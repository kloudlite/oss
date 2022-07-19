// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: jseval.proto

package jseval

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EvalIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Init    string                `protobuf:"bytes,1,opt,name=init,proto3" json:"init,omitempty"`
	FunName string                `protobuf:"bytes,4,opt,name=funName,proto3" json:"funName,omitempty"`
	Inputs  map[string]*anypb.Any `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EvalIn) Reset() {
	*x = EvalIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jseval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvalIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvalIn) ProtoMessage() {}

func (x *EvalIn) ProtoReflect() protoreflect.Message {
	mi := &file_jseval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvalIn.ProtoReflect.Descriptor instead.
func (*EvalIn) Descriptor() ([]byte, []int) {
	return file_jseval_proto_rawDescGZIP(), []int{0}
}

func (x *EvalIn) GetInit() string {
	if x != nil {
		return x.Init
	}
	return ""
}

func (x *EvalIn) GetFunName() string {
	if x != nil {
		return x.FunName
	}
	return ""
}

func (x *EvalIn) GetInputs() map[string]*anypb.Any {
	if x != nil {
		return x.Inputs
	}
	return nil
}

type EvalOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output map[string]*anypb.Any `protobuf:"bytes,3,rep,name=output,proto3" json:"output,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EvalOut) Reset() {
	*x = EvalOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jseval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvalOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvalOut) ProtoMessage() {}

func (x *EvalOut) ProtoReflect() protoreflect.Message {
	mi := &file_jseval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvalOut.ProtoReflect.Descriptor instead.
func (*EvalOut) Descriptor() ([]byte, []int) {
	return file_jseval_proto_rawDescGZIP(), []int{1}
}

func (x *EvalOut) GetOutput() map[string]*anypb.Any {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_jseval_proto protoreflect.FileDescriptor

var file_jseval_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6a, 0x73, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x06, 0x45, 0x76,
	0x61, 0x6c, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x75, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x75, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x49, 0x6e, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x1a,
	0x4f, 0x0a, 0x0b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x88, 0x01, 0x0a, 0x07, 0x45, 0x76, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x45,
	0x76, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x1a, 0x4f, 0x0a, 0x0b, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x23, 0x0a, 0x06, 0x4a,
	0x53, 0x45, 0x76, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x04, 0x45, 0x76, 0x61, 0x6c, 0x12, 0x07, 0x2e,
	0x45, 0x76, 0x61, 0x6c, 0x49, 0x6e, 0x1a, 0x08, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x4f, 0x75, 0x74,
	0x42, 0x19, 0x5a, 0x17, 0x6b, 0x6c, 0x6f, 0x75, 0x64, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x6f,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6a, 0x73, 0x65, 0x76, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_jseval_proto_rawDescOnce sync.Once
	file_jseval_proto_rawDescData = file_jseval_proto_rawDesc
)

func file_jseval_proto_rawDescGZIP() []byte {
	file_jseval_proto_rawDescOnce.Do(func() {
		file_jseval_proto_rawDescData = protoimpl.X.CompressGZIP(file_jseval_proto_rawDescData)
	})
	return file_jseval_proto_rawDescData
}

var file_jseval_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_jseval_proto_goTypes = []interface{}{
	(*EvalIn)(nil),    // 0: EvalIn
	(*EvalOut)(nil),   // 1: EvalOut
	nil,               // 2: EvalIn.InputsEntry
	nil,               // 3: EvalOut.OutputEntry
	(*anypb.Any)(nil), // 4: google.protobuf.Any
}
var file_jseval_proto_depIdxs = []int32{
	2, // 0: EvalIn.inputs:type_name -> EvalIn.InputsEntry
	3, // 1: EvalOut.output:type_name -> EvalOut.OutputEntry
	4, // 2: EvalIn.InputsEntry.value:type_name -> google.protobuf.Any
	4, // 3: EvalOut.OutputEntry.value:type_name -> google.protobuf.Any
	0, // 4: JSEval.Eval:input_type -> EvalIn
	1, // 5: JSEval.Eval:output_type -> EvalOut
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_jseval_proto_init() }
func file_jseval_proto_init() {
	if File_jseval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jseval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvalIn); i {
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
		file_jseval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvalOut); i {
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
			RawDescriptor: file_jseval_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jseval_proto_goTypes,
		DependencyIndexes: file_jseval_proto_depIdxs,
		MessageInfos:      file_jseval_proto_msgTypes,
	}.Build()
	File_jseval_proto = out.File
	file_jseval_proto_rawDesc = nil
	file_jseval_proto_goTypes = nil
	file_jseval_proto_depIdxs = nil
}
