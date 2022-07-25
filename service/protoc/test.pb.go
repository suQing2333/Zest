// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: test.proto

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

// this is a Test
type BaseTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestString string `protobuf:"bytes,1,opt,name=testString,proto3" json:"testString,omitempty"`
	TestBool   bool   `protobuf:"varint,2,opt,name=testBool,proto3" json:"testBool,omitempty"`
	TestInt    int32  `protobuf:"varint,3,opt,name=testInt,proto3" json:"testInt,omitempty"`
}

func (x *BaseTest) Reset() {
	*x = BaseTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseTest) ProtoMessage() {}

func (x *BaseTest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseTest.ProtoReflect.Descriptor instead.
func (*BaseTest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *BaseTest) GetTestString() string {
	if x != nil {
		return x.TestString
	}
	return ""
}

func (x *BaseTest) GetTestBool() bool {
	if x != nil {
		return x.TestBool
	}
	return false
}

func (x *BaseTest) GetTestInt() int32 {
	if x != nil {
		return x.TestInt
	}
	return 0
}

type RPCTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestString string `protobuf:"bytes,1,opt,name=testString,proto3" json:"testString,omitempty"`
	TestBool   bool   `protobuf:"varint,2,opt,name=testBool,proto3" json:"testBool,omitempty"`
	TestInt    int32  `protobuf:"varint,3,opt,name=testInt,proto3" json:"testInt,omitempty"`
}

func (x *RPCTestRequest) Reset() {
	*x = RPCTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCTestRequest) ProtoMessage() {}

func (x *RPCTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCTestRequest.ProtoReflect.Descriptor instead.
func (*RPCTestRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *RPCTestRequest) GetTestString() string {
	if x != nil {
		return x.TestString
	}
	return ""
}

func (x *RPCTestRequest) GetTestBool() bool {
	if x != nil {
		return x.TestBool
	}
	return false
}

func (x *RPCTestRequest) GetTestInt() int32 {
	if x != nil {
		return x.TestInt
	}
	return 0
}

type RPCTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestString string `protobuf:"bytes,1,opt,name=testString,proto3" json:"testString,omitempty"`
	TestBool   bool   `protobuf:"varint,2,opt,name=testBool,proto3" json:"testBool,omitempty"`
	TestInt    int32  `protobuf:"varint,3,opt,name=testInt,proto3" json:"testInt,omitempty"`
}

func (x *RPCTestResponse) Reset() {
	*x = RPCTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCTestResponse) ProtoMessage() {}

func (x *RPCTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCTestResponse.ProtoReflect.Descriptor instead.
func (*RPCTestResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *RPCTestResponse) GetTestString() string {
	if x != nil {
		return x.TestString
	}
	return ""
}

func (x *RPCTestResponse) GetTestBool() bool {
	if x != nil {
		return x.TestBool
	}
	return false
}

func (x *RPCTestResponse) GetTestInt() int32 {
	if x != nil {
		return x.TestInt
	}
	return 0
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x22, 0x60, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x22, 0x66, 0x0a, 0x0e, 0x52, 0x50, 0x43, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74,
	0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74,
	0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x22, 0x67,
	0x0a, 0x0f, 0x52, 0x50, 0x43, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x74, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x32, 0x84, 0x01, 0x0a, 0x0e, 0x52, 0x50, 0x43, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x74, 0x65,
	0x73, 0x74, 0x31, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x52, 0x50, 0x43,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x52, 0x50, 0x43, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x32, 0x12, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x52, 0x50, 0x43, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x52,
	0x50, 0x43, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test_proto_goTypes = []interface{}{
	(*BaseTest)(nil),        // 0: protoc.BaseTest
	(*RPCTestRequest)(nil),  // 1: protoc.RPCTestRequest
	(*RPCTestResponse)(nil), // 2: protoc.RPCTestResponse
}
var file_test_proto_depIdxs = []int32{
	1, // 0: protoc.RPCTestService.test1:input_type -> protoc.RPCTestRequest
	1, // 1: protoc.RPCTestService.test2:input_type -> protoc.RPCTestRequest
	2, // 2: protoc.RPCTestService.test1:output_type -> protoc.RPCTestResponse
	2, // 3: protoc.RPCTestService.test2:output_type -> protoc.RPCTestResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseTest); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCTestRequest); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCTestResponse); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
