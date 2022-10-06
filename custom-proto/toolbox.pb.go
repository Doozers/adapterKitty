// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: toolbox.proto

package __

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

type OperationSign int32

const (
	Operation_PLUS     OperationSign = 0
	Operation_MINUS    OperationSign = 1
	Operation_DIVIDE   OperationSign = 2
	Operation_MULTIPLY OperationSign = 3
)

// Enum value maps for OperationSign.
var (
	OperationSign_name = map[int32]string{
		0: "PLUS",
		1: "MINUS",
		2: "DIVIDE",
		3: "MULTIPLY",
	}
	OperationSign_value = map[string]int32{
		"PLUS":     0,
		"MINUS":    1,
		"DIVIDE":   2,
		"MULTIPLY": 3,
	}
)

func (x OperationSign) Enum() *OperationSign {
	p := new(OperationSign)
	*p = x
	return p
}

func (x OperationSign) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationSign) Descriptor() protoreflect.EnumDescriptor {
	return file_toolbox_proto_enumTypes[0].Descriptor()
}

func (OperationSign) Type() protoreflect.EnumType {
	return &file_toolbox_proto_enumTypes[0]
}

func (x OperationSign) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationSign.Descriptor instead.
func (OperationSign) EnumDescriptor() ([]byte, []int) {
	return file_toolbox_proto_rawDescGZIP(), []int{0, 0}
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op OperationSign `protobuf:"varint,1,opt,name=op,proto3,enum=OperationSign" json:"op,omitempty"`
	A  int32         `protobuf:"varint,2,opt,name=a,proto3" json:"a,omitempty"`
	B  int32         `protobuf:"varint,3,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toolbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_toolbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_toolbox_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetOp() OperationSign {
	if x != nil {
		return x.Op
	}
	return Operation_PLUS
}

func (x *Operation) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Operation) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toolbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_toolbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_toolbox_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toolbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_toolbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_toolbox_proto_rawDescGZIP(), []int{2}
}

func (x *Ping) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_toolbox_proto protoreflect.FileDescriptor

var file_toolbox_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x6f, 0x6f, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7f, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x02,
	0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x0c, 0x0a,
	0x01, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x35, 0x0a, 0x04, 0x73, 0x69, 0x67,
	0x6e, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4c, 0x55, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d,
	0x49, 0x4e, 0x55, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x49, 0x56, 0x49, 0x44, 0x45,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x59, 0x10, 0x03,
	0x22, 0x20, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x20, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_toolbox_proto_rawDescOnce sync.Once
	file_toolbox_proto_rawDescData = file_toolbox_proto_rawDesc
)

func file_toolbox_proto_rawDescGZIP() []byte {
	file_toolbox_proto_rawDescOnce.Do(func() {
		file_toolbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_toolbox_proto_rawDescData)
	})
	return file_toolbox_proto_rawDescData
}

var file_toolbox_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_toolbox_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_toolbox_proto_goTypes = []interface{}{
	(OperationSign)(0), // 0: operation.sign
	(*Operation)(nil),  // 1: operation
	(*Result)(nil),     // 2: result
	(*Ping)(nil),       // 3: ping
}
var file_toolbox_proto_depIdxs = []int32{
	0, // 0: operation.op:type_name -> operation.sign
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_toolbox_proto_init() }
func file_toolbox_proto_init() {
	if File_toolbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_toolbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_toolbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_toolbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
			RawDescriptor: file_toolbox_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_toolbox_proto_goTypes,
		DependencyIndexes: file_toolbox_proto_depIdxs,
		EnumInfos:         file_toolbox_proto_enumTypes,
		MessageInfos:      file_toolbox_proto_msgTypes,
	}.Build()
	File_toolbox_proto = out.File
	file_toolbox_proto_rawDesc = nil
	file_toolbox_proto_goTypes = nil
	file_toolbox_proto_depIdxs = nil
}
