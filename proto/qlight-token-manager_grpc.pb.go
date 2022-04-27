// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: qlight-token-manager.proto

package proto

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

// PluginQLightTokenRefresherClient is the client API for PluginQLightTokenRefresher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginQLightTokenRefresherClient interface {
	PluginQLightTokenManager(ctx context.Context, in *PluginQLightTokenManager_Request, opts ...grpc.CallOption) (*PluginQLightTokenManager_Response, error)
	TokenRefresh(ctx context.Context, in *TokenRefresh_Request, opts ...grpc.CallOption) (*TokenRefresh_Response, error)
}

type pluginQLightTokenRefresherClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginQLightTokenRefresherClient(cc grpc.ClientConnInterface) PluginQLightTokenRefresherClient {
	return &pluginQLightTokenRefresherClient{cc}
}

func (c *pluginQLightTokenRefresherClient) PluginQLightTokenManager(ctx context.Context, in *PluginQLightTokenManager_Request, opts ...grpc.CallOption) (*PluginQLightTokenManager_Response, error) {
	out := new(PluginQLightTokenManager_Response)
	err := c.cc.Invoke(ctx, "/proto.PluginQLightTokenRefresher/PluginQLightTokenManager", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginQLightTokenRefresherClient) TokenRefresh(ctx context.Context, in *TokenRefresh_Request, opts ...grpc.CallOption) (*TokenRefresh_Response, error) {
	out := new(TokenRefresh_Response)
	err := c.cc.Invoke(ctx, "/proto.PluginQLightTokenRefresher/TokenRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginQLightTokenRefresherServer is the server API for PluginQLightTokenRefresher service.
// All implementations must embed UnimplementedPluginQLightTokenRefresherServer
// for forward compatibility
type PluginQLightTokenRefresherServer interface {
	PluginQLightTokenManager(context.Context, *PluginQLightTokenManager_Request) (*PluginQLightTokenManager_Response, error)
	TokenRefresh(context.Context, *TokenRefresh_Request) (*TokenRefresh_Response, error)
	mustEmbedUnimplementedPluginQLightTokenRefresherServer()
}

// UnimplementedPluginQLightTokenRefresherServer must be embedded to have forward compatible implementations.
type UnimplementedPluginQLightTokenRefresherServer struct {
}

func (UnimplementedPluginQLightTokenRefresherServer) PluginQLightTokenManager(context.Context, *PluginQLightTokenManager_Request) (*PluginQLightTokenManager_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginQLightTokenManager not implemented")
}
func (UnimplementedPluginQLightTokenRefresherServer) TokenRefresh(context.Context, *TokenRefresh_Request) (*TokenRefresh_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenRefresh not implemented")
}
func (UnimplementedPluginQLightTokenRefresherServer) mustEmbedUnimplementedPluginQLightTokenRefresherServer() {
}

// UnsafePluginQLightTokenRefresherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginQLightTokenRefresherServer will
// result in compilation errors.
type UnsafePluginQLightTokenRefresherServer interface {
	mustEmbedUnimplementedPluginQLightTokenRefresherServer()
}

func RegisterPluginQLightTokenRefresherServer(s grpc.ServiceRegistrar, srv PluginQLightTokenRefresherServer) {
	s.RegisterService(&PluginQLightTokenRefresher_ServiceDesc, srv)
}

func _PluginQLightTokenRefresher_PluginQLightTokenManager_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginQLightTokenManager_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginQLightTokenRefresherServer).PluginQLightTokenManager(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PluginQLightTokenRefresher/PluginQLightTokenManager",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginQLightTokenRefresherServer).PluginQLightTokenManager(ctx, req.(*PluginQLightTokenManager_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginQLightTokenRefresher_TokenRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRefresh_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginQLightTokenRefresherServer).TokenRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PluginQLightTokenRefresher/TokenRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginQLightTokenRefresherServer).TokenRefresh(ctx, req.(*TokenRefresh_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginQLightTokenRefresher_ServiceDesc is the grpc.ServiceDesc for PluginQLightTokenRefresher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginQLightTokenRefresher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PluginQLightTokenRefresher",
	HandlerType: (*PluginQLightTokenRefresherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginQLightTokenManager",
			Handler:    _PluginQLightTokenRefresher_PluginQLightTokenManager_Handler,
		},
		{
			MethodName: "TokenRefresh",
			Handler:    _PluginQLightTokenRefresher_TokenRefresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qlight-token-manager.proto",
}
