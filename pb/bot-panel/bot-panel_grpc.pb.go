// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: bot-panel.proto

package bot_panel

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BotPanelClient is the client API for BotPanel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotPanelClient interface {
	SendCommand(ctx context.Context, in *SendVKCommandRequest, opts ...grpc.CallOption) (*SendVKCommandResponse, error)
	SendCommandWithoutPermissions(ctx context.Context, in *SendVKCommandRequest, opts ...grpc.CallOption) (*SendVKCommandResponse, error)
	HasDostup(ctx context.Context, in *HasDostupRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
}

type botPanelClient struct {
	cc grpc.ClientConnInterface
}

func NewBotPanelClient(cc grpc.ClientConnInterface) BotPanelClient {
	return &botPanelClient{cc}
}

func (c *botPanelClient) SendCommand(ctx context.Context, in *SendVKCommandRequest, opts ...grpc.CallOption) (*SendVKCommandResponse, error) {
	out := new(SendVKCommandResponse)
	err := c.cc.Invoke(ctx, "/BotPanel/SendCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botPanelClient) SendCommandWithoutPermissions(ctx context.Context, in *SendVKCommandRequest, opts ...grpc.CallOption) (*SendVKCommandResponse, error) {
	out := new(SendVKCommandResponse)
	err := c.cc.Invoke(ctx, "/BotPanel/SendCommandWithoutPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botPanelClient) HasDostup(ctx context.Context, in *HasDostupRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/BotPanel/HasDostup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotPanelServer is the server API for BotPanel service.
// All implementations should embed UnimplementedBotPanelServer
// for forward compatibility
type BotPanelServer interface {
	SendCommand(context.Context, *SendVKCommandRequest) (*SendVKCommandResponse, error)
	SendCommandWithoutPermissions(context.Context, *SendVKCommandRequest) (*SendVKCommandResponse, error)
	HasDostup(context.Context, *HasDostupRequest) (*wrapperspb.BoolValue, error)
}

// UnimplementedBotPanelServer should be embedded to have forward compatible implementations.
type UnimplementedBotPanelServer struct {
}

func (UnimplementedBotPanelServer) SendCommand(context.Context, *SendVKCommandRequest) (*SendVKCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommand not implemented")
}
func (UnimplementedBotPanelServer) SendCommandWithoutPermissions(context.Context, *SendVKCommandRequest) (*SendVKCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandWithoutPermissions not implemented")
}
func (UnimplementedBotPanelServer) HasDostup(context.Context, *HasDostupRequest) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasDostup not implemented")
}

// UnsafeBotPanelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotPanelServer will
// result in compilation errors.
type UnsafeBotPanelServer interface {
	mustEmbedUnimplementedBotPanelServer()
}

func RegisterBotPanelServer(s grpc.ServiceRegistrar, srv BotPanelServer) {
	s.RegisterService(&BotPanel_ServiceDesc, srv)
}

func _BotPanel_SendCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVKCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).SendCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/SendCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).SendCommand(ctx, req.(*SendVKCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotPanel_SendCommandWithoutPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVKCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).SendCommandWithoutPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/SendCommandWithoutPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).SendCommandWithoutPermissions(ctx, req.(*SendVKCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotPanel_HasDostup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasDostupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).HasDostup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/HasDostup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).HasDostup(ctx, req.(*HasDostupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BotPanel_ServiceDesc is the grpc.ServiceDesc for BotPanel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotPanel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BotPanel",
	HandlerType: (*BotPanelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCommand",
			Handler:    _BotPanel_SendCommand_Handler,
		},
		{
			MethodName: "SendCommandWithoutPermissions",
			Handler:    _BotPanel_SendCommandWithoutPermissions_Handler,
		},
		{
			MethodName: "HasDostup",
			Handler:    _BotPanel_HasDostup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot-panel.proto",
}