// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth/auth.proto
	auth/auth_types.proto

It has these top-level messages:
	CreateTokenRequest
	CreateTokenResponse
	CheckTokenRequest
	CheckTokenResponse
	ExtendTokenRequest
	ExtendTokenResponse
	UpdateAccessRequest
	GetUserTokensRequest
	GetUserTokensResponse
	DeleteTokenRequest
	DeleteUserTokensRequest
	StoredToken
	AccessObject
	ResourcesAccess
	StoredTokenForUser
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import uuid "git.containerum.net/ch/grpc-proto-files/common"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateTokenRequest struct {
	UserAgent   string           `protobuf:"bytes,1,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	Fingerprint string           `protobuf:"bytes,2,opt,name=fingerprint" json:"fingerprint,omitempty"`
	UserId      *uuid.UUID       `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIp      string           `protobuf:"bytes,4,opt,name=user_ip,json=userIp" json:"user_ip,omitempty"`
	UserRole    string           `protobuf:"bytes,5,opt,name=user_role,json=userRole" json:"user_role,omitempty"`
	RwAccess    bool             `protobuf:"varint,6,opt,name=rw_access,json=rwAccess" json:"rw_access,omitempty"`
	Access      *ResourcesAccess `protobuf:"bytes,7,opt,name=access" json:"access,omitempty"`
	PartTokenId *uuid.UUID       `protobuf:"bytes,8,opt,name=part_token_id,json=partTokenId" json:"part_token_id,omitempty"`
}

func (m *CreateTokenRequest) Reset()                    { *m = CreateTokenRequest{} }
func (m *CreateTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTokenRequest) ProtoMessage()               {}
func (*CreateTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateTokenRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *CreateTokenRequest) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *CreateTokenRequest) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *CreateTokenRequest) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

func (m *CreateTokenRequest) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

func (m *CreateTokenRequest) GetRwAccess() bool {
	if m != nil {
		return m.RwAccess
	}
	return false
}

func (m *CreateTokenRequest) GetAccess() *ResourcesAccess {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *CreateTokenRequest) GetPartTokenId() *uuid.UUID {
	if m != nil {
		return m.PartTokenId
	}
	return nil
}

type CreateTokenResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *CreateTokenResponse) Reset()                    { *m = CreateTokenResponse{} }
func (m *CreateTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateTokenResponse) ProtoMessage()               {}
func (*CreateTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CreateTokenResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type CheckTokenRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	UserAgent   string `protobuf:"bytes,2,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	FingerPrint string `protobuf:"bytes,3,opt,name=finger_print,json=fingerPrint" json:"finger_print,omitempty"`
	UserIp      string `protobuf:"bytes,4,opt,name=user_ip,json=userIp" json:"user_ip,omitempty"`
}

func (m *CheckTokenRequest) Reset()                    { *m = CheckTokenRequest{} }
func (m *CheckTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckTokenRequest) ProtoMessage()               {}
func (*CheckTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CheckTokenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CheckTokenRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *CheckTokenRequest) GetFingerPrint() string {
	if m != nil {
		return m.FingerPrint
	}
	return ""
}

func (m *CheckTokenRequest) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

type CheckTokenResponse struct {
	Access      *ResourcesAccess `protobuf:"bytes,1,opt,name=access" json:"access,omitempty"`
	UserId      *uuid.UUID       `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserRole    string           `protobuf:"bytes,3,opt,name=user_role,json=userRole" json:"user_role,omitempty"`
	TokenId     *uuid.UUID       `protobuf:"bytes,4,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	PartTokenId *uuid.UUID       `protobuf:"bytes,5,opt,name=part_token_id,json=partTokenId" json:"part_token_id,omitempty"`
}

func (m *CheckTokenResponse) Reset()                    { *m = CheckTokenResponse{} }
func (m *CheckTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckTokenResponse) ProtoMessage()               {}
func (*CheckTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CheckTokenResponse) GetAccess() *ResourcesAccess {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *CheckTokenResponse) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *CheckTokenResponse) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

func (m *CheckTokenResponse) GetTokenId() *uuid.UUID {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *CheckTokenResponse) GetPartTokenId() *uuid.UUID {
	if m != nil {
		return m.PartTokenId
	}
	return nil
}

type ExtendTokenRequest struct {
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	Fingerprint  string `protobuf:"bytes,2,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *ExtendTokenRequest) Reset()                    { *m = ExtendTokenRequest{} }
func (m *ExtendTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*ExtendTokenRequest) ProtoMessage()               {}
func (*ExtendTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ExtendTokenRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *ExtendTokenRequest) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

type ExtendTokenResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *ExtendTokenResponse) Reset()                    { *m = ExtendTokenResponse{} }
func (m *ExtendTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*ExtendTokenResponse) ProtoMessage()               {}
func (*ExtendTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ExtendTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *ExtendTokenResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type UpdateAccessRequest struct {
	UserId *uuid.UUID `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *UpdateAccessRequest) Reset()                    { *m = UpdateAccessRequest{} }
func (m *UpdateAccessRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccessRequest) ProtoMessage()               {}
func (*UpdateAccessRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateAccessRequest) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

type GetUserTokensRequest struct {
	UserId *uuid.UUID `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetUserTokensRequest) Reset()                    { *m = GetUserTokensRequest{} }
func (m *GetUserTokensRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserTokensRequest) ProtoMessage()               {}
func (*GetUserTokensRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetUserTokensRequest) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

type GetUserTokensResponse struct {
	Tokens []*StoredTokenForUser `protobuf:"bytes,1,rep,name=tokens" json:"tokens,omitempty"`
}

func (m *GetUserTokensResponse) Reset()                    { *m = GetUserTokensResponse{} }
func (m *GetUserTokensResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserTokensResponse) ProtoMessage()               {}
func (*GetUserTokensResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetUserTokensResponse) GetTokens() []*StoredTokenForUser {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type DeleteTokenRequest struct {
	TokenId *uuid.UUID `protobuf:"bytes,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	UserId  *uuid.UUID `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *DeleteTokenRequest) Reset()                    { *m = DeleteTokenRequest{} }
func (m *DeleteTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTokenRequest) ProtoMessage()               {}
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteTokenRequest) GetTokenId() *uuid.UUID {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *DeleteTokenRequest) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

type DeleteUserTokensRequest struct {
	UserId *uuid.UUID `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *DeleteUserTokensRequest) Reset()                    { *m = DeleteUserTokensRequest{} }
func (m *DeleteUserTokensRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserTokensRequest) ProtoMessage()               {}
func (*DeleteUserTokensRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteUserTokensRequest) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTokenRequest)(nil), "CreateTokenRequest")
	proto.RegisterType((*CreateTokenResponse)(nil), "CreateTokenResponse")
	proto.RegisterType((*CheckTokenRequest)(nil), "CheckTokenRequest")
	proto.RegisterType((*CheckTokenResponse)(nil), "CheckTokenResponse")
	proto.RegisterType((*ExtendTokenRequest)(nil), "ExtendTokenRequest")
	proto.RegisterType((*ExtendTokenResponse)(nil), "ExtendTokenResponse")
	proto.RegisterType((*UpdateAccessRequest)(nil), "UpdateAccessRequest")
	proto.RegisterType((*GetUserTokensRequest)(nil), "GetUserTokensRequest")
	proto.RegisterType((*GetUserTokensResponse)(nil), "GetUserTokensResponse")
	proto.RegisterType((*DeleteTokenRequest)(nil), "DeleteTokenRequest")
	proto.RegisterType((*DeleteUserTokensRequest)(nil), "DeleteUserTokensRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error)
	ExtendToken(ctx context.Context, in *ExtendTokenRequest, opts ...grpc.CallOption) (*ExtendTokenResponse, error)
	UpdateAccess(ctx context.Context, in *UpdateAccessRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	GetUserTokens(ctx context.Context, in *GetUserTokensRequest, opts ...grpc.CallOption) (*GetUserTokensResponse, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	DeleteUserTokens(ctx context.Context, in *DeleteUserTokensRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := grpc.Invoke(ctx, "/Auth/CreateToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error) {
	out := new(CheckTokenResponse)
	err := grpc.Invoke(ctx, "/Auth/CheckToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ExtendToken(ctx context.Context, in *ExtendTokenRequest, opts ...grpc.CallOption) (*ExtendTokenResponse, error) {
	out := new(ExtendTokenResponse)
	err := grpc.Invoke(ctx, "/Auth/ExtendToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateAccess(ctx context.Context, in *UpdateAccessRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/Auth/UpdateAccess", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserTokens(ctx context.Context, in *GetUserTokensRequest, opts ...grpc.CallOption) (*GetUserTokensResponse, error) {
	out := new(GetUserTokensResponse)
	err := grpc.Invoke(ctx, "/Auth/GetUserTokens", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/Auth/DeleteToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteUserTokens(ctx context.Context, in *DeleteUserTokensRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/Auth/DeleteUserTokens", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	CheckToken(context.Context, *CheckTokenRequest) (*CheckTokenResponse, error)
	ExtendToken(context.Context, *ExtendTokenRequest) (*ExtendTokenResponse, error)
	UpdateAccess(context.Context, *UpdateAccessRequest) (*google_protobuf.Empty, error)
	GetUserTokens(context.Context, *GetUserTokensRequest) (*GetUserTokensResponse, error)
	DeleteToken(context.Context, *DeleteTokenRequest) (*google_protobuf.Empty, error)
	DeleteUserTokens(context.Context, *DeleteUserTokensRequest) (*google_protobuf.Empty, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckToken(ctx, req.(*CheckTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ExtendToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ExtendToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/ExtendToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ExtendToken(ctx, req.(*ExtendTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UpdateAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UpdateAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/UpdateAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UpdateAccess(ctx, req.(*UpdateAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUserTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/GetUserTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserTokens(ctx, req.(*GetUserTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteUserTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteUserTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/DeleteUserTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteUserTokens(ctx, req.(*DeleteUserTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Auth_CreateToken_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _Auth_CheckToken_Handler,
		},
		{
			MethodName: "ExtendToken",
			Handler:    _Auth_ExtendToken_Handler,
		},
		{
			MethodName: "UpdateAccess",
			Handler:    _Auth_UpdateAccess_Handler,
		},
		{
			MethodName: "GetUserTokens",
			Handler:    _Auth_GetUserTokens_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Auth_DeleteToken_Handler,
		},
		{
			MethodName: "DeleteUserTokens",
			Handler:    _Auth_DeleteUserTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

func init() { proto.RegisterFile("auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x9b, 0x76, 0xeb, 0xba, 0x93, 0x4e, 0xbf, 0xcd, 0xe9, 0xb6, 0xa8, 0xd3, 0x0f, 0x75,
	0xe1, 0xa6, 0x08, 0xe6, 0x4a, 0x43, 0x80, 0x40, 0x08, 0x31, 0xf6, 0x07, 0xed, 0x0e, 0x05, 0xca,
	0x05, 0x08, 0x55, 0x59, 0x7a, 0xda, 0x46, 0x6b, 0xe3, 0xe0, 0x38, 0x1a, 0x7b, 0x0c, 0x1e, 0x80,
	0x67, 0xe1, 0x96, 0xc7, 0x42, 0xb6, 0xb3, 0x2d, 0x69, 0x52, 0x36, 0x24, 0x6e, 0x2a, 0xe5, 0x9c,
	0xe3, 0x63, 0xfb, 0xfb, 0xf9, 0xfa, 0x14, 0xfe, 0xf3, 0x12, 0x31, 0xe9, 0xc9, 0x1f, 0x1a, 0x71,
	0x26, 0x58, 0x7b, 0xc3, 0x67, 0xb3, 0x19, 0x0b, 0x7b, 0x49, 0x12, 0x0c, 0xd3, 0xd0, 0xe6, 0x75,
	0xcd, 0x40, 0x5c, 0x46, 0x18, 0xa7, 0xe1, 0x9d, 0x31, 0x63, 0xe3, 0x29, 0xf6, 0xd4, 0xd7, 0x59,
	0x32, 0xea, 0xe1, 0x2c, 0x12, 0x97, 0x3a, 0xe9, 0xfc, 0xa8, 0x02, 0x39, 0xe4, 0xe8, 0x09, 0xfc,
	0xc0, 0xce, 0x31, 0x74, 0xf1, 0x6b, 0x82, 0xb1, 0x20, 0xff, 0x03, 0x24, 0x31, 0xf2, 0x81, 0x37,
	0xc6, 0x50, 0xd8, 0x46, 0xc7, 0xe8, 0xae, 0xba, 0xab, 0x32, 0x72, 0x20, 0x03, 0xa4, 0x03, 0xe6,
	0x28, 0x08, 0xc7, 0xc8, 0x23, 0x1e, 0x84, 0xc2, 0xae, 0xaa, 0x7c, 0x36, 0x44, 0xee, 0xc1, 0x8a,
	0x6a, 0x10, 0x0c, 0xed, 0x5a, 0xc7, 0xe8, 0x9a, 0xfb, 0xcb, 0xb4, 0xdf, 0x3f, 0x3d, 0x72, 0xeb,
	0x32, 0x7a, 0x3a, 0x24, 0xdb, 0x57, 0xf9, 0xc8, 0x5e, 0x52, 0xab, 0x75, 0x22, 0x22, 0x3b, 0xa0,
	0xf6, 0x19, 0x70, 0x36, 0x45, 0x7b, 0x59, 0xa5, 0x1a, 0x32, 0xe0, 0xb2, 0x29, 0xca, 0x24, 0xbf,
	0x18, 0x78, 0xbe, 0x8f, 0x71, 0x6c, 0xd7, 0x3b, 0x46, 0xb7, 0xe1, 0x36, 0xf8, 0xc5, 0x81, 0xfa,
	0x26, 0x5d, 0xa8, 0xa7, 0x99, 0x15, 0xb5, 0xe3, 0x3a, 0x75, 0x31, 0x66, 0x09, 0xf7, 0x31, 0xd6,
	0x15, 0x6e, 0x9a, 0x27, 0x0f, 0x60, 0x2d, 0xf2, 0xb8, 0x18, 0x08, 0x79, 0x65, 0x79, 0xc4, 0x46,
	0xf6, 0x88, 0xa6, 0xcc, 0x29, 0x35, 0x4e, 0x87, 0xce, 0x17, 0xb0, 0x72, 0xf2, 0xc4, 0x11, 0x0b,
	0x63, 0x24, 0xbb, 0xd0, 0xd4, 0xbd, 0x74, 0x8f, 0x54, 0x21, 0x53, 0xc7, 0x54, 0x29, 0xb9, 0x0f,
	0x6b, 0x1c, 0x47, 0x1c, 0xe3, 0x49, 0x5a, 0xa3, 0x55, 0x6a, 0xa6, 0x41, 0x55, 0xe4, 0x7c, 0x37,
	0x60, 0xe3, 0x70, 0x82, 0xfe, 0x79, 0x4e, 0xfd, 0x3b, 0x74, 0xcf, 0x03, 0xaa, 0xce, 0x03, 0xda,
	0x85, 0xa6, 0xa6, 0x31, 0xd0, 0x84, 0x6a, 0x59, 0x42, 0xef, 0x14, 0xa1, 0x45, 0x04, 0x9c, 0x5f,
	0x06, 0x90, 0xec, 0x99, 0xd2, 0x2b, 0xdf, 0xc8, 0x6b, 0xdc, 0x22, 0x6f, 0x86, 0x7d, 0xb5, 0x8c,
	0x7d, 0x0e, 0x71, 0x6d, 0x0e, 0x71, 0x07, 0x1a, 0xd7, 0x58, 0x96, 0xb2, 0xab, 0x57, 0x84, 0x46,
	0x52, 0xa4, 0xb7, 0xbc, 0x90, 0xde, 0x67, 0x20, 0xc7, 0xdf, 0x04, 0x86, 0xc3, 0x9c, 0xbc, 0x05,
	0x32, 0x46, 0x91, 0xcc, 0xed, 0x16, 0x97, 0xd6, 0xc8, 0x35, 0xff, 0xc7, 0xd6, 0x78, 0x02, 0x56,
	0x3f, 0x1a, 0x7a, 0x02, 0x53, 0x75, 0xd3, 0xc3, 0x67, 0xc4, 0x35, 0x4a, 0xc4, 0x75, 0x9e, 0x42,
	0xeb, 0x2d, 0x8a, 0x7e, 0x8c, 0x5c, 0xb5, 0xb9, 0xf3, 0xba, 0x23, 0xd8, 0x9c, 0x5b, 0x97, 0xde,
	0xe7, 0x21, 0xd4, 0xd5, 0x21, 0x25, 0xf7, 0x5a, 0xd7, 0xdc, 0xb7, 0xe8, 0x7b, 0xc1, 0x38, 0xea,
	0x5b, 0x9f, 0x30, 0x2e, 0x97, 0xb8, 0x69, 0x89, 0xf3, 0x11, 0xc8, 0x11, 0x4e, 0x71, 0x6e, 0x9a,
	0x64, 0x99, 0x1a, 0xa5, 0x4c, 0x6f, 0xb1, 0x8c, 0xf3, 0x1c, 0xb6, 0x75, 0xdf, 0xbf, 0xbe, 0xd8,
	0xfe, 0xcf, 0x1a, 0x2c, 0x1d, 0x24, 0x62, 0x42, 0x5e, 0x80, 0x99, 0x79, 0xca, 0xc4, 0xa2, 0xc5,
	0xb9, 0xd7, 0x6e, 0xd1, 0x92, 0xd7, 0xee, 0x54, 0xc8, 0x33, 0x80, 0x9b, 0x27, 0x41, 0x08, 0x2d,
	0xbc, 0xd9, 0xb6, 0x45, 0x8b, 0x6f, 0xc6, 0xa9, 0xc8, 0x4d, 0x33, 0x26, 0x21, 0x16, 0x2d, 0xfa,
	0xb1, 0xdd, 0xa2, 0x25, 0x3e, 0x72, 0x2a, 0xe4, 0x15, 0x34, 0xb3, 0x0e, 0x20, 0x2d, 0x5a, 0x62,
	0x88, 0xf6, 0x16, 0xd5, 0xf3, 0x9d, 0x5e, 0xcd, 0x77, 0x7a, 0x2c, 0xe7, 0xbb, 0x53, 0x21, 0xaf,
	0x61, 0x2d, 0x87, 0x94, 0x6c, 0xd2, 0x32, 0x6b, 0xb4, 0xb7, 0x68, 0x29, 0x79, 0xa7, 0x42, 0x5e,
	0x82, 0x99, 0xc1, 0x49, 0x2c, 0x5a, 0x84, 0xfb, 0x87, 0xfd, 0x4f, 0x60, 0x7d, 0x1e, 0x1a, 0xb1,
	0xe9, 0x02, 0x8e, 0x8b, 0xfb, 0xbc, 0xa1, 0x9f, 0x1e, 0x8d, 0x03, 0x41, 0x7d, 0x16, 0x0a, 0x2f,
	0x08, 0x91, 0x27, 0x33, 0x1a, 0xa2, 0xe8, 0xf9, 0x93, 0xde, 0x98, 0x47, 0xfe, 0x9e, 0xaa, 0xdf,
	0x1b, 0x05, 0x53, 0x8c, 0xd5, 0x9f, 0xdf, 0x59, 0x5d, 0x45, 0x1e, 0xff, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xb3, 0x2d, 0x12, 0x43, 0x34, 0x07, 0x00, 0x00,
}
