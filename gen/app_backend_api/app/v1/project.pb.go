// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: app_backend_api/app/v1/project.proto

package appv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v1 "github.com/obnahsgnaw/appapi/gen/app_backend_api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_app_backend_api_app_v1_project_proto protoreflect.FileDescriptor

var file_app_backend_api_app_v1_project_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x70, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x70, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x61, 0x70,
	0x70, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xea, 0x02, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb4, 0x02, 0x0a,
	0x08, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x70, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61, 0x70, 0x70, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xbf, 0x01, 0x92, 0x41, 0x9b, 0x01, 0x12, 0x0c, 0xe9, 0xa1, 0xb9, 0xe7, 0x9b, 0xae,
	0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x1a, 0x57, 0xe4, 0xb8, 0xbb, 0xe8, 0xa6, 0x81, 0xe7, 0x94,
	0xa8, 0xe4, 0xba, 0x8e, 0xe4, 0xb8, 0x80, 0xe4, 0xba, 0x9b, 0xe4, 0xb8, 0x8b, 0xe6, 0x8b, 0x89,
	0xe9, 0x80, 0x89, 0xe6, 0x8b, 0xa9, 0xef, 0xbc, 0x8c, 0xe5, 0xba, 0x94, 0xe7, 0x94, 0xa8, 0xe5,
	0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe5, 0x90, 0x8e, 0xe6, 0x89, 0x8d, 0xe6, 0x9c, 0x89, 0xe4, 0xb8,
	0x8d, 0xe7, 0x94, 0xa8, 0xe4, 0xba, 0x8e, 0xe5, 0xba, 0x94, 0xe7, 0x94, 0xa8, 0xe5, 0x88, 0x9b,
	0xe5, 0xbb, 0xba, 0xe6, 0x97, 0xb6, 0xe7, 0x9a, 0x84, 0xe9, 0x80, 0x89, 0xe6, 0x8b, 0xa9, 0x62,
	0x1c, 0x0a, 0x09, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x00, 0x0a, 0x0f, 0x0a, 0x0b,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x00, 0x6a, 0x14, 0x0a,
	0x07, 0x78, 0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x09, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0xf0, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x1a, 0x19, 0x92, 0x41, 0x16, 0x12, 0x14, 0x32, 0x2e, 0xe5, 0xba, 0x94, 0xe7,
	0x94, 0xa8, 0xe9, 0xa1, 0xb9, 0xe7, 0x9b, 0xae, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x42, 0xdb,
	0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x62, 0x6e, 0x61, 0x68, 0x73,
	0x67, 0x6e, 0x61, 0x77, 0x2f, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x61, 0x70, 0x70, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41,
	0x41, 0x58, 0xaa, 0x02, 0x14, 0x41, 0x70, 0x70, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41,
	0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x41, 0x70, 0x70, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x70, 0x70, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x20, 0x41, 0x70, 0x70, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x69,
	0x5c, 0x41, 0x70, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x41, 0x70, 0x70, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_app_backend_api_app_v1_project_proto_goTypes = []interface{}{
	(*v1.PaginateAllRequest)(nil),           // 0: app_backend_api.common.v1.PaginateAllRequest
	(*v1.PaginateStringOptionResponse)(nil), // 1: app_backend_api.common.v1.PaginateStringOptionResponse
}
var file_app_backend_api_app_v1_project_proto_depIdxs = []int32{
	0, // 0: app_backend_api.app.v1.AppProjectQueryService.Paginate:input_type -> app_backend_api.common.v1.PaginateAllRequest
	1, // 1: app_backend_api.app.v1.AppProjectQueryService.Paginate:output_type -> app_backend_api.common.v1.PaginateStringOptionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_backend_api_app_v1_project_proto_init() }
func file_app_backend_api_app_v1_project_proto_init() {
	if File_app_backend_api_app_v1_project_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_backend_api_app_v1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_backend_api_app_v1_project_proto_goTypes,
		DependencyIndexes: file_app_backend_api_app_v1_project_proto_depIdxs,
	}.Build()
	File_app_backend_api_app_v1_project_proto = out.File
	file_app_backend_api_app_v1_project_proto_rawDesc = nil
	file_app_backend_api_app_v1_project_proto_goTypes = nil
	file_app_backend_api_app_v1_project_proto_depIdxs = nil
}
