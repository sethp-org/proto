// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: site-backend.proto

package site_backend

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SiteBackendClient is the client API for SiteBackend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SiteBackendClient interface {
	SendLogs(ctx context.Context, in *SendLogsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendLogsV2(ctx context.Context, in *SendLogsRequestV2, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
}

type siteBackendClient struct {
	cc grpc.ClientConnInterface
}

func NewSiteBackendClient(cc grpc.ClientConnInterface) SiteBackendClient {
	return &siteBackendClient{cc}
}

func (c *siteBackendClient) SendLogs(ctx context.Context, in *SendLogsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/SiteBackend/SendLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteBackendClient) SendLogsV2(ctx context.Context, in *SendLogsRequestV2, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/SiteBackend/SendLogsV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteBackendClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/SiteBackend/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SiteBackendServer is the server API for SiteBackend service.
// All implementations should embed UnimplementedSiteBackendServer
// for forward compatibility
type SiteBackendServer interface {
	SendLogs(context.Context, *SendLogsRequest) (*emptypb.Empty, error)
	SendLogsV2(context.Context, *SendLogsRequestV2) (*emptypb.Empty, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
}

// UnimplementedSiteBackendServer should be embedded to have forward compatible implementations.
type UnimplementedSiteBackendServer struct {
}

func (UnimplementedSiteBackendServer) SendLogs(context.Context, *SendLogsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLogs not implemented")
}
func (UnimplementedSiteBackendServer) SendLogsV2(context.Context, *SendLogsRequestV2) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLogsV2 not implemented")
}
func (UnimplementedSiteBackendServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}

// UnsafeSiteBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SiteBackendServer will
// result in compilation errors.
type UnsafeSiteBackendServer interface {
	mustEmbedUnimplementedSiteBackendServer()
}

func RegisterSiteBackendServer(s grpc.ServiceRegistrar, srv SiteBackendServer) {
	s.RegisterService(&SiteBackend_ServiceDesc, srv)
}

func _SiteBackend_SendLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteBackendServer).SendLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteBackend/SendLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteBackendServer).SendLogs(ctx, req.(*SendLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteBackend_SendLogsV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendLogsRequestV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteBackendServer).SendLogsV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteBackend/SendLogsV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteBackendServer).SendLogsV2(ctx, req.(*SendLogsRequestV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteBackend_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteBackendServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteBackend/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteBackendServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SiteBackend_ServiceDesc is the grpc.ServiceDesc for SiteBackend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SiteBackend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SiteBackend",
	HandlerType: (*SiteBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLogs",
			Handler:    _SiteBackend_SendLogs_Handler,
		},
		{
			MethodName: "SendLogsV2",
			Handler:    _SiteBackend_SendLogsV2_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _SiteBackend_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "site-backend.proto",
}
