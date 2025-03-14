// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/lazy/lazy_tree.proto

package lazy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type Node struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nested        *Node                  `protobuf:"bytes,99,opt,name=nested" json:"nested,omitempty"`
	Int32         *int32                 `protobuf:"varint,1,opt,name=int32" json:"int32,omitempty"`
	Int64         *int64                 `protobuf:"varint,2,opt,name=int64" json:"int64,omitempty"`
	Uint32        *uint32                `protobuf:"varint,3,opt,name=uint32" json:"uint32,omitempty"`
	Uint64        *uint64                `protobuf:"varint,4,opt,name=uint64" json:"uint64,omitempty"`
	Sint32        *int32                 `protobuf:"zigzag32,5,opt,name=sint32" json:"sint32,omitempty"`
	Sint64        *int64                 `protobuf:"zigzag64,6,opt,name=sint64" json:"sint64,omitempty"`
	Fixed32       *uint32                `protobuf:"fixed32,7,opt,name=fixed32" json:"fixed32,omitempty"`
	Fixed64       *uint64                `protobuf:"fixed64,8,opt,name=fixed64" json:"fixed64,omitempty"`
	Sfixed32      *int32                 `protobuf:"fixed32,9,opt,name=sfixed32" json:"sfixed32,omitempty"`
	Sfixed64      *int64                 `protobuf:"fixed64,10,opt,name=sfixed64" json:"sfixed64,omitempty"`
	Float         *float32               `protobuf:"fixed32,11,opt,name=float" json:"float,omitempty"`
	Double        *float64               `protobuf:"fixed64,12,opt,name=double" json:"double,omitempty"`
	Bool          *bool                  `protobuf:"varint,13,opt,name=bool" json:"bool,omitempty"`
	String_       *string                `protobuf:"bytes,14,opt,name=string" json:"string,omitempty"`
	Bytes         []byte                 `protobuf:"bytes,15,opt,name=bytes" json:"bytes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Node) Reset() {
	*x = Node{}
	mi := &file_internal_testprotos_lazy_lazy_tree_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_tree_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_tree_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetNested() *Node {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *Node) GetInt32() int32 {
	if x != nil && x.Int32 != nil {
		return *x.Int32
	}
	return 0
}

func (x *Node) GetInt64() int64 {
	if x != nil && x.Int64 != nil {
		return *x.Int64
	}
	return 0
}

func (x *Node) GetUint32() uint32 {
	if x != nil && x.Uint32 != nil {
		return *x.Uint32
	}
	return 0
}

func (x *Node) GetUint64() uint64 {
	if x != nil && x.Uint64 != nil {
		return *x.Uint64
	}
	return 0
}

func (x *Node) GetSint32() int32 {
	if x != nil && x.Sint32 != nil {
		return *x.Sint32
	}
	return 0
}

func (x *Node) GetSint64() int64 {
	if x != nil && x.Sint64 != nil {
		return *x.Sint64
	}
	return 0
}

func (x *Node) GetFixed32() uint32 {
	if x != nil && x.Fixed32 != nil {
		return *x.Fixed32
	}
	return 0
}

func (x *Node) GetFixed64() uint64 {
	if x != nil && x.Fixed64 != nil {
		return *x.Fixed64
	}
	return 0
}

func (x *Node) GetSfixed32() int32 {
	if x != nil && x.Sfixed32 != nil {
		return *x.Sfixed32
	}
	return 0
}

func (x *Node) GetSfixed64() int64 {
	if x != nil && x.Sfixed64 != nil {
		return *x.Sfixed64
	}
	return 0
}

func (x *Node) GetFloat() float32 {
	if x != nil && x.Float != nil {
		return *x.Float
	}
	return 0
}

func (x *Node) GetDouble() float64 {
	if x != nil && x.Double != nil {
		return *x.Double
	}
	return 0
}

func (x *Node) GetBool() bool {
	if x != nil && x.Bool != nil {
		return *x.Bool
	}
	return false
}

func (x *Node) GetString_() string {
	if x != nil && x.String_ != nil {
		return *x.String_
	}
	return ""
}

func (x *Node) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

var File_internal_testprotos_lazy_lazy_tree_proto protoreflect.FileDescriptor

var file_internal_testprotos_lazy_lazy_tree_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x5f,
	0x74, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x61, 0x7a, 0x79,
	0x5f, 0x74, 0x72, 0x65, 0x65, 0x22, 0x9b, 0x03, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x42,
	0x02, 0x28, 0x01, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x08, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var (
	file_internal_testprotos_lazy_lazy_tree_proto_rawDescOnce sync.Once
	file_internal_testprotos_lazy_lazy_tree_proto_rawDescData []byte
)

func file_internal_testprotos_lazy_lazy_tree_proto_rawDescGZIP() []byte {
	file_internal_testprotos_lazy_lazy_tree_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_lazy_lazy_tree_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_testprotos_lazy_lazy_tree_proto_rawDesc), len(file_internal_testprotos_lazy_lazy_tree_proto_rawDesc)))
	})
	return file_internal_testprotos_lazy_lazy_tree_proto_rawDescData
}

var file_internal_testprotos_lazy_lazy_tree_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_lazy_lazy_tree_proto_goTypes = []any{
	(*Node)(nil), // 0: lazy_tree.Node
}
var file_internal_testprotos_lazy_lazy_tree_proto_depIdxs = []int32{
	0, // 0: lazy_tree.Node.nested:type_name -> lazy_tree.Node
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_testprotos_lazy_lazy_tree_proto_init() }
func file_internal_testprotos_lazy_lazy_tree_proto_init() {
	if File_internal_testprotos_lazy_lazy_tree_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_lazy_lazy_tree_proto_rawDesc), len(file_internal_testprotos_lazy_lazy_tree_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_lazy_lazy_tree_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_lazy_lazy_tree_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_lazy_lazy_tree_proto_msgTypes,
	}.Build()
	File_internal_testprotos_lazy_lazy_tree_proto = out.File
	file_internal_testprotos_lazy_lazy_tree_proto_goTypes = nil
	file_internal_testprotos_lazy_lazy_tree_proto_depIdxs = nil
}
