// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: grpc-server/proto/test.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ResponseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ResponseRequest) Reset() {
	*x = ResponseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRequest) ProtoMessage() {}

func (x *ResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRequest.ProtoReflect.Descriptor instead.
func (*ResponseRequest) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_test_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_grpc_server_proto_test_proto protoreflect.FileDescriptor

var file_grpc_server_proto_test_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x51, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x41,
	0x70, 0x69, 0x12, 0x46, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x3a, 0x01, 0x2a, 0x22, 0x05, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_server_proto_test_proto_rawDescOnce sync.Once
	file_grpc_server_proto_test_proto_rawDescData = file_grpc_server_proto_test_proto_rawDesc
)

func file_grpc_server_proto_test_proto_rawDescGZIP() []byte {
	file_grpc_server_proto_test_proto_rawDescOnce.Do(func() {
		file_grpc_server_proto_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_server_proto_test_proto_rawDescData)
	})
	return file_grpc_server_proto_test_proto_rawDescData
}

var file_grpc_server_proto_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_server_proto_test_proto_goTypes = []interface{}{
	(*ResponseRequest)(nil), // 0: main.ResponseRequest
}
var file_grpc_server_proto_test_proto_depIdxs = []int32{
	0, // 0: main.TestApi.Echo:input_type -> main.ResponseRequest
	0, // 1: main.TestApi.Echo:output_type -> main.ResponseRequest
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_server_proto_test_proto_init() }
func file_grpc_server_proto_test_proto_init() {
	if File_grpc_server_proto_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_server_proto_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRequest); i {
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
			RawDescriptor: file_grpc_server_proto_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_server_proto_test_proto_goTypes,
		DependencyIndexes: file_grpc_server_proto_test_proto_depIdxs,
		MessageInfos:      file_grpc_server_proto_test_proto_msgTypes,
	}.Build()
	File_grpc_server_proto_test_proto = out.File
	file_grpc_server_proto_test_proto_rawDesc = nil
	file_grpc_server_proto_test_proto_goTypes = nil
	file_grpc_server_proto_test_proto_depIdxs = nil
}
