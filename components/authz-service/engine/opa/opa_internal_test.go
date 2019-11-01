package opa

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/stretchr/testify/require"

	"github.com/chef/automate/components/authz-service/engine"
	"github.com/chef/automate/lib/logger"
)

// This test file is for micro-benchmarks

// these package variables are required so the compiler does not optimize return values out
var result ast.Value
var authzResponse bool

func genericInput(subjects []string, resource string, action string) (ast.Value, error) {
	subs := make([]interface{}, len(subjects))
	for i, sub := range subjects {
		subs[i] = sub
	}
	input := map[string]interface{}{
		"subjects": subs,
		"resource": resource,
		"action":   action,
	}
	return ast.InterfaceToValue(input)
}

func BenchmarkGenericInput(b *testing.B) {
	var r ast.Value
	var err error
	for n := 0; n < b.N; n++ {
		// always record the result to prevent the compiler eliminating the function
		// call.
		r, err = genericInput([]string{"user:local:alice@example.com", "team:local:admins"}, "some:resource", "read")
		if err != nil {
			b.Fatal(err)
		}
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkSpecificInput(b *testing.B) {
	var r ast.Value
	subjects, resource, action := []string{"user:local:alice@example.com", "team:local:admins"}, "some:resource", "read"
	for n := 0; n < b.N; n++ {
		subs := make(ast.Array, len(subjects))
		for i, sub := range subjects {
			subs[i] = ast.NewTerm(ast.String(sub))
		}
		r = ast.NewObject(
			[2]*ast.Term{ast.NewTerm(ast.String("subjects")), ast.NewTerm(subs)},
			[2]*ast.Term{ast.NewTerm(ast.String("resource")), ast.NewTerm(ast.String(resource))},
			[2]*ast.Term{ast.NewTerm(ast.String("action")), ast.NewTerm(ast.String(action))},
		)
	}
	result = r
}

var errResult error

func BenchmarkInitPartialResult(b *testing.B) {
	var r error
	ctx := context.Background()

	l, err := logger.NewLogger("text", "debug")
	require.NoError(b, err, "init logger")
	s, err := New(ctx, l)
	require.NoError(b, err, "init state")

	initStore := func(s *State, i int) {
		policies := make([]map[string]interface{}, i)
		for j := 0; j < i; j++ {
			policies[j] = map[string]interface{}{
				"subjects": engine.Subject("token:any"),
				"action":   "read",
				"resource": "cfgmgmt:nodes",
				"effect":   "allow",
			}
		}
		s.store = inmem.NewFromObject(map[string]interface{}{
			"policies": policies,
		})
	}

	for i := 0; i < 5; i++ {
		b.Run(fmt.Sprintf("store with %d items", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				initStore(s, i)
				r = s.initPartialResult(ctx)
				if r != nil {
					b.Error(r)
				}
			}
			errResult = r
		})
	}
}

// Q: What takes longer: parsing the module or compiling it?

var parsedResult *ast.Module

func BenchmarkCompareParseCompile(b *testing.B) {
	policy := MustAsset("policy/authz.rego")
	compiler := ast.NewCompiler()
	parsed, err := ast.ParseModule("authz.rego", string(policy))
	if err != nil {
		b.Fatal(err)
	}

	b.Run("parsing", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			parsed, err = ast.ParseModule("authz.rego", string(policy))
			if err != nil {
				b.Error(err)
			}
		}
		parsedResult = parsed
	})

	b.Run("compiling", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			compiler.Compile(map[string]*ast.Module{"authz.rego": parsed})
			if compiler.Failed() {
				b.Error(compiler.Errors)
			}
		}
	})

	// consistency check
	b.Run("parse and compile", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			parsed, err = ast.ParseModule("authz.rego", string(policy))
			if err != nil {
				b.Error(err)
			}
			compiler.Compile(map[string]*ast.Module{"authz.rego": parsed})
			if compiler.Failed() {
				b.Error(compiler.Errors)
			}
		}
		parsedResult = parsed
	})
}

// A: parsing is worse:
// BenchmarkCompareParseCompile/parsing-8           100	  14742148 ns/op	 4060407 B/op	  119417 allocs/op
// BenchmarkCompareParseCompile/compiling-8         200	  12774183 ns/op	 2148726 B/op	  112140 allocs/op
// BenchmarkCompareParseCompile/parse_and_compile-8  50	  34038622 ns/op	 6883385 B/op	  272642 allocs/op

func BenchmarkInitPartialResultWithPolicies(b *testing.B) {
	ctx := context.Background()

	l, err := logger.NewLogger("text", "debug")
	require.NoError(b, err, "init logger")
	s, err := New(ctx, l)
	require.NoError(b, err, "init state")

	policyCount := []int{0, 5, 10, 20, 50, 100, 200}

	for _, count := range policyCount {
		policies := testPolicies(count)
		s.store = inmem.NewFromObject(map[string]interface{}{
			"policies": policies,
		})

		b.Run(fmt.Sprintf("store with default policies and %d random policies", count), func(b *testing.B) {
			var r error

			for n := 0; n < b.N; n++ {
				r = s.initPartialResult(ctx)
				if r != nil {
					b.Error(r)
				}
			}
			errResult = r
		})

	}
}

func BenchmarkIsAuthorizedWithInitPartialResult(b *testing.B) {
	ctx := context.Background()

	l, err := logger.NewLogger("text", "debug")
	require.NoError(b, err, "init logger")
	s, err := New(ctx, l)
	require.NoError(b, err, "init state")

	policyCount := []int{0, 5, 10, 20, 50, 100, 200}

	for _, count := range policyCount {
		policies := testPolicies(count)
		s.store = inmem.NewFromObject(map[string]interface{}{
			"policies": policies,
		})
		err = s.initPartialResult(ctx)
		require.NoError(b, err, "init partial result")

		b.Run(fmt.Sprintf("store with default policies and %d random policies", count), func(b *testing.B) {
			var resp bool
			var err error
			for n := 0; n < b.N; n++ {
				resp, err = s.IsAuthorized(ctx, []string{"user:local:test@example.com"}, "read", "cfgmgmt:nodes")
				if err != nil {
					b.Error(err)
				}
			}
			authzResponse = resp
		})
	}
}

var filteredPairsResp []engine.Pair

func BenchmarkFilterAuthorizedPairsWithPolicies(b *testing.B) {
	ctx := context.Background()

	l, err := logger.NewLogger("text", "debug")
	require.NoError(b, err, "init logger")
	s, err := New(ctx, l)
	require.NoError(b, err, "init state")

	policyCount := []int{0, 1, 5, 10, 20, 50, 100}

	for _, count := range policyCount {
		policies := testPolicies(count)
		s.store = inmem.NewFromObject(map[string]interface{}{
			"policies": policies,
		})

		b.Run(fmt.Sprintf("store with default policies and %d random policies", count), func(b *testing.B) {
			var resp []engine.Pair
			var err error
			for n := 0; n < b.N; n++ {
				resp, err = s.FilterAuthorizedPairs(ctx, []string{"user:local:test@example.com"},
					[]engine.Pair{{Action: "read", Resource: "cfgmgmt:nodes"}})
				if err != nil {
					b.Error(err)
				}
			}
			filteredPairsResp = resp
		})
	}

	pairCount := []int{0, 1, 5, 10, 20, 50, 100}
	for _, count := range pairCount {
		s.store = inmem.NewFromObject(map[string]interface{}{
			"policies": testPolicies(0),
		})
		pairs := testPairs(count)

		b.Run(fmt.Sprintf("store with default policies and %d random pairs", count), func(b *testing.B) {
			var resp []engine.Pair
			var err error
			for n := 0; n < b.N; n++ {
				resp, err = s.FilterAuthorizedPairs(ctx, []string{"user:local:test@example.com"}, pairs)
				if err != nil {
					b.Error(err)
				}
			}
			filteredPairsResp = resp
		})
	}

	teamCount := []int{0, 1, 5, 10, 20, 50, 100}
	for _, count := range teamCount {
		policies := testPolicies(0)
		s.store = inmem.NewFromObject(map[string]interface{}{
			"policies": policies,
		})

		b.Run(fmt.Sprintf("store with default policies and %d random teams", count), func(b *testing.B) {
			var resp []engine.Pair
			var err error
			for n := 0; n < b.N; n++ {
				resp, err = s.FilterAuthorizedPairs(ctx, append([]string{"user:local:test@example.com"}, randomTeams(count)...),
					[]engine.Pair{{Action: "read", Resource: "cfgmgmt:nodes"}})
				if err != nil {
					b.Error(err)
				}
			}
			filteredPairsResp = resp
		})
	}
}

// At time of benchmarking, this is around half a second for 1000 teams
func BenchmarkV2FilterAuthorizedPairsRealWorldExample(b *testing.B) {
	ctx := context.Background()

	l, err := logger.NewLogger("text", "debug")
	require.NoError(b, err, "init logger")
	s, err := New(ctx, l)
	require.NoError(b, err, "init state")

	pairs := []engine.Pair{
		engine.Pair{Resource: "compliance:reporting:nodes", Action: "compliance:reportNodes:list"},
		engine.Pair{Resource: "iam:policies", Action: "iam:policies:list"},
		engine.Pair{Resource: "iam:teams", Action: "iam:teams:create"},
		engine.Pair{Resource: "system:config", Action: "system:telemetryConfig:get"},
		engine.Pair{Resource: "compliance:profiles:market", Action: "compliance:marketProfiles:get"},
		engine.Pair{Resource: "retention:nodes", Action: "retention:nodes:update"},
		engine.Pair{Resource: "compliance:reporting:reports", Action: "compliance:reports:list"},
		engine.Pair{Resource: "infra:actions", Action: "infra:ingest:create"},
		engine.Pair{Resource: "system:service:logLevel", Action: "system:serviceLogLevel:set"},
		engine.Pair{Resource: "iam:rules", Action: "iam:rules:apply"},
		engine.Pair{Resource: "secrets:secrets", Action: "secrets:secrets:create"},
		engine.Pair{Resource: "system:iam:upgradeToV2", Action: "system:iam:upgrade"},
		engine.Pair{Resource: "iam:teams", Action: "iam:teams:update"},
		engine.Pair{Resource: "system:service:version", Action: "system:serviceVersion:get"},
		engine.Pair{Resource: "iam:policies", Action: "iam:policies:create"},
		engine.Pair{Resource: "compliance:reporting:stats:trend", Action: "compliance:reportTrend:get"},
		engine.Pair{Resource: "compliance:profiles", Action: "compliance:profiles:list"},
		engine.Pair{Resource: "infra:nodes", Action: "infra:nodes:delete"},
		engine.Pair{Resource: "compliance:reporting:stats:failures", Action: "compliance:reportFailures:get"},
		engine.Pair{Resource: "system:health", Action: "system:health:get"},
		engine.Pair{Resource: "compliance:scanner:jobs", Action: "compliance:scannerJobs:create"},
		engine.Pair{Resource: "compliance:profiles", Action: "compliance:profiles:create"},
		engine.Pair{Resource: "secrets:secrets", Action: "secrets:secrets:list"},
		engine.Pair{Resource: "applications:serviceGroups", Action: "applications:serviceGroups:list"},
		engine.Pair{Resource: "event:events", Action: "event:events:list"},
		engine.Pair{Resource: "iam:projects", Action: "iam:projects:create"},
		engine.Pair{Resource: "system:iam:resetToV1", Action: "system:iam:reset"},
		engine.Pair{Resource: "compliance:reporting:stats:profiles", Action: "compliance:reportProfiles:get"},
		engine.Pair{Resource: "infra:status", Action: "infra:ingest:get"},
		engine.Pair{Resource: "compliance:reporting:licenseusage", Action: "compliance:reportingLicenseUsage:list"},
		engine.Pair{Resource: "iam:policyVersion", Action: "iam:policies:get"},
		engine.Pair{Resource: "iam:projects", Action: "iam:projects:list"},
		engine.Pair{Resource: "iam:introspect", Action: "iam:introspect:get"},
		engine.Pair{Resource: "system:status", Action: "system:license:get"},
		engine.Pair{Resource: "iam:teams", Action: "iam:teams:list"},
		engine.Pair{Resource: "infra:nodes", Action: "infra:nodes:list"},
		engine.Pair{Resource: "iam:introspect", Action: "iam:introspect:getAllProjects"},
		engine.Pair{Resource: "retention:nodes", Action: "retention:nodes:get"},
		engine.Pair{Resource: "retention:serviceGroups", Action: "retention:serviceGroups:update"},
		engine.Pair{Resource: "compliance:scanner:jobs", Action: "compliance:scannerJobs:list"},
		engine.Pair{Resource: "infra:nodes", Action: "infra:nodes:create"},
		engine.Pair{Resource: "iam:tokens", Action: "iam:tokens:create"},
		engine.Pair{Resource: "iam:rules", Action: "iam:rules:cancel"},
		engine.Pair{Resource: "iam:tokens", Action: "iam:tokens:list"},
		engine.Pair{Resource: "infra:nodeManagers", Action: "infra:nodeManagers:create"},
		engine.Pair{Resource: "compliance:reporting:profiles", Action: "compliance:reportProfiles:list"},
		engine.Pair{Resource: "system:license", Action: "system:license:apply"},
		engine.Pair{Resource: "iam:introspect", Action: "iam:introspect:getAll"},
		engine.Pair{Resource: "infra:nodes", Action: "infra:ingestNodes:delete"},
		engine.Pair{Resource: "iam:roles", Action: "iam:roles:list"},
		engine.Pair{Resource: "retention:serviceGroups", Action: "retention:serviceGroups:get"},
		engine.Pair{Resource: "notifications:rules", Action: "notifications:notifyRules:validate"},
		engine.Pair{Resource: "compliance:reporting:suggestions", Action: "compliance:reportSuggestions:list"},
		engine.Pair{Resource: "system:service:version", Action: "system:serviceVersion:list"},
		engine.Pair{Resource: "iam:users", Action: "iam:users:list"},
		engine.Pair{Resource: "infra:nodes", Action: "infra:ingest:delete"},
		engine.Pair{Resource: "notifications:rules", Action: "notifications:notifyRules:create"},
		engine.Pair{Resource: "iam:rules", Action: "iam:rules:status"},
		engine.Pair{Resource: "iam:introspect", Action: "iam:introspect:getSome"},
		engine.Pair{Resource: "infra:nodeManagers", Action: "infra:nodeManagers:list"},
		engine.Pair{Resource: "applications:serviceGroups", Action: "applications:serviceGroups:delete"},
		engine.Pair{Resource: "compliance:reporting:control", Action: "compliance:controlItems:list"},
		engine.Pair{Resource: "compliance:reporting:report-ids", Action: "compliance:reportids:list"},
		engine.Pair{Resource: "compliance:reporting:stats:summary", Action: "compliance:reportSummary:get"},
		engine.Pair{Resource: "iam:users", Action: "iam:users:create"},
		engine.Pair{Resource: "system:license", Action: "system:license:request"},
		engine.Pair{Resource: "notifications:rules", Action: "notifications:notifyRules:list"},
		engine.Pair{Resource: "iam:roles", Action: "iam:roles:create"},
	}

	jsonFile, err := os.Open("example_v2/real_world_2p1_store.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var pr struct {
		Policies map[string]interface{} `json:policies`
		Roles    map[string]interface{} `json:roles`
	}
	json.Unmarshal(byteValue, &pr)

	s.v2p1Store = inmem.NewFromObject(map[string]interface{}{
		"policies": pr.Policies,
		"roles":    pr.Roles,
	})

	teamCount := []int{0, 1, 10, 30, 50, 100, 150, 300, 500, 1000, 10000}
	for _, count := range teamCount {
		b.Run(fmt.Sprintf("V2FilterAuthorizedPairs with real life input and %d teams", count), func(b *testing.B) {
			var resp []engine.Pair
			var err error
			for n := 0; n < b.N; n++ {
				resp, err = s.V2FilterAuthorizedPairs(ctx, append([]string{"user:local:test@example.com"}, randomTeams(count)...), pairs)
				if err != nil {
					b.Error(err)
				}
			}
			filteredPairsResp = resp
		})
	}
}

// helpers

func randomTeams(c int) []string {
	ret := make([]string, c)
	for i := 0; i < c; i++ {
		ret[i] = fmt.Sprintf("team:local:team%d", i)
	}
	return ret
}

func testPairs(c int) []engine.Pair {
	ret := make([]engine.Pair, c)
	for i := 0; i < c; i++ {
		var pol map[string]interface{}
		// repeat until our random policy has no wildcard action
		// the resource count contain a wildcard, but it really doesn't matter;
		// it'll be treated as a string
		for pol = randomPolicies(1)[0]; pol["action"] == "*"; pol = randomPolicies(1)[0] {
		}

		// append the count to ensure we don't end up with duplicates
		ret[i] = engine.Pair{
			Action:   engine.Action(fmt.Sprintf("%s%d", pol["action"], i)),
			Resource: engine.Resource(fmt.Sprintf("%s%d", pol["resource"], i)),
		}
	}
	return ret
}

func defaultPolicies() []map[string]interface{} {
	allUsers := []string{"user:*"}
	allTokens := []string{"token:*"}
	admins := []string{"team:local:admins"}
	anyAction := "*"
	defaultEffect := "allow"

	defPolicyResources := []string{
		"compliance:*",
		"cfgmgmt:*",
		"cfgmgmt",
		"auth_introspection:*",
		"events",
		"events:*",
		"nodemanagers",
		"nodemanagers:*",
		"nodes",
		"nodes:*",
		"secrets",
		"secrets:*",
		"service_info:*",
	}
	defPolicyCount := len(defPolicyResources) + 5

	defPolicies := make([]map[string]interface{}, defPolicyCount)
	defPolicies[0] = map[string]interface{}{
		"subjects": admins,
		"resource": "*",
		"action":   anyAction,
		"effect":   defaultEffect,
	}
	defPolicies[1] = map[string]interface{}{
		"subjects": allTokens,
		"resource": "ingest:*",
		"action":   anyAction,
		"effect":   defaultEffect,
	}
	defPolicies[2] = map[string]interface{}{
		"subjects": allTokens,
		"resource": "compliance:profiles",
		"action":   "search",
		"effect":   defaultEffect,
	}
	defPolicies[3] = map[string]interface{}{
		"subjects": allTokens,
		"resource": "compliance:profiles:storage:*",
		"action":   "read",
		"effect":   defaultEffect,
	}
	defPolicies[4] = map[string]interface{}{
		"subjects": allTokens,
		"resource": "compliance:profiles:storage:*",
		"action":   "upload",
		"effect":   defaultEffect,
	}

	for j := 5; j < defPolicyCount; j++ {
		defPolicies[j] = map[string]interface{}{
			"subjects": allUsers,
			"action":   anyAction,
			"resource": defPolicyResources[j-5],
			"effect":   "allow",
		}
	}

	return defPolicies
}

// r is the number of random policies to include, appended to the default ones
func testPolicies(r int) []map[string]interface{} {
	return append(defaultPolicies(), randomPolicies(r)...)
}

func randomPolicies(i int) []map[string]interface{} {
	subjects := []string{"user:local:admin", "team:*", "team:local:sec", "team:local:admin", "user:ldap:*", "token:*"}
	resources := []string{
		"nodes",
		"nodes:123",
		"cfgmgmt:nodes:*",
		"notifications:rules",
		"cfgmgmt:stats",
		"compliance:profiles:alice",
		"compliance:reporting:reports:123",
		"events:*",
		"events:tasks",
		"secrets:*",
		"secrets:12345",
		"service_info:telemetry",
	}
	actions := []string{"read", "update", "delete", "create", "*", "search", "rerun", "count", "stop", "start"}

	newPolicies := make([]map[string]interface{}, i)
	for j := 0; j < i; j++ {
		rand.Seed(time.Now().Unix())
		subIndex := rand.Intn(len(subjects))
		resIndex := rand.Intn(len(resources))
		actIndex := rand.Intn(len(actions))

		newPolicies[j] = map[string]interface{}{
			"subjects": subjects[subIndex],
			"action":   actions[actIndex],
			"resource": resources[resIndex],
			"effect":   "allow",
		}
	}

	return newPolicies
}
