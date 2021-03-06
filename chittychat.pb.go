// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: chittychat.proto

package test

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JoinMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T      int64  `protobuf:"varint,1,opt,name=t,proto3" json:"t,omitempty"`
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *JoinMessage) Reset() {
	*x = JoinMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinMessage) ProtoMessage() {}

func (x *JoinMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinMessage.ProtoReflect.Descriptor instead.
func (*JoinMessage) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{0}
}

func (x *JoinMessage) GetT() int64 {
	if x != nil {
		return x.T
	}
	return 0
}

func (x *JoinMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

type LeaveMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T      int64  `protobuf:"varint,1,opt,name=t,proto3" json:"t,omitempty"`
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *LeaveMessage) Reset() {
	*x = LeaveMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveMessage) ProtoMessage() {}

func (x *LeaveMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveMessage.ProtoReflect.Descriptor instead.
func (*LeaveMessage) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{1}
}

func (x *LeaveMessage) GetT() int64 {
	if x != nil {
		return x.T
	}
	return 0
}

func (x *LeaveMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T       int64  `protobuf:"varint,1,opt,name=t,proto3" json:"t,omitempty"`
	Sender  string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatMessage) GetT() int64 {
	if x != nil {
		return x.T
	}
	return 0
}

func (x *ChatMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MessageAcknowledgement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T               int64  `protobuf:"varint,1,opt,name=t,proto3" json:"t,omitempty"`
	Acknowledgement string `protobuf:"bytes,2,opt,name=acknowledgement,proto3" json:"acknowledgement,omitempty"`
}

func (x *MessageAcknowledgement) Reset() {
	*x = MessageAcknowledgement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAcknowledgement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAcknowledgement) ProtoMessage() {}

func (x *MessageAcknowledgement) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAcknowledgement.ProtoReflect.Descriptor instead.
func (*MessageAcknowledgement) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{3}
}

func (x *MessageAcknowledgement) GetT() int64 {
	if x != nil {
		return x.T
	}
	return 0
}

func (x *MessageAcknowledgement) GetAcknowledgement() string {
	if x != nil {
		return x.Acknowledgement
	}
	return ""
}

type ChatHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T       int64    `protobuf:"varint,1,opt,name=t,proto3" json:"t,omitempty"`
	Sender  []string `protobuf:"bytes,2,rep,name=sender,proto3" json:"sender,omitempty"`
	Message []string `protobuf:"bytes,3,rep,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatHistory) Reset() {
	*x = ChatHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatHistory) ProtoMessage() {}

func (x *ChatHistory) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatHistory.ProtoReflect.Descriptor instead.
func (*ChatHistory) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatHistory) GetT() int64 {
	if x != nil {
		return x.T
	}
	return 0
}

func (x *ChatHistory) GetSender() []string {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *ChatHistory) GetMessage() []string {
	if x != nil {
		return x.Message
	}
	return nil
}

type Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Null) Reset() {
	*x = Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null) ProtoMessage() {}

func (x *Null) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Null.ProtoReflect.Descriptor instead.
func (*Null) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{5}
}

var File_chittychat_proto protoreflect.FileDescriptor

var file_chittychat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x22, 0x33, 0x0a, 0x0b, 0x4a, 0x6f,
	0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22,
	0x34, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x01, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41,
	0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0c,
	0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x32, 0xc5, 0x01,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x1e, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x28, 0x01, 0x12, 0x39, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x12, 0x13, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x4a, 0x6f, 0x69, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x33, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x14, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0c, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x4e,
	0x75, 0x6c, 0x6c, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chittychat_proto_rawDescOnce sync.Once
	file_chittychat_proto_rawDescData = file_chittychat_proto_rawDesc
)

func file_chittychat_proto_rawDescGZIP() []byte {
	file_chittychat_proto_rawDescOnce.Do(func() {
		file_chittychat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chittychat_proto_rawDescData)
	})
	return file_chittychat_proto_rawDescData
}

var file_chittychat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chittychat_proto_goTypes = []interface{}{
	(*JoinMessage)(nil),            // 0: GRPCex.JoinMessage
	(*LeaveMessage)(nil),           // 1: GRPCex.LeaveMessage
	(*ChatMessage)(nil),            // 2: GRPCex.ChatMessage
	(*MessageAcknowledgement)(nil), // 3: GRPCex.MessageAcknowledgement
	(*ChatHistory)(nil),            // 4: GRPCex.ChatHistory
	(*Null)(nil),                   // 5: GRPCex.Null
}
var file_chittychat_proto_depIdxs = []int32{
	2, // 0: GRPCex.ChatService.SendMessage:input_type -> GRPCex.ChatMessage
	0, // 1: GRPCex.ChatService.Subscribe:input_type -> GRPCex.JoinMessage
	1, // 2: GRPCex.ChatService.Unsubscribe:input_type -> GRPCex.LeaveMessage
	3, // 3: GRPCex.ChatService.SendMessage:output_type -> GRPCex.MessageAcknowledgement
	2, // 4: GRPCex.ChatService.Subscribe:output_type -> GRPCex.ChatMessage
	5, // 5: GRPCex.ChatService.Unsubscribe:output_type -> GRPCex.Null
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chittychat_proto_init() }
func file_chittychat_proto_init() {
	if File_chittychat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chittychat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinMessage); i {
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
		file_chittychat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveMessage); i {
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
		file_chittychat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_chittychat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAcknowledgement); i {
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
		file_chittychat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatHistory); i {
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
		file_chittychat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Null); i {
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
			RawDescriptor: file_chittychat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chittychat_proto_goTypes,
		DependencyIndexes: file_chittychat_proto_depIdxs,
		MessageInfos:      file_chittychat_proto_msgTypes,
	}.Build()
	File_chittychat_proto = out.File
	file_chittychat_proto_rawDesc = nil
	file_chittychat_proto_goTypes = nil
	file_chittychat_proto_depIdxs = nil
}
