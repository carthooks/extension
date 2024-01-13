// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: extensions.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CarthooksExtensionClient is the client API for CarthooksExtension service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarthooksExtensionClient interface {
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoResponse, error)
	GetCustomFieldTypes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFieldsResponse, error)
	GetI18NPack(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetI18NPackResponse, error)
	GetServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetServicesResponse, error)
	GetConnectors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetConnectorsResponse, error)
	CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error)
	GetAssets(ctx context.Context, in *GetAssetsRequest, opts ...grpc.CallOption) (*GetAssetsResponse, error)
	GetAuthMethods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAuthMethodsResponse, error)
	AuthValidate(ctx context.Context, in *AuthValidateRequest, opts ...grpc.CallOption) (*AuthValidateResponse, error)
	OnHttpRequest(ctx context.Context, in *HttpRequest, opts ...grpc.CallOption) (*HttpResponse, error)
}

type carthooksExtensionClient struct {
	cc grpc.ClientConnInterface
}

func NewCarthooksExtensionClient(cc grpc.ClientConnInterface) CarthooksExtensionClient {
	return &carthooksExtensionClient{cc}
}

func (c *carthooksExtensionClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carthooksExtensionClient) GetCustomFieldTypes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFieldsResponse, error) {
	out := new(GetFieldsResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/GetCustomFieldTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carthooksExtensionClient) GetI18NPack(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetI18NPackResponse, error) {
	out := new(GetI18NPackResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/GetI18nPack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carthooksExtensionClient) GetServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carthooksExtensionClient) GetConnectors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetConnectorsResponse, error) {
	out := new(GetConnectorsResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/GetConnectors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carthooksExtensionClient) CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error) {
	out := new(CreateConnectionResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/CreateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carthooksExtensionClient) GetAssets(ctx context.Context, in *GetAssetsRequest, opts ...grpc.CallOption) (*GetAssetsResponse, error) {
	out := new(GetAssetsResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/GetAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carthooksExtensionClient) GetAuthMethods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAuthMethodsResponse, error) {
	out := new(GetAuthMethodsResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/GetAuthMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carthooksExtensionClient) AuthValidate(ctx context.Context, in *AuthValidateRequest, opts ...grpc.CallOption) (*AuthValidateResponse, error) {
	out := new(AuthValidateResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/AuthValidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carthooksExtensionClient) OnHttpRequest(ctx context.Context, in *HttpRequest, opts ...grpc.CallOption) (*HttpResponse, error) {
	out := new(HttpResponse)
	err := c.cc.Invoke(ctx, "/service.CarthooksExtension/OnHttpRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarthooksExtensionServer is the server API for CarthooksExtension service.
// All implementations must embed UnimplementedCarthooksExtensionServer
// for forward compatibility
type CarthooksExtensionServer interface {
	Info(context.Context, *Empty) (*InfoResponse, error)
	GetCustomFieldTypes(context.Context, *Empty) (*GetFieldsResponse, error)
	GetI18NPack(context.Context, *Empty) (*GetI18NPackResponse, error)
	GetServices(context.Context, *Empty) (*GetServicesResponse, error)
	GetConnectors(context.Context, *Empty) (*GetConnectorsResponse, error)
	CreateConnection(context.Context, *CreateConnectionRequest) (*CreateConnectionResponse, error)
	GetAssets(context.Context, *GetAssetsRequest) (*GetAssetsResponse, error)
	GetAuthMethods(context.Context, *Empty) (*GetAuthMethodsResponse, error)
	AuthValidate(context.Context, *AuthValidateRequest) (*AuthValidateResponse, error)
	OnHttpRequest(context.Context, *HttpRequest) (*HttpResponse, error)
	mustEmbedUnimplementedCarthooksExtensionServer()
}

// UnimplementedCarthooksExtensionServer must be embedded to have forward compatible implementations.
type UnimplementedCarthooksExtensionServer struct {
}

func (UnimplementedCarthooksExtensionServer) Info(context.Context, *Empty) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedCarthooksExtensionServer) GetCustomFieldTypes(context.Context, *Empty) (*GetFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomFieldTypes not implemented")
}
func (UnimplementedCarthooksExtensionServer) GetI18NPack(context.Context, *Empty) (*GetI18NPackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetI18NPack not implemented")
}
func (UnimplementedCarthooksExtensionServer) GetServices(context.Context, *Empty) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedCarthooksExtensionServer) GetConnectors(context.Context, *Empty) (*GetConnectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectors not implemented")
}
func (UnimplementedCarthooksExtensionServer) CreateConnection(context.Context, *CreateConnectionRequest) (*CreateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (UnimplementedCarthooksExtensionServer) GetAssets(context.Context, *GetAssetsRequest) (*GetAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssets not implemented")
}
func (UnimplementedCarthooksExtensionServer) GetAuthMethods(context.Context, *Empty) (*GetAuthMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthMethods not implemented")
}
func (UnimplementedCarthooksExtensionServer) AuthValidate(context.Context, *AuthValidateRequest) (*AuthValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthValidate not implemented")
}
func (UnimplementedCarthooksExtensionServer) OnHttpRequest(context.Context, *HttpRequest) (*HttpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnHttpRequest not implemented")
}
func (UnimplementedCarthooksExtensionServer) mustEmbedUnimplementedCarthooksExtensionServer() {}

// UnsafeCarthooksExtensionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarthooksExtensionServer will
// result in compilation errors.
type UnsafeCarthooksExtensionServer interface {
	mustEmbedUnimplementedCarthooksExtensionServer()
}

func RegisterCarthooksExtensionServer(s grpc.ServiceRegistrar, srv CarthooksExtensionServer) {
	s.RegisterService(&CarthooksExtension_ServiceDesc, srv)
}

func _CarthooksExtension_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).Info(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarthooksExtension_GetCustomFieldTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).GetCustomFieldTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/GetCustomFieldTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).GetCustomFieldTypes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarthooksExtension_GetI18NPack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).GetI18NPack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/GetI18nPack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).GetI18NPack(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarthooksExtension_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).GetServices(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarthooksExtension_GetConnectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).GetConnectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/GetConnectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).GetConnectors(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarthooksExtension_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/CreateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).CreateConnection(ctx, req.(*CreateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarthooksExtension_GetAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).GetAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/GetAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).GetAssets(ctx, req.(*GetAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarthooksExtension_GetAuthMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).GetAuthMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/GetAuthMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).GetAuthMethods(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarthooksExtension_AuthValidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).AuthValidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/AuthValidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).AuthValidate(ctx, req.(*AuthValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarthooksExtension_OnHttpRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarthooksExtensionServer).OnHttpRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CarthooksExtension/OnHttpRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarthooksExtensionServer).OnHttpRequest(ctx, req.(*HttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CarthooksExtension_ServiceDesc is the grpc.ServiceDesc for CarthooksExtension service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarthooksExtension_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.CarthooksExtension",
	HandlerType: (*CarthooksExtensionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _CarthooksExtension_Info_Handler,
		},
		{
			MethodName: "GetCustomFieldTypes",
			Handler:    _CarthooksExtension_GetCustomFieldTypes_Handler,
		},
		{
			MethodName: "GetI18nPack",
			Handler:    _CarthooksExtension_GetI18NPack_Handler,
		},
		{
			MethodName: "GetServices",
			Handler:    _CarthooksExtension_GetServices_Handler,
		},
		{
			MethodName: "GetConnectors",
			Handler:    _CarthooksExtension_GetConnectors_Handler,
		},
		{
			MethodName: "CreateConnection",
			Handler:    _CarthooksExtension_CreateConnection_Handler,
		},
		{
			MethodName: "GetAssets",
			Handler:    _CarthooksExtension_GetAssets_Handler,
		},
		{
			MethodName: "GetAuthMethods",
			Handler:    _CarthooksExtension_GetAuthMethods_Handler,
		},
		{
			MethodName: "AuthValidate",
			Handler:    _CarthooksExtension_AuthValidate_Handler,
		},
		{
			MethodName: "OnHttpRequest",
			Handler:    _CarthooksExtension_OnHttpRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "extensions.proto",
}
