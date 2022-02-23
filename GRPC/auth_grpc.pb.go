// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authProto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

//go:generate mockgen -source=auth_grpc.pb.go -destination=mocks/grpc_mock.go
// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	GetUserWithRights(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*Response, error)
	BindUserAndRole(ctx context.Context, in *User, opts ...grpc.CallOption) (*ResultBinding, error)
	CheckToken(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*Result, error)
	TokenGenerationByRefresh(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*GeneratedTokens, error)
	TokenGenerationById(ctx context.Context, in *User, opts ...grpc.CallOption) (*GeneratedTokens, error)
	GetSalt(ctx context.Context, in *ReqSalt, opts ...grpc.CallOption) (*Salt, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetUserWithRights(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetUserWithRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) BindUserAndRole(ctx context.Context, in *User, opts ...grpc.CallOption) (*ResultBinding, error) {
	out := new(ResultBinding)
	err := c.cc.Invoke(ctx, "/auth.Auth/BindUserAndRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckToken(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/auth.Auth/CheckToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) TokenGenerationByRefresh(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*GeneratedTokens, error) {
	out := new(GeneratedTokens)
	err := c.cc.Invoke(ctx, "/auth.Auth/TokenGenerationByRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) TokenGenerationById(ctx context.Context, in *User, opts ...grpc.CallOption) (*GeneratedTokens, error) {
	out := new(GeneratedTokens)
	err := c.cc.Invoke(ctx, "/auth.Auth/TokenGenerationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetSalt(ctx context.Context, in *ReqSalt, opts ...grpc.CallOption) (*Salt, error) {
	out := new(Salt)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetSalt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	GetUserWithRights(context.Context, *AccessToken) (*Response, error)
	BindUserAndRole(context.Context, *User) (*ResultBinding, error)
	CheckToken(context.Context, *AccessToken) (*Result, error)
	TokenGenerationByRefresh(context.Context, *RefreshToken) (*GeneratedTokens, error)
	TokenGenerationById(context.Context, *User) (*GeneratedTokens, error)
	GetSalt(context.Context, *ReqSalt) (*Salt, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) GetUserWithRights(context.Context, *AccessToken) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserWithRights not implemented")
}
func (UnimplementedAuthServer) BindUserAndRole(context.Context, *User) (*ResultBinding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindUserAndRole not implemented")
}
func (UnimplementedAuthServer) CheckToken(context.Context, *AccessToken) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedAuthServer) TokenGenerationByRefresh(context.Context, *RefreshToken) (*GeneratedTokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenGenerationByRefresh not implemented")
}
func (UnimplementedAuthServer) TokenGenerationById(context.Context, *User) (*GeneratedTokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenGenerationById not implemented")
}
func (UnimplementedAuthServer) GetSalt(context.Context, *ReqSalt) (*Salt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSalt not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_GetUserWithRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserWithRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetUserWithRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserWithRights(ctx, req.(*AccessToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_BindUserAndRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).BindUserAndRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/BindUserAndRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).BindUserAndRole(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckToken(ctx, req.(*AccessToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_TokenGenerationByRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).TokenGenerationByRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/TokenGenerationByRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).TokenGenerationByRefresh(ctx, req.(*RefreshToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_TokenGenerationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).TokenGenerationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/TokenGenerationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).TokenGenerationById(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetSalt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSalt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetSalt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetSalt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetSalt(ctx, req.(*ReqSalt))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserWithRights",
			Handler:    _Auth_GetUserWithRights_Handler,
		},
		{
			MethodName: "BindUserAndRole",
			Handler:    _Auth_BindUserAndRole_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _Auth_CheckToken_Handler,
		},
		{
			MethodName: "TokenGenerationByRefresh",
			Handler:    _Auth_TokenGenerationByRefresh_Handler,
		},
		{
			MethodName: "TokenGenerationById",
			Handler:    _Auth_TokenGenerationById_Handler,
		},
		{
			MethodName: "GetSalt",
			Handler:    _Auth_GetSalt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
