// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: auditor/auditor.proto

package pba

import (
	context "context"
	reflect "reflect"

	grpc "google.golang.org/grpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auditor_auditor_proto protoreflect.FileDescriptor

var file_auditor_auditor_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x32, 0x09, 0x0a, 0x07, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x6f, 0x72, 0x42, 0x90, 0x01, 0x0a, 0x1c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x42, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x3b, 0x70, 0x62, 0x61, 0xaa, 0x02, 0x14, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0xca, 0x02,
	0x14, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x5c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auditor_auditor_proto_goTypes = []interface{}{}
var file_auditor_auditor_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auditor_auditor_proto_init() }
func file_auditor_auditor_proto_init() {
	if File_auditor_auditor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auditor_auditor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auditor_auditor_proto_goTypes,
		DependencyIndexes: file_auditor_auditor_proto_depIdxs,
	}.Build()
	File_auditor_auditor_proto = out.File
	file_auditor_auditor_proto_rawDesc = nil
	file_auditor_auditor_proto_goTypes = nil
	file_auditor_auditor_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuditorClient is the client API for Auditor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuditorClient interface {
}

type auditorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditorClient(cc grpc.ClientConnInterface) AuditorClient {
	return &auditorClient{cc}
}

// AuditorServer is the server API for Auditor service.
type AuditorServer interface {
}

// UnimplementedAuditorServer can be embedded to have forward compatible implementations.
type UnimplementedAuditorServer struct {
}

func RegisterAuditorServer(s *grpc.Server, srv AuditorServer) {
	s.RegisterService(&_Auditor_serviceDesc, srv)
}

var _Auditor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "photon.auditor.Auditor",
	HandlerType: (*AuditorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "auditor/auditor.proto",
}
