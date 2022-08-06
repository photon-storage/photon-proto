// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: consensus/finalized_block_root_container.proto

package pbc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FinalizedBlockRootContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentRoot []byte `protobuf:"bytes,1,opt,name=parent_root,json=parentRoot,proto3" json:"parent_root,omitempty"`
	ChildRoot  []byte `protobuf:"bytes,2,opt,name=child_root,json=childRoot,proto3" json:"child_root,omitempty"`
}

func (x *FinalizedBlockRootContainer) Reset() {
	*x = FinalizedBlockRootContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_finalized_block_root_container_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizedBlockRootContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizedBlockRootContainer) ProtoMessage() {}

func (x *FinalizedBlockRootContainer) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_finalized_block_root_container_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizedBlockRootContainer.ProtoReflect.Descriptor instead.
func (*FinalizedBlockRootContainer) Descriptor() ([]byte, []int) {
	return file_consensus_finalized_block_root_container_proto_rawDescGZIP(), []int{0}
}

func (x *FinalizedBlockRootContainer) GetParentRoot() []byte {
	if x != nil {
		return x.ParentRoot
	}
	return nil
}

func (x *FinalizedBlockRootContainer) GetChildRoot() []byte {
	if x != nil {
		return x.ChildRoot
	}
	return nil
}

var File_consensus_finalized_block_root_container_proto protoreflect.FileDescriptor

var file_consensus_finalized_block_root_container_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x22, 0x5d, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x6f, 0x6f,
	0x74, 0x42, 0xac, 0x01, 0x0a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x42, 0x20, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2d, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x3b, 0x70, 0x62, 0x63, 0xaa, 0x02,
	0x16, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0xca, 0x02, 0x16, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e,
	0x5c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consensus_finalized_block_root_container_proto_rawDescOnce sync.Once
	file_consensus_finalized_block_root_container_proto_rawDescData = file_consensus_finalized_block_root_container_proto_rawDesc
)

func file_consensus_finalized_block_root_container_proto_rawDescGZIP() []byte {
	file_consensus_finalized_block_root_container_proto_rawDescOnce.Do(func() {
		file_consensus_finalized_block_root_container_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_finalized_block_root_container_proto_rawDescData)
	})
	return file_consensus_finalized_block_root_container_proto_rawDescData
}

var file_consensus_finalized_block_root_container_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_consensus_finalized_block_root_container_proto_goTypes = []interface{}{
	(*FinalizedBlockRootContainer)(nil), // 0: photon.consensus.FinalizedBlockRootContainer
}
var file_consensus_finalized_block_root_container_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_consensus_finalized_block_root_container_proto_init() }
func file_consensus_finalized_block_root_container_proto_init() {
	if File_consensus_finalized_block_root_container_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consensus_finalized_block_root_container_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizedBlockRootContainer); i {
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
			RawDescriptor: file_consensus_finalized_block_root_container_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consensus_finalized_block_root_container_proto_goTypes,
		DependencyIndexes: file_consensus_finalized_block_root_container_proto_depIdxs,
		MessageInfos:      file_consensus_finalized_block_root_container_proto_msgTypes,
	}.Build()
	File_consensus_finalized_block_root_container_proto = out.File
	file_consensus_finalized_block_root_container_proto_rawDesc = nil
	file_consensus_finalized_block_root_container_proto_goTypes = nil
	file_consensus_finalized_block_root_container_proto_depIdxs = nil
}
