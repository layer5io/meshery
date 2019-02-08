// Code generated by protoc-gen-go. DO NOT EDIT.
// source: meshops.proto

package meshes

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type CreateMeshInstanceRequest struct {
	K8SConfig            []byte   `protobuf:"bytes,1,opt,name=k8sConfig,proto3" json:"k8sConfig,omitempty"`
	ContextName          string   `protobuf:"bytes,2,opt,name=contextName,proto3" json:"contextName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMeshInstanceRequest) Reset()         { *m = CreateMeshInstanceRequest{} }
func (m *CreateMeshInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMeshInstanceRequest) ProtoMessage()    {}
func (*CreateMeshInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_881788560c20cf7b, []int{0}
}

func (m *CreateMeshInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMeshInstanceRequest.Unmarshal(m, b)
}
func (m *CreateMeshInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMeshInstanceRequest.Marshal(b, m, deterministic)
}
func (m *CreateMeshInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMeshInstanceRequest.Merge(m, src)
}
func (m *CreateMeshInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMeshInstanceRequest.Size(m)
}
func (m *CreateMeshInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMeshInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMeshInstanceRequest proto.InternalMessageInfo

func (m *CreateMeshInstanceRequest) GetK8SConfig() []byte {
	if m != nil {
		return m.K8SConfig
	}
	return nil
}

func (m *CreateMeshInstanceRequest) GetContextName() string {
	if m != nil {
		return m.ContextName
	}
	return ""
}

type CreateMeshInstanceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMeshInstanceResponse) Reset()         { *m = CreateMeshInstanceResponse{} }
func (m *CreateMeshInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMeshInstanceResponse) ProtoMessage()    {}
func (*CreateMeshInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_881788560c20cf7b, []int{1}
}

func (m *CreateMeshInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMeshInstanceResponse.Unmarshal(m, b)
}
func (m *CreateMeshInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMeshInstanceResponse.Marshal(b, m, deterministic)
}
func (m *CreateMeshInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMeshInstanceResponse.Merge(m, src)
}
func (m *CreateMeshInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMeshInstanceResponse.Size(m)
}
func (m *CreateMeshInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMeshInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMeshInstanceResponse proto.InternalMessageInfo

type MeshNameRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshNameRequest) Reset()         { *m = MeshNameRequest{} }
func (m *MeshNameRequest) String() string { return proto.CompactTextString(m) }
func (*MeshNameRequest) ProtoMessage()    {}
func (*MeshNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_881788560c20cf7b, []int{2}
}

func (m *MeshNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshNameRequest.Unmarshal(m, b)
}
func (m *MeshNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshNameRequest.Marshal(b, m, deterministic)
}
func (m *MeshNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshNameRequest.Merge(m, src)
}
func (m *MeshNameRequest) XXX_Size() int {
	return xxx_messageInfo_MeshNameRequest.Size(m)
}
func (m *MeshNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MeshNameRequest proto.InternalMessageInfo

type MeshNameResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshNameResponse) Reset()         { *m = MeshNameResponse{} }
func (m *MeshNameResponse) String() string { return proto.CompactTextString(m) }
func (*MeshNameResponse) ProtoMessage()    {}
func (*MeshNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_881788560c20cf7b, []int{3}
}

func (m *MeshNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshNameResponse.Unmarshal(m, b)
}
func (m *MeshNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshNameResponse.Marshal(b, m, deterministic)
}
func (m *MeshNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshNameResponse.Merge(m, src)
}
func (m *MeshNameResponse) XXX_Size() int {
	return xxx_messageInfo_MeshNameResponse.Size(m)
}
func (m *MeshNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MeshNameResponse proto.InternalMessageInfo

func (m *MeshNameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplyRuleRequest struct {
	OpName               string   `protobuf:"bytes,1,opt,name=opName,proto3" json:"opName,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRuleRequest) Reset()         { *m = ApplyRuleRequest{} }
func (m *ApplyRuleRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyRuleRequest) ProtoMessage()    {}
func (*ApplyRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_881788560c20cf7b, []int{4}
}

func (m *ApplyRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRuleRequest.Unmarshal(m, b)
}
func (m *ApplyRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRuleRequest.Marshal(b, m, deterministic)
}
func (m *ApplyRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRuleRequest.Merge(m, src)
}
func (m *ApplyRuleRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyRuleRequest.Size(m)
}
func (m *ApplyRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRuleRequest proto.InternalMessageInfo

func (m *ApplyRuleRequest) GetOpName() string {
	if m != nil {
		return m.OpName
	}
	return ""
}

func (m *ApplyRuleRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ApplyRuleRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ApplyRuleResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRuleResponse) Reset()         { *m = ApplyRuleResponse{} }
func (m *ApplyRuleResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyRuleResponse) ProtoMessage()    {}
func (*ApplyRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_881788560c20cf7b, []int{5}
}

func (m *ApplyRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRuleResponse.Unmarshal(m, b)
}
func (m *ApplyRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRuleResponse.Marshal(b, m, deterministic)
}
func (m *ApplyRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRuleResponse.Merge(m, src)
}
func (m *ApplyRuleResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyRuleResponse.Size(m)
}
func (m *ApplyRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRuleResponse proto.InternalMessageInfo

func (m *ApplyRuleResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SupportedOperationsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupportedOperationsRequest) Reset()         { *m = SupportedOperationsRequest{} }
func (m *SupportedOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*SupportedOperationsRequest) ProtoMessage()    {}
func (*SupportedOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_881788560c20cf7b, []int{6}
}

func (m *SupportedOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportedOperationsRequest.Unmarshal(m, b)
}
func (m *SupportedOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportedOperationsRequest.Marshal(b, m, deterministic)
}
func (m *SupportedOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedOperationsRequest.Merge(m, src)
}
func (m *SupportedOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_SupportedOperationsRequest.Size(m)
}
func (m *SupportedOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedOperationsRequest proto.InternalMessageInfo

type SupportedOperationsResponse struct {
	Ops                  map[string]string `protobuf:"bytes,1,rep,name=ops,proto3" json:"ops,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Error                string            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SupportedOperationsResponse) Reset()         { *m = SupportedOperationsResponse{} }
func (m *SupportedOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*SupportedOperationsResponse) ProtoMessage()    {}
func (*SupportedOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_881788560c20cf7b, []int{7}
}

func (m *SupportedOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportedOperationsResponse.Unmarshal(m, b)
}
func (m *SupportedOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportedOperationsResponse.Marshal(b, m, deterministic)
}
func (m *SupportedOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedOperationsResponse.Merge(m, src)
}
func (m *SupportedOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_SupportedOperationsResponse.Size(m)
}
func (m *SupportedOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedOperationsResponse proto.InternalMessageInfo

func (m *SupportedOperationsResponse) GetOps() map[string]string {
	if m != nil {
		return m.Ops
	}
	return nil
}

func (m *SupportedOperationsResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SupportedOperation struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupportedOperation) Reset()         { *m = SupportedOperation{} }
func (m *SupportedOperation) String() string { return proto.CompactTextString(m) }
func (*SupportedOperation) ProtoMessage()    {}
func (*SupportedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_881788560c20cf7b, []int{8}
}

func (m *SupportedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportedOperation.Unmarshal(m, b)
}
func (m *SupportedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportedOperation.Marshal(b, m, deterministic)
}
func (m *SupportedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedOperation.Merge(m, src)
}
func (m *SupportedOperation) XXX_Size() int {
	return xxx_messageInfo_SupportedOperation.Size(m)
}
func (m *SupportedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedOperation proto.InternalMessageInfo

func (m *SupportedOperation) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SupportedOperation) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateMeshInstanceRequest)(nil), "meshes.CreateMeshInstanceRequest")
	proto.RegisterType((*CreateMeshInstanceResponse)(nil), "meshes.CreateMeshInstanceResponse")
	proto.RegisterType((*MeshNameRequest)(nil), "meshes.MeshNameRequest")
	proto.RegisterType((*MeshNameResponse)(nil), "meshes.MeshNameResponse")
	proto.RegisterType((*ApplyRuleRequest)(nil), "meshes.ApplyRuleRequest")
	proto.RegisterType((*ApplyRuleResponse)(nil), "meshes.ApplyRuleResponse")
	proto.RegisterType((*SupportedOperationsRequest)(nil), "meshes.SupportedOperationsRequest")
	proto.RegisterType((*SupportedOperationsResponse)(nil), "meshes.SupportedOperationsResponse")
	proto.RegisterMapType((map[string]string)(nil), "meshes.SupportedOperationsResponse.OpsEntry")
	proto.RegisterType((*SupportedOperation)(nil), "meshes.SupportedOperation")
}

func init() { proto.RegisterFile("meshops.proto", fileDescriptor_881788560c20cf7b) }

var fileDescriptor_881788560c20cf7b = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdb, 0xae, 0xd2, 0x40,
	0x14, 0xa5, 0xad, 0x12, 0xd8, 0x78, 0x81, 0xd1, 0x68, 0xa9, 0x3c, 0xd4, 0x31, 0x31, 0x98, 0x98,
	0x3e, 0x60, 0x62, 0x88, 0x31, 0x1a, 0x43, 0x8c, 0xf1, 0x41, 0x48, 0xca, 0xa3, 0x31, 0xb1, 0x96,
	0xad, 0x10, 0x60, 0x66, 0x9c, 0x99, 0x92, 0xc3, 0x37, 0x9d, 0x1f, 0x3a, 0x9f, 0x73, 0x32, 0xbd,
	0x12, 0xe8, 0x21, 0xbc, 0xcd, 0xbe, 0xad, 0xbd, 0x76, 0xd7, 0x2a, 0x3c, 0xdc, 0xa2, 0x5a, 0x72,
	0xa1, 0x02, 0x21, 0xb9, 0xe6, 0xa4, 0x69, 0x42, 0x54, 0xf4, 0x27, 0xf4, 0x27, 0x12, 0x23, 0x8d,
	0x3f, 0x50, 0x2d, 0xbf, 0x33, 0xa5, 0x23, 0x16, 0x63, 0x88, 0xff, 0x13, 0x54, 0x9a, 0x0c, 0xa0,
	0xbd, 0x1e, 0xab, 0x09, 0x67, 0x7f, 0x57, 0xff, 0x5c, 0xcb, 0xb7, 0x86, 0x0f, 0xc2, 0x2a, 0x41,
	0x7c, 0xe8, 0xc4, 0x9c, 0x69, 0xbc, 0xd2, 0xd3, 0x68, 0x8b, 0xae, 0xed, 0x5b, 0xc3, 0x76, 0x78,
	0x98, 0xa2, 0x03, 0xf0, 0xea, 0xc0, 0x95, 0xe0, 0x4c, 0x21, 0xed, 0xc1, 0x63, 0x93, 0x37, 0x9d,
	0xf9, 0x42, 0xfa, 0x1a, 0xba, 0x55, 0x2a, 0x6b, 0x23, 0x04, 0xee, 0x31, 0x83, 0x6f, 0xa5, 0xf8,
	0xe9, 0x9b, 0x2e, 0xa0, 0xfb, 0x45, 0x88, 0xcd, 0x3e, 0x4c, 0x36, 0x25, 0xd9, 0x67, 0xd0, 0xe4,
	0x62, 0x5a, 0x75, 0xe6, 0x91, 0x39, 0xc2, 0xcc, 0x28, 0x11, 0xc5, 0x05, 0xc9, 0x2a, 0x41, 0x3c,
	0x68, 0x25, 0x0a, 0x65, 0xba, 0xc1, 0x49, 0x8b, 0x65, 0x4c, 0xdf, 0x40, 0xef, 0x60, 0x4b, 0x4e,
	0xe7, 0x29, 0xdc, 0x47, 0x29, 0xb9, 0xcc, 0xb7, 0x64, 0x81, 0xb9, 0x74, 0x9e, 0x08, 0xc1, 0xa5,
	0xc6, 0xc5, 0x4c, 0xa0, 0x8c, 0xf4, 0x8a, 0x33, 0x55, 0x9c, 0x75, 0x6d, 0xc1, 0x8b, 0xda, 0x72,
	0x8e, 0xf9, 0x09, 0x1c, 0x2e, 0x94, 0x6b, 0xf9, 0xce, 0xb0, 0x33, 0x7a, 0x1b, 0x64, 0xd2, 0x04,
	0x67, 0x26, 0x82, 0x99, 0x50, 0x5f, 0x99, 0x96, 0xfb, 0xd0, 0x0c, 0x56, 0x9c, 0xec, 0x03, 0x4e,
	0xde, 0x7b, 0x68, 0x15, 0x6d, 0xa4, 0x0b, 0xce, 0x1a, 0xf7, 0x39, 0x67, 0xf3, 0x34, 0x33, 0xbb,
	0x68, 0x93, 0x14, 0x9f, 0x24, 0x0b, 0x3e, 0xd8, 0x63, 0x8b, 0x7e, 0x04, 0x72, 0xba, 0xfa, 0x52,
	0x84, 0xd1, 0x8d, 0x0d, 0x1d, 0xa3, 0xe1, 0x1c, 0xe5, 0x6e, 0x15, 0x23, 0xf9, 0x05, 0xe4, 0xd4,
	0x03, 0xe4, 0x65, 0x71, 0xe4, 0x9d, 0xe6, 0xf3, 0xe8, 0xb9, 0x96, 0xdc, 0x42, 0x0d, 0xf2, 0x19,
	0x5a, 0x85, 0x63, 0xc8, 0xf3, 0x62, 0xe2, 0xc8, 0x56, 0x9e, 0x7b, 0x5a, 0x28, 0x01, 0xbe, 0xc1,
	0xa3, 0x54, 0xe4, 0xea, 0xd2, 0xb2, 0xfb, 0xd8, 0x62, 0x5e, 0xbf, 0xa6, 0x52, 0x02, 0xfd, 0x86,
	0x27, 0x35, 0x8a, 0x11, 0x7a, 0x56, 0xce, 0x0c, 0xf7, 0xd5, 0x05, 0x92, 0xd3, 0xc6, 0x9f, 0x66,
	0xfa, 0xeb, 0xbe, 0xbb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x20, 0x51, 0x4f, 0xaa, 0xcb, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MeshServiceClient is the client API for MeshService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeshServiceClient interface {
	CreateMeshInstance(ctx context.Context, in *CreateMeshInstanceRequest, opts ...grpc.CallOption) (*CreateMeshInstanceResponse, error)
	MeshName(ctx context.Context, in *MeshNameRequest, opts ...grpc.CallOption) (*MeshNameResponse, error)
	ApplyOperation(ctx context.Context, in *ApplyRuleRequest, opts ...grpc.CallOption) (*ApplyRuleResponse, error)
	SupportedOperations(ctx context.Context, in *SupportedOperationsRequest, opts ...grpc.CallOption) (*SupportedOperationsResponse, error)
}

type meshServiceClient struct {
	cc *grpc.ClientConn
}

func NewMeshServiceClient(cc *grpc.ClientConn) MeshServiceClient {
	return &meshServiceClient{cc}
}

func (c *meshServiceClient) CreateMeshInstance(ctx context.Context, in *CreateMeshInstanceRequest, opts ...grpc.CallOption) (*CreateMeshInstanceResponse, error) {
	out := new(CreateMeshInstanceResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/CreateMeshInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) MeshName(ctx context.Context, in *MeshNameRequest, opts ...grpc.CallOption) (*MeshNameResponse, error) {
	out := new(MeshNameResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/MeshName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) ApplyOperation(ctx context.Context, in *ApplyRuleRequest, opts ...grpc.CallOption) (*ApplyRuleResponse, error) {
	out := new(ApplyRuleResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/ApplyOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) SupportedOperations(ctx context.Context, in *SupportedOperationsRequest, opts ...grpc.CallOption) (*SupportedOperationsResponse, error) {
	out := new(SupportedOperationsResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/SupportedOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeshServiceServer is the server API for MeshService service.
type MeshServiceServer interface {
	CreateMeshInstance(context.Context, *CreateMeshInstanceRequest) (*CreateMeshInstanceResponse, error)
	MeshName(context.Context, *MeshNameRequest) (*MeshNameResponse, error)
	ApplyOperation(context.Context, *ApplyRuleRequest) (*ApplyRuleResponse, error)
	SupportedOperations(context.Context, *SupportedOperationsRequest) (*SupportedOperationsResponse, error)
}

func RegisterMeshServiceServer(s *grpc.Server, srv MeshServiceServer) {
	s.RegisterService(&_MeshService_serviceDesc, srv)
}

func _MeshService_CreateMeshInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMeshInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).CreateMeshInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/CreateMeshInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).CreateMeshInstance(ctx, req.(*CreateMeshInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_MeshName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeshNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).MeshName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/MeshName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).MeshName(ctx, req.(*MeshNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_ApplyOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).ApplyOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/ApplyOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).ApplyOperation(ctx, req.(*ApplyRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_SupportedOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportedOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).SupportedOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/SupportedOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).SupportedOperations(ctx, req.(*SupportedOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MeshService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "meshes.MeshService",
	HandlerType: (*MeshServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMeshInstance",
			Handler:    _MeshService_CreateMeshInstance_Handler,
		},
		{
			MethodName: "MeshName",
			Handler:    _MeshService_MeshName_Handler,
		},
		{
			MethodName: "ApplyOperation",
			Handler:    _MeshService_ApplyOperation_Handler,
		},
		{
			MethodName: "SupportedOperations",
			Handler:    _MeshService_SupportedOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meshops.proto",
}
