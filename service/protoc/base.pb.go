// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: base.proto

package protoc

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

// gate接收到客户端发送的协议后会将消息封装成BaseProto
// 在其他进程处理了协议后就可以通过baseInfo返回消息到客户端
type BaseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd      int32  `protobuf:"varint,1,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	Subcmd   int32  `protobuf:"varint,2,opt,name=Subcmd,proto3" json:"Subcmd,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	BackInfo []byte `protobuf:"bytes,4,opt,name=BackInfo,proto3" json:"BackInfo,omitempty"`
}

func (x *BaseProto) Reset() {
	*x = BaseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseProto) ProtoMessage() {}

func (x *BaseProto) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseProto.ProtoReflect.Descriptor instead.
func (*BaseProto) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *BaseProto) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *BaseProto) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *BaseProto) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BaseProto) GetBackInfo() []byte {
	if x != nil {
		return x.BackInfo
	}
	return nil
}

type BackInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=Module,proto3" json:"Module,omitempty"`
	Sid    int32  `protobuf:"varint,2,opt,name=Sid,proto3" json:"Sid,omitempty"`
	ConnID int64  `protobuf:"varint,3,opt,name=ConnID,proto3" json:"ConnID,omitempty"`
}

func (x *BackInfo) Reset() {
	*x = BackInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackInfo) ProtoMessage() {}

func (x *BackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackInfo.ProtoReflect.Descriptor instead.
func (*BackInfo) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *BackInfo) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *BackInfo) GetSid() int32 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *BackInfo) GetConnID() int64 {
	if x != nil {
		return x.ConnID
	}
	return 0
}

// 客户端与服务端通信过程中发送的协议
type CSProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd    int32  `protobuf:"varint,1,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	Subcmd int32  `protobuf:"varint,2,opt,name=Subcmd,proto3" json:"Subcmd,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *CSProto) Reset() {
	*x = CSProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSProto) ProtoMessage() {}

func (x *CSProto) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSProto.ProtoReflect.Descriptor instead.
func (*CSProto) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *CSProto) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *CSProto) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *CSProto) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type RPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriService string `protobuf:"bytes,1,opt,name=OriService,proto3" json:"OriService,omitempty"`
	RPCID      int64  `protobuf:"varint,2,opt,name=RPCID,proto3" json:"RPCID,omitempty"`
	CallFunc   string `protobuf:"bytes,3,opt,name=CallFunc,proto3" json:"CallFunc,omitempty"`
	CallArgs   []byte `protobuf:"bytes,4,opt,name=CallArgs,proto3" json:"CallArgs,omitempty"`
}

func (x *RPCRequest) Reset() {
	*x = RPCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCRequest) ProtoMessage() {}

func (x *RPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCRequest.ProtoReflect.Descriptor instead.
func (*RPCRequest) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *RPCRequest) GetOriService() string {
	if x != nil {
		return x.OriService
	}
	return ""
}

func (x *RPCRequest) GetRPCID() int64 {
	if x != nil {
		return x.RPCID
	}
	return 0
}

func (x *RPCRequest) GetCallFunc() string {
	if x != nil {
		return x.CallFunc
	}
	return ""
}

func (x *RPCRequest) GetCallArgs() []byte {
	if x != nil {
		return x.CallArgs
	}
	return nil
}

type RPCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RPCID int64  `protobuf:"varint,1,opt,name=RPCID,proto3" json:"RPCID,omitempty"`
	Reply []byte `protobuf:"bytes,2,opt,name=Reply,proto3" json:"Reply,omitempty"`
}

func (x *RPCResponse) Reset() {
	*x = RPCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCResponse) ProtoMessage() {}

func (x *RPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCResponse.ProtoReflect.Descriptor instead.
func (*RPCResponse) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{4}
}

func (x *RPCResponse) GetRPCID() int64 {
	if x != nil {
		return x.RPCID
	}
	return 0
}

func (x *RPCResponse) GetReply() []byte {
	if x != nil {
		return x.Reply
	}
	return nil
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x22, 0x65, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x43, 0x6d, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x75, 0x62, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4c, 0x0a, 0x08, 0x42,
	0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x53, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x53, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x07, 0x43, 0x53, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x63, 0x6d, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x75, 0x62, 0x63, 0x6d, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x7a, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x72, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x72, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x52, 0x50, 0x43, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x52, 0x50, 0x43, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x46, 0x75,
	0x6e, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x46, 0x75,
	0x6e, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x22, 0x39,
	0x0a, 0x0b, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x52, 0x50, 0x43, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x52, 0x50,
	0x43, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_base_proto_goTypes = []interface{}{
	(*BaseProto)(nil),   // 0: protoc.BaseProto
	(*BackInfo)(nil),    // 1: protoc.BackInfo
	(*CSProto)(nil),     // 2: protoc.CSProto
	(*RPCRequest)(nil),  // 3: protoc.RPCRequest
	(*RPCResponse)(nil), // 4: protoc.RPCResponse
}
var file_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseProto); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackInfo); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSProto); i {
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
		file_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCRequest); i {
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
		file_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCResponse); i {
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
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
