// Code generated by protoc-gen-go. DO NOT EDIT.
// source: annotations/annotations.proto

package annotations

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AnnotationsTestEnum int32

const (
	AnnotationsTestEnum_ANNOTATIONS_TEST_ENUM_VALUE AnnotationsTestEnum = 0
)

func (e AnnotationsTestEnum) Type() protoreflect.EnumType {
	return xxx_File_annotations_annotations_proto_enumTypes[0]
}
func (e AnnotationsTestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var AnnotationsTestEnum_name = map[int32]string{
	0: "ANNOTATIONS_TEST_ENUM_VALUE",
}

var AnnotationsTestEnum_value = map[string]int32{
	"ANNOTATIONS_TEST_ENUM_VALUE": 0,
}

func (x AnnotationsTestEnum) Enum() *AnnotationsTestEnum {
	p := new(AnnotationsTestEnum)
	*p = x
	return p
}

func (x AnnotationsTestEnum) String() string {
	return proto.EnumName(AnnotationsTestEnum_name, int32(x))
}

func (x *AnnotationsTestEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AnnotationsTestEnum_value, data, "AnnotationsTestEnum")
	if err != nil {
		return err
	}
	*x = AnnotationsTestEnum(value)
	return nil
}

func (AnnotationsTestEnum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_annotations_annotations_proto_rawdesc_gzipped, []int{0}
}

type AnnotationsTestMessage struct {
	AnnotationsTestField *string  `protobuf:"bytes,1,opt,name=AnnotationsTestField" json:"AnnotationsTestField,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnnotationsTestMessage) ProtoReflect() protoreflect.Message {
	return xxx_File_annotations_annotations_proto_messageTypes[0].MessageOf(m)
}
func (m *AnnotationsTestMessage) Reset()         { *m = AnnotationsTestMessage{} }
func (m *AnnotationsTestMessage) String() string { return proto.CompactTextString(m) }
func (*AnnotationsTestMessage) ProtoMessage()    {}
func (*AnnotationsTestMessage) Descriptor() ([]byte, []int) {
	return xxx_File_annotations_annotations_proto_rawdesc_gzipped, []int{0}
}

func (m *AnnotationsTestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationsTestMessage.Unmarshal(m, b)
}
func (m *AnnotationsTestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationsTestMessage.Marshal(b, m, deterministic)
}
func (m *AnnotationsTestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationsTestMessage.Merge(m, src)
}
func (m *AnnotationsTestMessage) XXX_Size() int {
	return xxx_messageInfo_AnnotationsTestMessage.Size(m)
}
func (m *AnnotationsTestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationsTestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationsTestMessage proto.InternalMessageInfo

func (m *AnnotationsTestMessage) GetAnnotationsTestField() string {
	if m != nil && m.AnnotationsTestField != nil {
		return *m.AnnotationsTestField
	}
	return ""
}

func init() {
	proto.RegisterFile("annotations/annotations.proto", xxx_File_annotations_annotations_proto_rawdesc_gzipped)
	proto.RegisterEnum("goproto.protoc.annotations.AnnotationsTestEnum", AnnotationsTestEnum_name, AnnotationsTestEnum_value)
	proto.RegisterType((*AnnotationsTestMessage)(nil), "goproto.protoc.annotations.AnnotationsTestMessage")
}

var xxx_File_annotations_annotations_proto_rawdesc = []byte{
	// 265 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x1d, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4c, 0x0a, 0x16, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2a, 0x36, 0x0a, 0x13, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10,
	0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
}

var xxx_File_annotations_annotations_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_annotations_annotations_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_annotations_annotations_proto protoreflect.FileDescriptor

var xxx_File_annotations_annotations_proto_enumTypes = make([]protoreflect.EnumType, 1)
var xxx_File_annotations_annotations_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_annotations_annotations_proto_goTypes = []interface{}{
	(AnnotationsTestEnum)(0),       // 0: goproto.protoc.annotations.AnnotationsTestEnum
	(*AnnotationsTestMessage)(nil), // 1: goproto.protoc.annotations.AnnotationsTestMessage
}
var xxx_File_annotations_annotations_proto_depIdxs = []int32{}

func init() {
	messageTypes := make([]protoreflect.MessageType, 1)
	File_annotations_annotations_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_annotations_annotations_proto_rawdesc,
		GoTypes:            xxx_File_annotations_annotations_proto_goTypes,
		DependencyIndexes:  xxx_File_annotations_annotations_proto_depIdxs,
		EnumOutputTypes:    xxx_File_annotations_annotations_proto_enumTypes,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_annotations_annotations_proto_goTypes[1:][:1]
	for i, mt := range messageTypes {
		xxx_File_annotations_annotations_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_annotations_annotations_proto_messageTypes[i].PBType = mt
	}
	xxx_File_annotations_annotations_proto_goTypes = nil
	xxx_File_annotations_annotations_proto_depIdxs = nil
}
