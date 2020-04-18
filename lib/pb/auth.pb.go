// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/auth.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AuthRole int32

const (
	AuthRole_USER  AuthRole = 0
	AuthRole_ADMIN AuthRole = 1
)

var AuthRole_name = map[int32]string{
	0: "USER",
	1: "ADMIN",
}

var AuthRole_value = map[string]int32{
	"USER":  0,
	"ADMIN": 1,
}

func (x AuthRole) String() string {
	return proto.EnumName(AuthRole_name, int32(x))
}

func (AuthRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c43ab4a3390919ea, []int{0}
}

type ValidateRequest struct {
	// access token given by user
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c43ab4a3390919ea, []int{0}
}

func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (m *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(m, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type ValidateResponse struct {
	// id of a user
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// time before expiration
	ExpireAfter *duration.Duration `protobuf:"bytes,2,opt,name=expire_after,json=expireAfter,proto3" json:"expire_after,omitempty"`
	// role of a user
	Role                 AuthRole `protobuf:"varint,3,opt,name=role,proto3,enum=pb.AuthRole" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateResponse) Reset()         { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()    {}
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c43ab4a3390919ea, []int{1}
}

func (m *ValidateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateResponse.Unmarshal(m, b)
}
func (m *ValidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateResponse.Marshal(b, m, deterministic)
}
func (m *ValidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateResponse.Merge(m, src)
}
func (m *ValidateResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateResponse.Size(m)
}
func (m *ValidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateResponse proto.InternalMessageInfo

func (m *ValidateResponse) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ValidateResponse) GetExpireAfter() *duration.Duration {
	if m != nil {
		return m.ExpireAfter
	}
	return nil
}

func (m *ValidateResponse) GetRole() AuthRole {
	if m != nil {
		return m.Role
	}
	return AuthRole_USER
}

func init() {
	proto.RegisterEnum("pb.AuthRole", AuthRole_name, AuthRole_value)
	proto.RegisterType((*ValidateRequest)(nil), "pb.ValidateRequest")
	proto.RegisterType((*ValidateResponse)(nil), "pb.ValidateResponse")
}

func init() {
	proto.RegisterFile("pb/auth.proto", fileDescriptor_c43ab4a3390919ea)
}

var fileDescriptor_c43ab4a3390919ea = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4d, 0x4f, 0xf2, 0x40,
	0x14, 0x85, 0xdf, 0xe1, 0xad, 0x08, 0x17, 0x54, 0x32, 0x9a, 0x88, 0x98, 0x48, 0x65, 0x45, 0x5c,
	0x0c, 0x09, 0xea, 0x4e, 0x17, 0x24, 0xb8, 0x60, 0xa1, 0x8b, 0xf1, 0x63, 0x4b, 0xa6, 0xf4, 0x02,
	0x8d, 0x6d, 0x67, 0x9c, 0x8f, 0x44, 0x7f, 0x83, 0x7f, 0xda, 0xb4, 0x43, 0x63, 0xec, 0x72, 0xce,
	0x79, 0xee, 0xe4, 0x39, 0x70, 0xa0, 0xa2, 0x89, 0x70, 0x76, 0xcb, 0x94, 0x96, 0x56, 0xd2, 0x86,
	0x8a, 0x06, 0xc3, 0x8d, 0x94, 0x9b, 0x14, 0x27, 0x65, 0x12, 0xb9, 0xf5, 0xc4, 0x26, 0x19, 0x1a,
	0x2b, 0x32, 0xe5, 0xa1, 0xc1, 0x45, 0x1d, 0x88, 0x9d, 0x16, 0x36, 0x91, 0xf9, 0xae, 0x3f, 0xaf,
	0xf7, 0x98, 0x29, 0xfb, 0xe5, 0xcb, 0xd1, 0x0d, 0x1c, 0xbd, 0x89, 0x34, 0x89, 0x85, 0x45, 0x8e,
	0x1f, 0x0e, 0x8d, 0xa5, 0x97, 0xd0, 0x15, 0xab, 0x15, 0x1a, 0xb3, 0xb4, 0xf2, 0x1d, 0xf3, 0x3e,
	0x09, 0xc9, 0xb8, 0xcd, 0x3b, 0x3e, 0x7b, 0x29, 0xa2, 0xd1, 0x37, 0x81, 0xde, 0xef, 0x99, 0x51,
	0x32, 0x37, 0x48, 0x4f, 0x61, 0xdf, 0x19, 0xd4, 0xcb, 0x24, 0x2e, 0x4f, 0x02, 0xde, 0x2c, 0x9e,
	0x8b, 0x98, 0xde, 0x41, 0x17, 0x3f, 0x55, 0xa2, 0x71, 0x29, 0xd6, 0x16, 0x75, 0xbf, 0x11, 0x92,
	0x71, 0x67, 0x7a, 0xc6, 0xbc, 0x17, 0xab, 0xbc, 0xd8, 0x7c, 0xe7, 0xcd, 0x3b, 0x1e, 0x9f, 0x15,
	0x34, 0x0d, 0x21, 0xd0, 0x32, 0xc5, 0xfe, 0xff, 0x90, 0x8c, 0x0f, 0xa7, 0x5d, 0xa6, 0x22, 0x36,
	0x73, 0x76, 0xcb, 0x65, 0x8a, 0xbc, 0x6c, 0xae, 0x86, 0xd0, 0xaa, 0x12, 0xda, 0x82, 0xe0, 0xf5,
	0xf9, 0x81, 0xf7, 0xfe, 0xd1, 0x36, 0xec, 0xcd, 0xe6, 0x8f, 0x8b, 0xa7, 0x1e, 0x99, 0xde, 0x43,
	0x50, 0x00, 0xf4, 0x16, 0x5a, 0x95, 0x35, 0x3d, 0x2e, 0x3e, 0xaa, 0x4d, 0x1f, 0x9c, 0xfc, 0x0d,
	0xfd, 0xb0, 0xa8, 0x59, 0x1a, 0x5e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x80, 0xbc, 0x2b,
	0x9d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	// Validate validates access token, and returns info about it.
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/pb.Auth/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	// Validate validates access token, and returns info about it.
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) Validate(ctx context.Context, req *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auth/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _Auth_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/auth.proto",
}