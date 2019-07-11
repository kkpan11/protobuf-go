// Code generated by protoc-gen-go. DO NOT EDIT.
// source: irregular/test.proto

package irregular

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

type Message struct {
	state           protoimpl.MessageState
	OptionalMessage *IrregularMessage            `protobuf:"bytes,1,opt,name=optional_message,json=optionalMessage" json:"optional_message,omitempty"`
	RepeatedMessage []*IrregularMessage          `protobuf:"bytes,2,rep,name=repeated_message,json=repeatedMessage" json:"repeated_message,omitempty"`
	RequiredMessage *IrregularMessage            `protobuf:"bytes,3,req,name=required_message,json=requiredMessage" json:"required_message,omitempty"`
	MapMessage      map[string]*IrregularMessage `protobuf:"bytes,4,rep,name=map_message,json=mapMessage" json:"map_message,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Union:
	//	*Message_OneofMessage
	Union         isMessage_Union `protobuf_oneof:"union"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message) Reset() {
	*x = Message{}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_irregular_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Type instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_irregular_test_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetOptionalMessage() *IrregularMessage {
	if x != nil {
		return x.OptionalMessage
	}
	return nil
}

func (x *Message) GetRepeatedMessage() []*IrregularMessage {
	if x != nil {
		return x.RepeatedMessage
	}
	return nil
}

func (x *Message) GetRequiredMessage() *IrregularMessage {
	if x != nil {
		return x.RequiredMessage
	}
	return nil
}

func (x *Message) GetMapMessage() map[string]*IrregularMessage {
	if x != nil {
		return x.MapMessage
	}
	return nil
}

func (m *Message) GetUnion() isMessage_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (x *Message) GetOneofMessage() *IrregularMessage {
	if x, ok := x.GetUnion().(*Message_OneofMessage); ok {
		return x.OneofMessage
	}
	return nil
}

type isMessage_Union interface {
	isMessage_Union()
}

type Message_OneofMessage struct {
	OneofMessage *IrregularMessage `protobuf:"bytes,5,opt,name=oneof_message,json=oneofMessage,oneof"`
}

func (*Message_OneofMessage) isMessage_Union() {}

var File_irregular_test_proto protoreflect.FileDescriptor

var file_irregular_test_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x1a,
	0x19, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2f, 0x69, 0x72, 0x72, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x04, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x49, 0x72, 0x72, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x10,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x2e, 0x49, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x49, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x0b, 0x6d, 0x61, 0x70, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72,
	0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x6d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x49, 0x72, 0x72,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x68, 0x0a,
	0x0f, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x49, 0x72, 0x72, 0x65,
	0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72,
}

var (
	file_irregular_test_proto_rawDescOnce sync.Once
	file_irregular_test_proto_rawDescData = file_irregular_test_proto_rawDesc
)

func file_irregular_test_proto_rawDescGZIP() []byte {
	file_irregular_test_proto_rawDescOnce.Do(func() {
		file_irregular_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_irregular_test_proto_rawDescData)
	})
	return file_irregular_test_proto_rawDescData
}

var file_irregular_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_irregular_test_proto_goTypes = []interface{}{
	(*Message)(nil),          // 0: goproto.proto.irregular.Message
	nil,                      // 1: goproto.proto.irregular.Message.MapMessageEntry
	(*IrregularMessage)(nil), // 2: goproto.proto.irregular.IrregularMessage
}
var file_irregular_test_proto_depIdxs = []int32{
	2, // goproto.proto.irregular.Message.optional_message:type_name -> goproto.proto.irregular.IrregularMessage
	2, // goproto.proto.irregular.Message.repeated_message:type_name -> goproto.proto.irregular.IrregularMessage
	2, // goproto.proto.irregular.Message.required_message:type_name -> goproto.proto.irregular.IrregularMessage
	1, // goproto.proto.irregular.Message.map_message:type_name -> goproto.proto.irregular.Message.MapMessageEntry
	2, // goproto.proto.irregular.Message.oneof_message:type_name -> goproto.proto.irregular.IrregularMessage
	2, // goproto.proto.irregular.Message.MapMessageEntry.value:type_name -> goproto.proto.irregular.IrregularMessage
	6, // starting offset of method output_type sub-list
	6, // starting offset of method input_type sub-list
	6, // starting offset of extension type_name sub-list
	6, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_irregular_test_proto_init() }
func file_irregular_test_proto_init() {
	if File_irregular_test_proto != nil {
		return
	}
	file_irregular_irregular_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_irregular_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 6:
				return &v.sizeCache
			case 7:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_irregular_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_OneofMessage)(nil),
	}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_irregular_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_irregular_test_proto_goTypes,
		DependencyIndexes: file_irregular_test_proto_depIdxs,
		MessageInfos:      file_irregular_test_proto_msgTypes,
	}.Build()
	File_irregular_test_proto = out.File
	file_irregular_test_proto_rawDesc = nil
	file_irregular_test_proto_goTypes = nil
	file_irregular_test_proto_depIdxs = nil
}
