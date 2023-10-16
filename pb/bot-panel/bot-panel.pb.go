// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: bot-panel.proto

package bot_panel

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HasDostupRequest_BotPanelDostup int32

const (
	HasDostupRequest_SEND_COMMAND HasDostupRequest_BotPanelDostup = 0
)

// Enum value maps for HasDostupRequest_BotPanelDostup.
var (
	HasDostupRequest_BotPanelDostup_name = map[int32]string{
		0: "SEND_COMMAND",
	}
	HasDostupRequest_BotPanelDostup_value = map[string]int32{
		"SEND_COMMAND": 0,
	}
)

func (x HasDostupRequest_BotPanelDostup) Enum() *HasDostupRequest_BotPanelDostup {
	p := new(HasDostupRequest_BotPanelDostup)
	*p = x
	return p
}

func (x HasDostupRequest_BotPanelDostup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HasDostupRequest_BotPanelDostup) Descriptor() protoreflect.EnumDescriptor {
	return file_bot_panel_proto_enumTypes[0].Descriptor()
}

func (HasDostupRequest_BotPanelDostup) Type() protoreflect.EnumType {
	return &file_bot_panel_proto_enumTypes[0]
}

func (x HasDostupRequest_BotPanelDostup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HasDostupRequest_BotPanelDostup.Descriptor instead.
func (HasDostupRequest_BotPanelDostup) EnumDescriptor() ([]byte, []int) {
	return file_bot_panel_proto_rawDescGZIP(), []int{1, 0}
}

type SendVKCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId  int32  `protobuf:"varint,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	UserId  int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Command string `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *SendVKCommandRequest) Reset() {
	*x = SendVKCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_panel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVKCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVKCommandRequest) ProtoMessage() {}

func (x *SendVKCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_panel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVKCommandRequest.ProtoReflect.Descriptor instead.
func (*SendVKCommandRequest) Descriptor() ([]byte, []int) {
	return file_bot_panel_proto_rawDescGZIP(), []int{0}
}

func (x *SendVKCommandRequest) GetPeerId() int32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *SendVKCommandRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SendVKCommandRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type HasDostupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId int32                           `protobuf:"varint,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	UserId int32                           `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Dostup HasDostupRequest_BotPanelDostup `protobuf:"varint,3,opt,name=dostup,proto3,enum=HasDostupRequest_BotPanelDostup" json:"dostup,omitempty"`
}

func (x *HasDostupRequest) Reset() {
	*x = HasDostupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_panel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasDostupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasDostupRequest) ProtoMessage() {}

func (x *HasDostupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_panel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasDostupRequest.ProtoReflect.Descriptor instead.
func (*HasDostupRequest) Descriptor() ([]byte, []int) {
	return file_bot_panel_proto_rawDescGZIP(), []int{1}
}

func (x *HasDostupRequest) GetPeerId() int32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *HasDostupRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *HasDostupRequest) GetDostup() HasDostupRequest_BotPanelDostup {
	if x != nil {
		return x.Dostup
	}
	return HasDostupRequest_SEND_COMMAND
}

type SendVKCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendVKCommandResponse) Reset() {
	*x = SendVKCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_panel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVKCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVKCommandResponse) ProtoMessage() {}

func (x *SendVKCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_panel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVKCommandResponse.ProtoReflect.Descriptor instead.
func (*SendVKCommandResponse) Descriptor() ([]byte, []int) {
	return file_bot_panel_proto_rawDescGZIP(), []int{2}
}

type BotPanelGetOnlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int32  `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *BotPanelGetOnlineRequest) Reset() {
	*x = BotPanelGetOnlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_panel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotPanelGetOnlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotPanelGetOnlineRequest) ProtoMessage() {}

func (x *BotPanelGetOnlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_panel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotPanelGetOnlineRequest.ProtoReflect.Descriptor instead.
func (*BotPanelGetOnlineRequest) Descriptor() ([]byte, []int) {
	return file_bot_panel_proto_rawDescGZIP(), []int{3}
}

func (x *BotPanelGetOnlineRequest) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *BotPanelGetOnlineRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type BotPanelGetOnlineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	RegIp    string `protobuf:"bytes,2,opt,name=reg_ip,json=regIp,proto3" json:"reg_ip,omitempty"`
	LastIp   string `protobuf:"bytes,3,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
}

func (x *BotPanelGetOnlineResponse) Reset() {
	*x = BotPanelGetOnlineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_panel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotPanelGetOnlineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotPanelGetOnlineResponse) ProtoMessage() {}

func (x *BotPanelGetOnlineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_panel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotPanelGetOnlineResponse.ProtoReflect.Descriptor instead.
func (*BotPanelGetOnlineResponse) Descriptor() ([]byte, []int) {
	return file_bot_panel_proto_rawDescGZIP(), []int{4}
}

func (x *BotPanelGetOnlineResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *BotPanelGetOnlineResponse) GetRegIp() string {
	if x != nil {
		return x.RegIp
	}
	return ""
}

func (x *BotPanelGetOnlineResponse) GetLastIp() string {
	if x != nil {
		return x.LastIp
	}
	return ""
}

var File_bot_panel_proto protoreflect.FileDescriptor

var file_bot_panel_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x74, 0x2d, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x62, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x4b, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x48, 0x61, 0x73, 0x44, 0x6f, 0x73,
	0x74, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06,
	0x64, 0x6f, 0x73, 0x74, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x48,
	0x61, 0x73, 0x44, 0x6f, 0x73, 0x74, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x42, 0x6f, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x44, 0x6f, 0x73, 0x74, 0x75, 0x70, 0x52, 0x06,
	0x64, 0x6f, 0x73, 0x74, 0x75, 0x70, 0x22, 0x22, 0x0a, 0x0e, 0x42, 0x6f, 0x74, 0x50, 0x61, 0x6e,
	0x65, 0x6c, 0x44, 0x6f, 0x73, 0x74, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x4e, 0x44,
	0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x00, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65,
	0x6e, 0x64, 0x56, 0x4b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x18, 0x42, 0x6f, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x47,
	0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x19, 0x42, 0x6f, 0x74, 0x50,
	0x61, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x67, 0x49, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49,
	0x70, 0x32, 0x98, 0x02, 0x0a, 0x08, 0x42, 0x6f, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x12, 0x3c,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x15, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x56, 0x4b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x4b, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x1d,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x57, 0x69, 0x74, 0x68, 0x6f,
	0x75, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x56, 0x4b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x4b, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09,
	0x48, 0x61, 0x73, 0x44, 0x6f, 0x73, 0x74, 0x75, 0x70, 0x12, 0x11, 0x2e, 0x48, 0x61, 0x73, 0x44,
	0x6f, 0x73, 0x74, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x2e, 0x42, 0x6f, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x42, 0x6f, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14,
	0x62, 0x6f, 0x74, 0x2d, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2f, 0x3b, 0x62, 0x6f, 0x74, 0x5f, 0x70,
	0x61, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bot_panel_proto_rawDescOnce sync.Once
	file_bot_panel_proto_rawDescData = file_bot_panel_proto_rawDesc
)

func file_bot_panel_proto_rawDescGZIP() []byte {
	file_bot_panel_proto_rawDescOnce.Do(func() {
		file_bot_panel_proto_rawDescData = protoimpl.X.CompressGZIP(file_bot_panel_proto_rawDescData)
	})
	return file_bot_panel_proto_rawDescData
}

var file_bot_panel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bot_panel_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bot_panel_proto_goTypes = []interface{}{
	(HasDostupRequest_BotPanelDostup)(0), // 0: HasDostupRequest.BotPanelDostup
	(*SendVKCommandRequest)(nil),         // 1: SendVKCommandRequest
	(*HasDostupRequest)(nil),             // 2: HasDostupRequest
	(*SendVKCommandResponse)(nil),        // 3: SendVKCommandResponse
	(*BotPanelGetOnlineRequest)(nil),     // 4: BotPanelGetOnlineRequest
	(*BotPanelGetOnlineResponse)(nil),    // 5: BotPanelGetOnlineResponse
	(*wrapperspb.BoolValue)(nil),         // 6: google.protobuf.BoolValue
}
var file_bot_panel_proto_depIdxs = []int32{
	0, // 0: HasDostupRequest.dostup:type_name -> HasDostupRequest.BotPanelDostup
	1, // 1: BotPanel.SendCommand:input_type -> SendVKCommandRequest
	1, // 2: BotPanel.SendCommandWithoutPermissions:input_type -> SendVKCommandRequest
	2, // 3: BotPanel.HasDostup:input_type -> HasDostupRequest
	4, // 4: BotPanel.GetOnline:input_type -> BotPanelGetOnlineRequest
	3, // 5: BotPanel.SendCommand:output_type -> SendVKCommandResponse
	3, // 6: BotPanel.SendCommandWithoutPermissions:output_type -> SendVKCommandResponse
	6, // 7: BotPanel.HasDostup:output_type -> google.protobuf.BoolValue
	5, // 8: BotPanel.GetOnline:output_type -> BotPanelGetOnlineResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bot_panel_proto_init() }
func file_bot_panel_proto_init() {
	if File_bot_panel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bot_panel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVKCommandRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bot_panel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasDostupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bot_panel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVKCommandResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bot_panel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotPanelGetOnlineRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bot_panel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotPanelGetOnlineResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bot_panel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bot_panel_proto_goTypes,
		DependencyIndexes: file_bot_panel_proto_depIdxs,
		EnumInfos:         file_bot_panel_proto_enumTypes,
		MessageInfos:      file_bot_panel_proto_msgTypes,
	}.Build()
	File_bot_panel_proto = out.File
	file_bot_panel_proto_rawDesc = nil
	file_bot_panel_proto_goTypes = nil
	file_bot_panel_proto_depIdxs = nil
}
