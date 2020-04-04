// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/response/policyfiles.proto

package response

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Policyfiles struct {
	// List of policy files.
	Policies             []*PolicyfileListItem `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Policyfiles) Reset()         { *m = Policyfiles{} }
func (m *Policyfiles) String() string { return proto.CompactTextString(m) }
func (*Policyfiles) ProtoMessage()    {}
func (*Policyfiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{0}
}

func (m *Policyfiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policyfiles.Unmarshal(m, b)
}
func (m *Policyfiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policyfiles.Marshal(b, m, deterministic)
}
func (m *Policyfiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policyfiles.Merge(m, src)
}
func (m *Policyfiles) XXX_Size() int {
	return xxx_messageInfo_Policyfiles.Size(m)
}
func (m *Policyfiles) XXX_DiscardUnknown() {
	xxx_messageInfo_Policyfiles.DiscardUnknown(m)
}

var xxx_messageInfo_Policyfiles proto.InternalMessageInfo

func (m *Policyfiles) GetPolicies() []*PolicyfileListItem {
	if m != nil {
		return m.Policies
	}
	return nil
}

type PolicyfileListItem struct {
	// Name of the policy file.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Revision ID of the policy file.
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Policy group of the policy file.
	PolicyGroup          string   `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyfileListItem) Reset()         { *m = PolicyfileListItem{} }
func (m *PolicyfileListItem) String() string { return proto.CompactTextString(m) }
func (*PolicyfileListItem) ProtoMessage()    {}
func (*PolicyfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{1}
}

func (m *PolicyfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyfileListItem.Unmarshal(m, b)
}
func (m *PolicyfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyfileListItem.Marshal(b, m, deterministic)
}
func (m *PolicyfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyfileListItem.Merge(m, src)
}
func (m *PolicyfileListItem) XXX_Size() int {
	return xxx_messageInfo_PolicyfileListItem.Size(m)
}
func (m *PolicyfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyfileListItem proto.InternalMessageInfo

func (m *PolicyfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PolicyfileListItem) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func (m *PolicyfileListItem) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

type Policyfile struct {
	// Name of the policy file.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of the policy group.
	PolicyGroup string `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	// Revision Id of the policy.
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Run list associated with the policy.
	RunList []string `protobuf:"bytes,4,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	// Included policy lock files.
	IncludedPolicyLocks *IncludedPolicyLocks `protobuf:"bytes,5,opt,name=included_policy_locks,json=includedPolicyLocks,proto3" json:"included_policy_locks,omitempty"`
	// Named the run list associated with the policy.
	NamedRunList []*NamedRunList `protobuf:"bytes,6,rep,name=named_run_list,json=namedRunList,proto3" json:"named_run_list,omitempty"`
	// Stringified json of the default attributes.
	DefaultAttributes string `protobuf:"bytes,7,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Stringified json of the override attributes.
	OverrideAttributes string `protobuf:"bytes,8,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
	// Solution dependencies of the policy file.
	SolutionDependencies string   `protobuf:"bytes,9,opt,name=solution_dependencies,json=solutionDependencies,proto3" json:"solution_dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policyfile) Reset()         { *m = Policyfile{} }
func (m *Policyfile) String() string { return proto.CompactTextString(m) }
func (*Policyfile) ProtoMessage()    {}
func (*Policyfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{2}
}

func (m *Policyfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policyfile.Unmarshal(m, b)
}
func (m *Policyfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policyfile.Marshal(b, m, deterministic)
}
func (m *Policyfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policyfile.Merge(m, src)
}
func (m *Policyfile) XXX_Size() int {
	return xxx_messageInfo_Policyfile.Size(m)
}
func (m *Policyfile) XXX_DiscardUnknown() {
	xxx_messageInfo_Policyfile.DiscardUnknown(m)
}

var xxx_messageInfo_Policyfile proto.InternalMessageInfo

func (m *Policyfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policyfile) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *Policyfile) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func (m *Policyfile) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func (m *Policyfile) GetIncludedPolicyLocks() *IncludedPolicyLocks {
	if m != nil {
		return m.IncludedPolicyLocks
	}
	return nil
}

func (m *Policyfile) GetNamedRunList() []*NamedRunList {
	if m != nil {
		return m.NamedRunList
	}
	return nil
}

func (m *Policyfile) GetDefaultAttributes() string {
	if m != nil {
		return m.DefaultAttributes
	}
	return ""
}

func (m *Policyfile) GetOverrideAttributes() string {
	if m != nil {
		return m.OverrideAttributes
	}
	return ""
}

func (m *Policyfile) GetSolutionDependencies() string {
	if m != nil {
		return m.SolutionDependencies
	}
	return ""
}

type IncludedPolicyLocks struct {
	// Name of the included policy file.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Revision id of the included policy file.
	RevisionId string `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Source options of the included policy file.
	SourceOptions        *SourceOptions `protobuf:"bytes,3,opt,name=source_options,json=sourceOptions,proto3" json:"source_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IncludedPolicyLocks) Reset()         { *m = IncludedPolicyLocks{} }
func (m *IncludedPolicyLocks) String() string { return proto.CompactTextString(m) }
func (*IncludedPolicyLocks) ProtoMessage()    {}
func (*IncludedPolicyLocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{3}
}

func (m *IncludedPolicyLocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncludedPolicyLocks.Unmarshal(m, b)
}
func (m *IncludedPolicyLocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncludedPolicyLocks.Marshal(b, m, deterministic)
}
func (m *IncludedPolicyLocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncludedPolicyLocks.Merge(m, src)
}
func (m *IncludedPolicyLocks) XXX_Size() int {
	return xxx_messageInfo_IncludedPolicyLocks.Size(m)
}
func (m *IncludedPolicyLocks) XXX_DiscardUnknown() {
	xxx_messageInfo_IncludedPolicyLocks.DiscardUnknown(m)
}

var xxx_messageInfo_IncludedPolicyLocks proto.InternalMessageInfo

func (m *IncludedPolicyLocks) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IncludedPolicyLocks) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func (m *IncludedPolicyLocks) GetSourceOptions() *SourceOptions {
	if m != nil {
		return m.SourceOptions
	}
	return nil
}

type SourceOptions struct {
	// Path of the source options.
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceOptions) Reset()         { *m = SourceOptions{} }
func (m *SourceOptions) String() string { return proto.CompactTextString(m) }
func (*SourceOptions) ProtoMessage()    {}
func (*SourceOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{4}
}

func (m *SourceOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceOptions.Unmarshal(m, b)
}
func (m *SourceOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceOptions.Marshal(b, m, deterministic)
}
func (m *SourceOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceOptions.Merge(m, src)
}
func (m *SourceOptions) XXX_Size() int {
	return xxx_messageInfo_SourceOptions.Size(m)
}
func (m *SourceOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SourceOptions proto.InternalMessageInfo

func (m *SourceOptions) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type NamedRunList struct {
	// Name of the run list.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Run list associated with the policy.
	RunList              []string `protobuf:"bytes,2,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamedRunList) Reset()         { *m = NamedRunList{} }
func (m *NamedRunList) String() string { return proto.CompactTextString(m) }
func (*NamedRunList) ProtoMessage()    {}
func (*NamedRunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{5}
}

func (m *NamedRunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedRunList.Unmarshal(m, b)
}
func (m *NamedRunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedRunList.Marshal(b, m, deterministic)
}
func (m *NamedRunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedRunList.Merge(m, src)
}
func (m *NamedRunList) XXX_Size() int {
	return xxx_messageInfo_NamedRunList.Size(m)
}
func (m *NamedRunList) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedRunList.DiscardUnknown(m)
}

var xxx_messageInfo_NamedRunList proto.InternalMessageInfo

func (m *NamedRunList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedRunList) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func init() {
	proto.RegisterType((*Policyfiles)(nil), "chef.automate.api.infra_proxy.response.Policyfiles")
	proto.RegisterType((*PolicyfileListItem)(nil), "chef.automate.api.infra_proxy.response.PolicyfileListItem")
	proto.RegisterType((*Policyfile)(nil), "chef.automate.api.infra_proxy.response.Policyfile")
	proto.RegisterType((*IncludedPolicyLocks)(nil), "chef.automate.api.infra_proxy.response.IncludedPolicyLocks")
	proto.RegisterType((*SourceOptions)(nil), "chef.automate.api.infra_proxy.response.SourceOptions")
	proto.RegisterType((*NamedRunList)(nil), "chef.automate.api.infra_proxy.response.NamedRunList")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/response/policyfiles.proto", fileDescriptor_64fca7030dbc5441)
}

var fileDescriptor_64fca7030dbc5441 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x25, 0x49, 0x6d, 0x93, 0xbb, 0x69, 0xc1, 0x89, 0x85, 0xf5, 0xc9, 0xb8, 0x82, 0xe4, 0xc5,
	0x5d, 0x48, 0xf4, 0xa5, 0xea, 0x83, 0x22, 0x48, 0xa0, 0xa8, 0xac, 0xe0, 0x43, 0x11, 0x96, 0xc9,
	0xee, 0x4d, 0x33, 0xb8, 0x99, 0x19, 0xe6, 0xa3, 0xb4, 0xff, 0xc8, 0xbf, 0xe5, 0x3f, 0x91, 0x99,
	0xcd, 0x26, 0x5b, 0x12, 0xc8, 0xbe, 0xcd, 0xdc, 0x33, 0xe7, 0xdc, 0x73, 0x3f, 0x18, 0x98, 0x51,
	0xc9, 0x12, 0xbc, 0x37, 0xa8, 0x38, 0x2d, 0x13, 0xc6, 0x97, 0x8a, 0x66, 0x52, 0x89, 0xfb, 0x87,
	0x44, 0xa1, 0x96, 0x82, 0x6b, 0x4c, 0xa4, 0x28, 0x59, 0xfe, 0xb0, 0x64, 0x25, 0xea, 0x58, 0x2a,
	0x61, 0x04, 0x79, 0x9d, 0xaf, 0x70, 0x19, 0x53, 0x6b, 0xc4, 0x9a, 0x1a, 0x8c, 0xa9, 0x64, 0x71,
	0x83, 0x19, 0xd7, 0xcc, 0x08, 0x21, 0xf8, 0xb1, 0x23, 0x93, 0x5f, 0xd0, 0xf7, 0x5a, 0x0c, 0x75,
	0xd8, 0x19, 0xf7, 0x26, 0xc1, 0xf4, 0x2a, 0x6e, 0xa7, 0x14, 0xef, 0x64, 0xae, 0x99, 0x36, 0x73,
	0x83, 0xeb, 0x74, 0xab, 0x15, 0x95, 0x40, 0xf6, 0x71, 0x42, 0xe0, 0x84, 0xd3, 0x35, 0x86, 0x9d,
	0x71, 0x67, 0x32, 0x48, 0xfd, 0x99, 0xbc, 0x80, 0x40, 0xe1, 0x1d, 0xd3, 0x4c, 0xf0, 0x8c, 0x15,
	0x61, 0xcf, 0x43, 0x50, 0x87, 0xe6, 0x05, 0x79, 0x09, 0xc3, 0xaa, 0xdc, 0xec, 0x56, 0x09, 0x2b,
	0xc3, 0xae, 0x7f, 0x11, 0x54, 0xb1, 0xaf, 0x2e, 0x14, 0xfd, 0xeb, 0x01, 0xec, 0xd2, 0x1d, 0x4c,
	0x73, 0x5c, 0xe5, 0xb8, 0x93, 0xe7, 0xd0, 0x57, 0x96, 0x67, 0x25, 0xd3, 0x26, 0x3c, 0x19, 0xf7,
	0x26, 0x83, 0xf4, 0x4c, 0x59, 0xee, 0xaa, 0x23, 0x02, 0x2e, 0x19, 0xcf, 0x4b, 0x5b, 0x60, 0x91,
	0x6d, 0xf2, 0x94, 0x22, 0xff, 0xa3, 0xc3, 0x27, 0xe3, 0xce, 0x24, 0x98, 0xbe, 0x6f, 0xdb, 0xd4,
	0xf9, 0x46, 0xa4, 0xaa, 0xe6, 0xda, 0x49, 0xa4, 0x23, 0xb6, 0x1f, 0x24, 0x37, 0x70, 0xe1, 0xea,
	0x2a, 0xb2, 0xad, 0xa3, 0x53, 0x3f, 0xbe, 0xb7, 0x6d, 0x33, 0x7d, 0x73, 0xec, 0xb4, 0xb2, 0x9f,
	0x0e, 0x79, 0xe3, 0x46, 0xde, 0x00, 0x29, 0x70, 0x49, 0x6d, 0x69, 0x32, 0x6a, 0x8c, 0x62, 0x0b,
	0x6b, 0x50, 0x87, 0x67, 0xbe, 0x1f, 0x4f, 0x37, 0xc8, 0xa7, 0x2d, 0x40, 0x12, 0x18, 0x89, 0x3b,
	0x54, 0x8a, 0x15, 0xd8, 0x7c, 0xdf, 0xf7, 0xef, 0x49, 0x0d, 0x35, 0x08, 0x33, 0xb8, 0xd4, 0xa2,
	0xb4, 0xc6, 0x35, 0xba, 0x40, 0x89, 0xbc, 0x40, 0xee, 0x37, 0x70, 0xe0, 0x29, 0xcf, 0x6a, 0xf0,
	0x4b, 0x03, 0x8b, 0xfe, 0x76, 0x60, 0x74, 0xa0, 0x3b, 0x6d, 0x76, 0xaa, 0xbb, 0x37, 0xc9, 0xdf,
	0x70, 0xa1, 0x85, 0x55, 0x39, 0x66, 0x42, 0xba, 0x4c, 0xda, 0x4f, 0x3b, 0x98, 0xbe, 0x6b, 0xdb,
	0xbd, 0x9f, 0x9e, 0xfd, 0xbd, 0x22, 0xa7, 0xe7, 0xba, 0x79, 0x8d, 0x5e, 0xc1, 0xf9, 0x23, 0xdc,
	0x79, 0x94, 0xd4, 0xac, 0x6a, 0x8f, 0xee, 0x1c, 0x7d, 0x84, 0x61, 0x73, 0x04, 0x07, 0xeb, 0x68,
	0x2e, 0x5c, 0xf7, 0xd1, 0xc2, 0x7d, 0xfe, 0x70, 0x73, 0x75, 0xcb, 0xcc, 0xca, 0x2e, 0xe2, 0x5c,
	0xac, 0x13, 0xe7, 0x3a, 0xa9, 0x5d, 0x27, 0x47, 0xff, 0x8f, 0xc5, 0xa9, 0xff, 0x34, 0x66, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xd9, 0xb7, 0xfe, 0x6b, 0x04, 0x00, 0x00,
}
