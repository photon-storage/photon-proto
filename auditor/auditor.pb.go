// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: auditor/auditor.proto

package pba

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Information about the auditor version.
type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A string about the auditor version.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auditor_auditor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auditor_auditor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_auditor_auditor_proto_rawDescGZIP(), []int{0}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_auditor_auditor_proto protoreflect.FileDescriptor

var file_auditor_auditor_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x32, 0x50, 0x0a, 0x07, 0x41, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x45, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x90, 0x01, 0x0a, 0x1c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x6f, 0x72, 0x42, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x6f, 0x72, 0x3b, 0x70, 0x62, 0x61, 0xaa, 0x02, 0x14, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0xca,
	0x02, 0x14, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x5c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auditor_auditor_proto_rawDescOnce sync.Once
	file_auditor_auditor_proto_rawDescData = file_auditor_auditor_proto_rawDesc
)

func file_auditor_auditor_proto_rawDescGZIP() []byte {
	file_auditor_auditor_proto_rawDescOnce.Do(func() {
		file_auditor_auditor_proto_rawDescData = protoimpl.X.CompressGZIP(file_auditor_auditor_proto_rawDescData)
	})
	return file_auditor_auditor_proto_rawDescData
}

var file_auditor_auditor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_auditor_auditor_proto_goTypes = []interface{}{
	(*VersionResponse)(nil), // 0: photon.auditor.VersionResponse
	(*emptypb.Empty)(nil),   // 1: google.protobuf.Empty
}
var file_auditor_auditor_proto_depIdxs = []int32{
	1, // 0: photon.auditor.Auditor.GetVersion:input_type -> google.protobuf.Empty
	0, // 1: photon.auditor.Auditor.GetVersion:output_type -> photon.auditor.VersionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auditor_auditor_proto_init() }
func file_auditor_auditor_proto_init() {
	if File_auditor_auditor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auditor_auditor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
			RawDescriptor: file_auditor_auditor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auditor_auditor_proto_goTypes,
		DependencyIndexes: file_auditor_auditor_proto_depIdxs,
		MessageInfos:      file_auditor_auditor_proto_msgTypes,
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
	// Retrieve information about the running Auditor server.
	GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type auditorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditorClient(cc grpc.ClientConnInterface) AuditorClient {
	return &auditorClient{cc}
}

func (c *auditorClient) GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/photon.auditor.Auditor/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditorServer is the server API for Auditor service.
type AuditorServer interface {
	// Retrieve information about the running Auditor server.
	GetVersion(context.Context, *emptypb.Empty) (*VersionResponse, error)
}

// UnimplementedAuditorServer can be embedded to have forward compatible implementations.
type UnimplementedAuditorServer struct {
}

func (*UnimplementedAuditorServer) GetVersion(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterAuditorServer(s *grpc.Server, srv AuditorServer) {
	s.RegisterService(&_Auditor_serviceDesc, srv)
}

func _Auditor_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditorServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photon.auditor.Auditor/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditorServer).GetVersion(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auditor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "photon.auditor.Auditor",
	HandlerType: (*AuditorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Auditor_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auditor/auditor.proto",
}
