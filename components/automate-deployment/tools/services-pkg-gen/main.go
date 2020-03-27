package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/chef/automate/components/automate-deployment/pkg/assets"
	"github.com/chef/automate/components/automate-deployment/pkg/bind"
	"github.com/chef/automate/components/automate-deployment/pkg/habpkg"
	"github.com/chef/automate/lib/product"
)

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package generated

var ProductMetadataJSON = ` + "`" +
	`
{{ .ProductMetadataJSON }}
` + "`\n"))

const deploymentServiceName = "deployment-service"

func fatal(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s", msg, err.Error())
	os.Exit(1)
}

func removeDeploymentServiceFromPackages(packages []*product.Package) []*product.Package {
	for i := range packages {
		if packages[i].Name.Name == deploymentServiceName {
			return append(packages[:i], packages[i+1:]...)
		}
	}
	panic("deployment-service not found")
}

func removeDeploymentServiceFromCollections(collections []*product.Collection) {
	for _, c := range collections {
	COL:
		for i := range c.Services {
			if c.Services[i].Name == deploymentServiceName {
				c.Services = append(c.Services[:i], c.Services[i+1:]...)
				break COL
			}
		}
	}

}

func sortPackages(packages []*product.Package) []*product.Package {
	binds, err := bind.ParseData([]byte(assets.BindData))
	if err != nil {
		fatal("failed to parse bind data", err)
	}

	allPkgs := []habpkg.HabPkg{}
	packageSet := make(map[string]*product.Package)
	for _, p := range packages {
		allPkgs = append(allPkgs, habpkg.New(p.Name.Origin, p.Name.Name))
		packageSet[p.Name.Name] = p
	}

	allSortedPkgs, err := bind.TopoSortAll(allPkgs, binds)
	if err != nil {
		fatal("failed to sort services", err)
	}

	sorted := make([]*product.Package, len(allSortedPkgs))
	for i := range allSortedPkgs {
		sorted[i] = packageSet[allSortedPkgs[i].Name()]
	}
	return sorted
}

func main() {
	repoRoot := os.Args[1]
	metadata, err := product.Parse(repoRoot)
	if err != nil {
		panic(err)
	}

	packages := removeDeploymentServiceFromPackages(metadata.Packages)
	metadata.Packages = sortPackages(packages)
	removeDeploymentServiceFromCollections(metadata.Collections)

	outStruct := product.Metadata{
		Packages:        metadata.Packages,
		DeletedPackages: metadata.DeletedPackages,
		Collections:     metadata.Collections,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	data, err := json.MarshalIndent(outStruct, "", "  ")
	if err != nil {
		panic(err)
	}
	f, err := os.Create(os.Args[2])
	if err != nil {
		fatal("failed to write file", err)
	}
	defer f.Close()
	err = packageTemplate.Execute(f, struct {
		ProductMetadataJSON string
	}{
		ProductMetadataJSON: string(data),
	})
	if err != nil {
		fatal("failed to write template", err)
	}
}
