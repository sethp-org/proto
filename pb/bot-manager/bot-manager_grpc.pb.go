// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: bot-manager.proto

package bot_manager

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

// BotManagerClient is the client API for BotManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotManagerClient interface {
	SendCommand(ctx context.Context, in *SendBotCommandRequest, opts ...grpc.CallOption) (*SendBotCommandResponse, error)
	SendCommandV2(ctx context.Context, in *SendBotCommandV2Request, opts ...grpc.CallOption) (*SendBotCommandV2Response, error)
	GetOnline(ctx context.Context, in *BotGetOnlineRequest, opts ...grpc.CallOption) (*BotGetOnlineResponse, error)
	GetOrgMember(ctx context.Context, in *BotGetOrgMemberRequest, opts ...grpc.CallOption) (*BotGetOrgMemberResponse, error)
	RakBot(ctx context.Context, in *BotRakBotRequest, opts ...grpc.CallOption) (*BotRakBotResponse, error)
}

type botManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewBotManagerClient(cc grpc.ClientConnInterface) BotManagerClient {
	return &botManagerClient{cc}
}

func (c *botManagerClient) SendCommand(ctx context.Context, in *SendBotCommandRequest, opts ...grpc.CallOption) (*SendBotCommandResponse, error) {
	out := new(SendBotCommandResponse)
	err := c.cc.Invoke(ctx, "/BotManager/SendCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botManagerClient) SendCommandV2(ctx context.Context, in *SendBotCommandV2Request, opts ...grpc.CallOption) (*SendBotCommandV2Response, error) {
	out := new(SendBotCommandV2Response)
	err := c.cc.Invoke(ctx, "/BotManager/SendCommandV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botManagerClient) GetOnline(ctx context.Context, in *BotGetOnlineRequest, opts ...grpc.CallOption) (*BotGetOnlineResponse, error) {
	out := new(BotGetOnlineResponse)
	err := c.cc.Invoke(ctx, "/BotManager/GetOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botManagerClient) GetOrgMember(ctx context.Context, in *BotGetOrgMemberRequest, opts ...grpc.CallOption) (*BotGetOrgMemberResponse, error) {
	out := new(BotGetOrgMemberResponse)
	err := c.cc.Invoke(ctx, "/BotManager/GetOrgMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botManagerClient) RakBot(ctx context.Context, in *BotRakBotRequest, opts ...grpc.CallOption) (*BotRakBotResponse, error) {
	out := new(BotRakBotResponse)
	err := c.cc.Invoke(ctx, "/BotManager/RakBot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotManagerServer is the server API for BotManager service.
// All implementations should embed UnimplementedBotManagerServer
// for forward compatibility
type BotManagerServer interface {
	SendCommand(context.Context, *SendBotCommandRequest) (*SendBotCommandResponse, error)
	SendCommandV2(context.Context, *SendBotCommandV2Request) (*SendBotCommandV2Response, error)
	GetOnline(context.Context, *BotGetOnlineRequest) (*BotGetOnlineResponse, error)
	GetOrgMember(context.Context, *BotGetOrgMemberRequest) (*BotGetOrgMemberResponse, error)
	RakBot(context.Context, *BotRakBotRequest) (*BotRakBotResponse, error)
}

// UnimplementedBotManagerServer should be embedded to have forward compatible implementations.
type UnimplementedBotManagerServer struct {
}

func (UnimplementedBotManagerServer) SendCommand(context.Context, *SendBotCommandRequest) (*SendBotCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommand not implemented")
}
func (UnimplementedBotManagerServer) SendCommandV2(context.Context, *SendBotCommandV2Request) (*SendBotCommandV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandV2 not implemented")
}
func (UnimplementedBotManagerServer) GetOnline(context.Context, *BotGetOnlineRequest) (*BotGetOnlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnline not implemented")
}
func (UnimplementedBotManagerServer) GetOrgMember(context.Context, *BotGetOrgMemberRequest) (*BotGetOrgMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgMember not implemented")
}
func (UnimplementedBotManagerServer) RakBot(context.Context, *BotRakBotRequest) (*BotRakBotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RakBot not implemented")
}

// UnsafeBotManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotManagerServer will
// result in compilation errors.
type UnsafeBotManagerServer interface {
	mustEmbedUnimplementedBotManagerServer()
}

func RegisterBotManagerServer(s grpc.ServiceRegistrar, srv BotManagerServer) {
	s.RegisterService(&BotManager_ServiceDesc, srv)
}

func _BotManager_SendCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendBotCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotManagerServer).SendCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotManager/SendCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotManagerServer).SendCommand(ctx, req.(*SendBotCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotManager_SendCommandV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendBotCommandV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotManagerServer).SendCommandV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotManager/SendCommandV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotManagerServer).SendCommandV2(ctx, req.(*SendBotCommandV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotManager_GetOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotGetOnlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotManagerServer).GetOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotManager/GetOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotManagerServer).GetOnline(ctx, req.(*BotGetOnlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotManager_GetOrgMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotGetOrgMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotManagerServer).GetOrgMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotManager/GetOrgMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotManagerServer).GetOrgMember(ctx, req.(*BotGetOrgMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotManager_RakBot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotRakBotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotManagerServer).RakBot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotManager/RakBot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotManagerServer).RakBot(ctx, req.(*BotRakBotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BotManager_ServiceDesc is the grpc.ServiceDesc for BotManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BotManager",
	HandlerType: (*BotManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCommand",
			Handler:    _BotManager_SendCommand_Handler,
		},
		{
			MethodName: "SendCommandV2",
			Handler:    _BotManager_SendCommandV2_Handler,
		},
		{
			MethodName: "GetOnline",
			Handler:    _BotManager_GetOnline_Handler,
		},
		{
			MethodName: "GetOrgMember",
			Handler:    _BotManager_GetOrgMember_Handler,
		},
		{
			MethodName: "RakBot",
			Handler:    _BotManager_RakBot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot-manager.proto",
}
