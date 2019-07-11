// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/api.proto

package apipb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sourcecontextpb "google.golang.org/protobuf/types/known/sourcecontextpb"
	typepb "google.golang.org/protobuf/types/known/typepb"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

// Api is a light-weight descriptor for an API Interface.
//
// Interfaces are also described as "protocol buffer services" in some contexts,
// such as by the "service" keyword in a .proto file, but they are different
// from API Services, which represent a concrete implementation of an interface
// as opposed to simply a description of methods and bindings. They are also
// sometimes simply referred to as "APIs" in other contexts, such as the name of
// this message itself. See https://cloud.google.com/apis/design/glossary for
// detailed terminology.
type Api struct {
	state protoimpl.MessageState
	// The fully qualified name of this interface, including package name
	// followed by the interface's simple name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The methods of this interface, in unspecified order.
	Methods []*Method `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
	// Any metadata attached to the interface.
	Options []*typepb.Option `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	// A version string for this interface. If specified, must have the form
	// `major-version.minor-version`, as in `1.10`. If the minor version is
	// omitted, it defaults to zero. If the entire version field is empty, the
	// major version is derived from the package name, as outlined below. If the
	// field is not empty, the version in the package name will be verified to be
	// consistent with what is provided here.
	//
	// The versioning schema uses [semantic
	// versioning](http://semver.org) where the major version number
	// indicates a breaking change and the minor version an additive,
	// non-breaking change. Both version numbers are signals to users
	// what to expect from different versions, and should be carefully
	// chosen based on the product plan.
	//
	// The major version is also reflected in the package name of the
	// interface, which must end in `v<major-version>`, as in
	// `google.feature.v1`. For major versions 0 and 1, the suffix can
	// be omitted. Zero major versions must only be used for
	// experimental, non-GA interfaces.
	//
	//
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Source context for the protocol buffer service represented by this
	// message.
	SourceContext *sourcecontextpb.SourceContext `protobuf:"bytes,5,opt,name=source_context,json=sourceContext,proto3" json:"source_context,omitempty"`
	// Included interfaces. See [Mixin][].
	Mixins []*Mixin `protobuf:"bytes,6,rep,name=mixins,proto3" json:"mixins,omitempty"`
	// The source syntax of the service.
	Syntax        typepb.Syntax `protobuf:"varint,7,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Api) Reset() {
	*x = Api{}
}

func (x *Api) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Api) ProtoMessage() {}

func (x *Api) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Api.ProtoReflect.Type instead.
func (*Api) Descriptor() ([]byte, []int) {
	return file_google_protobuf_api_proto_rawDescGZIP(), []int{0}
}

func (x *Api) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Api) GetMethods() []*Method {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Api) GetOptions() []*typepb.Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Api) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Api) GetSourceContext() *sourcecontextpb.SourceContext {
	if x != nil {
		return x.SourceContext
	}
	return nil
}

func (x *Api) GetMixins() []*Mixin {
	if x != nil {
		return x.Mixins
	}
	return nil
}

func (x *Api) GetSyntax() typepb.Syntax {
	if x != nil {
		return x.Syntax
	}
	return typepb.Syntax_SYNTAX_PROTO2
}

// Method represents a method of an API interface.
type Method struct {
	state protoimpl.MessageState
	// The simple name of this method.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A URL of the input message type.
	RequestTypeUrl string `protobuf:"bytes,2,opt,name=request_type_url,json=requestTypeUrl,proto3" json:"request_type_url,omitempty"`
	// If true, the request is streamed.
	RequestStreaming bool `protobuf:"varint,3,opt,name=request_streaming,json=requestStreaming,proto3" json:"request_streaming,omitempty"`
	// The URL of the output message type.
	ResponseTypeUrl string `protobuf:"bytes,4,opt,name=response_type_url,json=responseTypeUrl,proto3" json:"response_type_url,omitempty"`
	// If true, the response is streamed.
	ResponseStreaming bool `protobuf:"varint,5,opt,name=response_streaming,json=responseStreaming,proto3" json:"response_streaming,omitempty"`
	// Any metadata attached to the method.
	Options []*typepb.Option `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty"`
	// The source syntax of this method.
	Syntax        typepb.Syntax `protobuf:"varint,7,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Method) Reset() {
	*x = Method{}
}

func (x *Method) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Method) ProtoMessage() {}

func (x *Method) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Method.ProtoReflect.Type instead.
func (*Method) Descriptor() ([]byte, []int) {
	return file_google_protobuf_api_proto_rawDescGZIP(), []int{1}
}

func (x *Method) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Method) GetRequestTypeUrl() string {
	if x != nil {
		return x.RequestTypeUrl
	}
	return ""
}

func (x *Method) GetRequestStreaming() bool {
	if x != nil {
		return x.RequestStreaming
	}
	return false
}

func (x *Method) GetResponseTypeUrl() string {
	if x != nil {
		return x.ResponseTypeUrl
	}
	return ""
}

func (x *Method) GetResponseStreaming() bool {
	if x != nil {
		return x.ResponseStreaming
	}
	return false
}

func (x *Method) GetOptions() []*typepb.Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Method) GetSyntax() typepb.Syntax {
	if x != nil {
		return x.Syntax
	}
	return typepb.Syntax_SYNTAX_PROTO2
}

// Declares an API Interface to be included in this interface. The including
// interface must redeclare all the methods from the included interface, but
// documentation and options are inherited as follows:
//
// - If after comment and whitespace stripping, the documentation
//   string of the redeclared method is empty, it will be inherited
//   from the original method.
//
// - Each annotation belonging to the service config (http,
//   visibility) which is not set in the redeclared method will be
//   inherited.
//
// - If an http annotation is inherited, the path pattern will be
//   modified as follows. Any version prefix will be replaced by the
//   version of the including interface plus the [root][] path if
//   specified.
//
// Example of a simple mixin:
//
//     package google.acl.v1;
//     service AccessControl {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v1/{resource=**}:getAcl";
//       }
//     }
//
//     package google.storage.v2;
//     service Storage {
//       rpc GetAcl(GetAclRequest) returns (Acl);
//
//       // Get a data record.
//       rpc GetData(GetDataRequest) returns (Data) {
//         option (google.api.http).get = "/v2/{resource=**}";
//       }
//     }
//
// Example of a mixin configuration:
//
//     apis:
//     - name: google.storage.v2.Storage
//       mixins:
//       - name: google.acl.v1.AccessControl
//
// The mixin construct implies that all methods in `AccessControl` are
// also declared with same name and request/response types in
// `Storage`. A documentation generator or annotation processor will
// see the effective `Storage.GetAcl` method after inherting
// documentation and annotations as follows:
//
//     service Storage {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v2/{resource=**}:getAcl";
//       }
//       ...
//     }
//
// Note how the version in the path pattern changed from `v1` to `v2`.
//
// If the `root` field in the mixin is specified, it should be a
// relative path under which inherited HTTP paths are placed. Example:
//
//     apis:
//     - name: google.storage.v2.Storage
//       mixins:
//       - name: google.acl.v1.AccessControl
//         root: acls
//
// This implies the following inherited HTTP annotation:
//
//     service Storage {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v2/acls/{resource=**}:getAcl";
//       }
//       ...
//     }
type Mixin struct {
	state protoimpl.MessageState
	// The fully qualified name of the interface which is included.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If non-empty specifies a path under which inherited HTTP paths
	// are rooted.
	Root          string `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Mixin) Reset() {
	*x = Mixin{}
}

func (x *Mixin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mixin) ProtoMessage() {}

func (x *Mixin) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mixin.ProtoReflect.Type instead.
func (*Mixin) Descriptor() ([]byte, []int) {
	return file_google_protobuf_api_proto_rawDescGZIP(), []int{2}
}

func (x *Mixin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Mixin) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

var File_google_protobuf_api_proto protoreflect.FileDescriptor

var file_google_protobuf_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x24, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1,
	0x02, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x31, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x2e, 0x0a, 0x06, 0x6d, 0x69, 0x78, 0x69, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x52, 0x06, 0x6d, 0x69, 0x78, 0x69, 0x6e,
	0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x74,
	0x61, 0x78, 0x22, 0xb2, 0x02, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x52,
	0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x22, 0x2f, 0x0a, 0x05, 0x4d, 0x69, 0x78, 0x69, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x42, 0x76, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42,
	0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_protobuf_api_proto_rawDescOnce sync.Once
	file_google_protobuf_api_proto_rawDescData = file_google_protobuf_api_proto_rawDesc
)

func file_google_protobuf_api_proto_rawDescGZIP() []byte {
	file_google_protobuf_api_proto_rawDescOnce.Do(func() {
		file_google_protobuf_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_api_proto_rawDescData)
	})
	return file_google_protobuf_api_proto_rawDescData
}

var file_google_protobuf_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_protobuf_api_proto_goTypes = []interface{}{
	(*Api)(nil),                           // 0: google.protobuf.Api
	(*Method)(nil),                        // 1: google.protobuf.Method
	(*Mixin)(nil),                         // 2: google.protobuf.Mixin
	(*typepb.Option)(nil),                 // 3: google.protobuf.Option
	(*sourcecontextpb.SourceContext)(nil), // 4: google.protobuf.SourceContext
	(typepb.Syntax)(0),                    // 5: google.protobuf.Syntax
}
var file_google_protobuf_api_proto_depIdxs = []int32{
	1, // google.protobuf.Api.methods:type_name -> google.protobuf.Method
	3, // google.protobuf.Api.options:type_name -> google.protobuf.Option
	4, // google.protobuf.Api.source_context:type_name -> google.protobuf.SourceContext
	2, // google.protobuf.Api.mixins:type_name -> google.protobuf.Mixin
	5, // google.protobuf.Api.syntax:type_name -> google.protobuf.Syntax
	3, // google.protobuf.Method.options:type_name -> google.protobuf.Option
	5, // google.protobuf.Method.syntax:type_name -> google.protobuf.Syntax
	7, // starting offset of method output_type sub-list
	7, // starting offset of method input_type sub-list
	7, // starting offset of extension type_name sub-list
	7, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_google_protobuf_api_proto_init() }
func file_google_protobuf_api_proto_init() {
	if File_google_protobuf_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_protobuf_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Api); i {
			case 0:
				return &v.state
			case 8:
				return &v.sizeCache
			case 9:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_protobuf_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Method); i {
			case 0:
				return &v.state
			case 8:
				return &v.sizeCache
			case 9:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_protobuf_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mixin); i {
			case 0:
				return &v.state
			case 3:
				return &v.sizeCache
			case 4:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_google_protobuf_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_api_proto_goTypes,
		DependencyIndexes: file_google_protobuf_api_proto_depIdxs,
		MessageInfos:      file_google_protobuf_api_proto_msgTypes,
	}.Build()
	File_google_protobuf_api_proto = out.File
	file_google_protobuf_api_proto_rawDesc = nil
	file_google_protobuf_api_proto_goTypes = nil
	file_google_protobuf_api_proto_depIdxs = nil
}
