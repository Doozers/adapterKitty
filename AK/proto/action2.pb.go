// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: action2.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action2Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Num  int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Action2Payload) Reset() {
	*x = Action2Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action2Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action2Payload) ProtoMessage() {}

func (x *Action2Payload) ProtoReflect() protoreflect.Message {
	mi := &file_action2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action2Payload.ProtoReflect.Descriptor instead.
func (*Action2Payload) Descriptor() ([]byte, []int) {
	return file_action2_proto_rawDescGZIP(), []int{0}
}

func (x *Action2Payload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action2Payload) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_action2_proto protoreflect.FileDescriptor

var file_action2_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x36, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x42, 0x20, 0x5a, 0x1e, 0x41, 0x4b, 0x2d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_action2_proto_rawDescOnce sync.Once
	file_action2_proto_rawDescData = file_action2_proto_rawDesc
)

func file_action2_proto_rawDescGZIP() []byte {
	file_action2_proto_rawDescOnce.Do(func() {
		file_action2_proto_rawDescData = protoimpl.X.CompressGZIP(file_action2_proto_rawDescData)
	})
	return file_action2_proto_rawDescData
}

var file_action2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_action2_proto_goTypes = []interface{}{
	(*Action2Payload)(nil), // 0: Action2Payload
}
var file_action2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_action2_proto_init() }
func file_action2_proto_init() {
	if File_action2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_action2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action2Payload); i {
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
			RawDescriptor: file_action2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_action2_proto_goTypes,
		DependencyIndexes: file_action2_proto_depIdxs,
		MessageInfos:      file_action2_proto_msgTypes,
	}.Build()
	File_action2_proto = out.File
	file_action2_proto_rawDesc = nil
	file_action2_proto_goTypes = nil
	file_action2_proto_depIdxs = nil
}