// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: protoc/author.proto

package blogpost

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protoc_author_proto protoreflect.FileDescriptor

var file_protoc_author_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x28, 0x0a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x05, 0x2e, 0x50, 0x6f,
	0x6e, 0x67, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x6f,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protoc_author_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: Empty
	(*Pong)(nil),  // 1: Pong
}
var file_protoc_author_proto_depIdxs = []int32{
	0, // 0: AuthorService.Ping:input_type -> Empty
	1, // 1: AuthorService.Ping:output_type -> Pong
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoc_author_proto_init() }
func file_protoc_author_proto_init() {
	if File_protoc_author_proto != nil {
		return
	}
	file_protoc_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protoc_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoc_author_proto_goTypes,
		DependencyIndexes: file_protoc_author_proto_depIdxs,
	}.Build()
	File_protoc_author_proto = out.File
	file_protoc_author_proto_rawDesc = nil
	file_protoc_author_proto_goTypes = nil
	file_protoc_author_proto_depIdxs = nil
}
