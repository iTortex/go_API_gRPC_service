// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: modulegrpc/modulegrpc.proto

package module

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

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modulegrpc_modulegrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_modulegrpc_modulegrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_modulegrpc_modulegrpc_proto_rawDescGZIP(), []int{0}
}

func (x *URL) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ShortURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Shortname string `protobuf:"bytes,2,opt,name=shortname,proto3" json:"shortname,omitempty"`
}

func (x *ShortURL) Reset() {
	*x = ShortURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modulegrpc_modulegrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortURL) ProtoMessage() {}

func (x *ShortURL) ProtoReflect() protoreflect.Message {
	mi := &file_modulegrpc_modulegrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortURL.ProtoReflect.Descriptor instead.
func (*ShortURL) Descriptor() ([]byte, []int) {
	return file_modulegrpc_modulegrpc_proto_rawDescGZIP(), []int{1}
}

func (x *ShortURL) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShortURL) GetShortname() string {
	if x != nil {
		return x.Shortname
	}
	return ""
}

var File_modulegrpc_modulegrpc_proto protoreflect.FileDescriptor

var file_modulegrpc_modulegrpc_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x22, 0x19, 0x0a, 0x03, 0x55, 0x52, 0x4c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0x72, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x52, 0x4c, 0x1a, 0x14,
	0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x52, 0x4c, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x3b, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modulegrpc_modulegrpc_proto_rawDescOnce sync.Once
	file_modulegrpc_modulegrpc_proto_rawDescData = file_modulegrpc_modulegrpc_proto_rawDesc
)

func file_modulegrpc_modulegrpc_proto_rawDescGZIP() []byte {
	file_modulegrpc_modulegrpc_proto_rawDescOnce.Do(func() {
		file_modulegrpc_modulegrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_modulegrpc_modulegrpc_proto_rawDescData)
	})
	return file_modulegrpc_modulegrpc_proto_rawDescData
}

var file_modulegrpc_modulegrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_modulegrpc_modulegrpc_proto_goTypes = []interface{}{
	(*URL)(nil),      // 0: modulegrpc.URL
	(*ShortURL)(nil), // 1: modulegrpc.ShortURL
}
var file_modulegrpc_modulegrpc_proto_depIdxs = []int32{
	0, // 0: modulegrpc.UserManagment.Create:input_type -> modulegrpc.URL
	1, // 1: modulegrpc.UserManagment.Get:input_type -> modulegrpc.ShortURL
	1, // 2: modulegrpc.UserManagment.Create:output_type -> modulegrpc.ShortURL
	0, // 3: modulegrpc.UserManagment.Get:output_type -> modulegrpc.URL
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modulegrpc_modulegrpc_proto_init() }
func file_modulegrpc_modulegrpc_proto_init() {
	if File_modulegrpc_modulegrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modulegrpc_modulegrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
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
		file_modulegrpc_modulegrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortURL); i {
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
			RawDescriptor: file_modulegrpc_modulegrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modulegrpc_modulegrpc_proto_goTypes,
		DependencyIndexes: file_modulegrpc_modulegrpc_proto_depIdxs,
		MessageInfos:      file_modulegrpc_modulegrpc_proto_msgTypes,
	}.Build()
	File_modulegrpc_modulegrpc_proto = out.File
	file_modulegrpc_modulegrpc_proto_rawDesc = nil
	file_modulegrpc_modulegrpc_proto_goTypes = nil
	file_modulegrpc_modulegrpc_proto_depIdxs = nil
}
