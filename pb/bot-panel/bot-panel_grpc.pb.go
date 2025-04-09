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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	GetOnline(ctx context.Context, in *BotPanelGetOnlineRequest, opts ...grpc.CallOption) (*BotPanelGetOnlineResponse, error)
	GetOrgMember(ctx context.Context, in *BotPanelGetOrgMemberRequest, opts ...grpc.CallOption) (*BotPanelGetOrgMemberResponse, error)
	RakBot(ctx context.Context, in *BotPanelRakBotRequest, opts ...grpc.CallOption) (*BotPanelRakBotResponse, error)
	PropsList(ctx context.Context, in *BotPanelPropsListRequest, opts ...grpc.CallOption) (*BotPanelPropsListResponse, error)
	PropsConfirm(ctx context.Context, in *BotPanelPropsConfirmRequest, opts ...grpc.CallOption) (*BotPanelPropsConfirmResponse, error)
	DialogSend(ctx context.Context, in *BotPanelDialogSendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetID(ctx context.Context, in *BotPanelGetIDRequest, opts ...grpc.CallOption) (*BotPanelGetIDResponse, error)
	GetOrgMembers(ctx context.Context, in *BotPanelGetOrgMembersRequest, opts ...grpc.CallOption) (*BotPanelGetOrgMembersResponse, error)
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

func (c *botPanelClient) GetOnline(ctx context.Context, in *BotPanelGetOnlineRequest, opts ...grpc.CallOption) (*BotPanelGetOnlineResponse, error) {
	out := new(BotPanelGetOnlineResponse)
	err := c.cc.Invoke(ctx, "/BotPanel/GetOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botPanelClient) GetOrgMember(ctx context.Context, in *BotPanelGetOrgMemberRequest, opts ...grpc.CallOption) (*BotPanelGetOrgMemberResponse, error) {
	out := new(BotPanelGetOrgMemberResponse)
	err := c.cc.Invoke(ctx, "/BotPanel/GetOrgMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botPanelClient) RakBot(ctx context.Context, in *BotPanelRakBotRequest, opts ...grpc.CallOption) (*BotPanelRakBotResponse, error) {
	out := new(BotPanelRakBotResponse)
	err := c.cc.Invoke(ctx, "/BotPanel/RakBot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botPanelClient) PropsList(ctx context.Context, in *BotPanelPropsListRequest, opts ...grpc.CallOption) (*BotPanelPropsListResponse, error) {
	out := new(BotPanelPropsListResponse)
	err := c.cc.Invoke(ctx, "/BotPanel/PropsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botPanelClient) PropsConfirm(ctx context.Context, in *BotPanelPropsConfirmRequest, opts ...grpc.CallOption) (*BotPanelPropsConfirmResponse, error) {
	out := new(BotPanelPropsConfirmResponse)
	err := c.cc.Invoke(ctx, "/BotPanel/PropsConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botPanelClient) DialogSend(ctx context.Context, in *BotPanelDialogSendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/BotPanel/DialogSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botPanelClient) GetID(ctx context.Context, in *BotPanelGetIDRequest, opts ...grpc.CallOption) (*BotPanelGetIDResponse, error) {
	out := new(BotPanelGetIDResponse)
	err := c.cc.Invoke(ctx, "/BotPanel/GetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botPanelClient) GetOrgMembers(ctx context.Context, in *BotPanelGetOrgMembersRequest, opts ...grpc.CallOption) (*BotPanelGetOrgMembersResponse, error) {
	out := new(BotPanelGetOrgMembersResponse)
	err := c.cc.Invoke(ctx, "/BotPanel/GetOrgMembers", in, out, opts...)
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
	GetOnline(context.Context, *BotPanelGetOnlineRequest) (*BotPanelGetOnlineResponse, error)
	GetOrgMember(context.Context, *BotPanelGetOrgMemberRequest) (*BotPanelGetOrgMemberResponse, error)
	RakBot(context.Context, *BotPanelRakBotRequest) (*BotPanelRakBotResponse, error)
	PropsList(context.Context, *BotPanelPropsListRequest) (*BotPanelPropsListResponse, error)
	PropsConfirm(context.Context, *BotPanelPropsConfirmRequest) (*BotPanelPropsConfirmResponse, error)
	DialogSend(context.Context, *BotPanelDialogSendRequest) (*emptypb.Empty, error)
	GetID(context.Context, *BotPanelGetIDRequest) (*BotPanelGetIDResponse, error)
	GetOrgMembers(context.Context, *BotPanelGetOrgMembersRequest) (*BotPanelGetOrgMembersResponse, error)
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
func (UnimplementedBotPanelServer) GetOnline(context.Context, *BotPanelGetOnlineRequest) (*BotPanelGetOnlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnline not implemented")
}
func (UnimplementedBotPanelServer) GetOrgMember(context.Context, *BotPanelGetOrgMemberRequest) (*BotPanelGetOrgMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgMember not implemented")
}
func (UnimplementedBotPanelServer) RakBot(context.Context, *BotPanelRakBotRequest) (*BotPanelRakBotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RakBot not implemented")
}
func (UnimplementedBotPanelServer) PropsList(context.Context, *BotPanelPropsListRequest) (*BotPanelPropsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PropsList not implemented")
}
func (UnimplementedBotPanelServer) PropsConfirm(context.Context, *BotPanelPropsConfirmRequest) (*BotPanelPropsConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PropsConfirm not implemented")
}
func (UnimplementedBotPanelServer) DialogSend(context.Context, *BotPanelDialogSendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DialogSend not implemented")
}
func (UnimplementedBotPanelServer) GetID(context.Context, *BotPanelGetIDRequest) (*BotPanelGetIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetID not implemented")
}
func (UnimplementedBotPanelServer) GetOrgMembers(context.Context, *BotPanelGetOrgMembersRequest) (*BotPanelGetOrgMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgMembers not implemented")
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

func _BotPanel_GetOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotPanelGetOnlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).GetOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/GetOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).GetOnline(ctx, req.(*BotPanelGetOnlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotPanel_GetOrgMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotPanelGetOrgMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).GetOrgMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/GetOrgMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).GetOrgMember(ctx, req.(*BotPanelGetOrgMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotPanel_RakBot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotPanelRakBotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).RakBot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/RakBot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).RakBot(ctx, req.(*BotPanelRakBotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotPanel_PropsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotPanelPropsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).PropsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/PropsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).PropsList(ctx, req.(*BotPanelPropsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotPanel_PropsConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotPanelPropsConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).PropsConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/PropsConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).PropsConfirm(ctx, req.(*BotPanelPropsConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotPanel_DialogSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotPanelDialogSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).DialogSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/DialogSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).DialogSend(ctx, req.(*BotPanelDialogSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotPanel_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotPanelGetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/GetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).GetID(ctx, req.(*BotPanelGetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotPanel_GetOrgMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotPanelGetOrgMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotPanelServer).GetOrgMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotPanel/GetOrgMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotPanelServer).GetOrgMembers(ctx, req.(*BotPanelGetOrgMembersRequest))
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
		{
			MethodName: "GetOnline",
			Handler:    _BotPanel_GetOnline_Handler,
		},
		{
			MethodName: "GetOrgMember",
			Handler:    _BotPanel_GetOrgMember_Handler,
		},
		{
			MethodName: "RakBot",
			Handler:    _BotPanel_RakBot_Handler,
		},
		{
			MethodName: "PropsList",
			Handler:    _BotPanel_PropsList_Handler,
		},
		{
			MethodName: "PropsConfirm",
			Handler:    _BotPanel_PropsConfirm_Handler,
		},
		{
			MethodName: "DialogSend",
			Handler:    _BotPanel_DialogSend_Handler,
		},
		{
			MethodName: "GetID",
			Handler:    _BotPanel_GetID_Handler,
		},
		{
			MethodName: "GetOrgMembers",
			Handler:    _BotPanel_GetOrgMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot-panel.proto",
}
