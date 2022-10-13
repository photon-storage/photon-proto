// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: consensus/contract.proto

package pbc

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/photon-storage/photon-proto/ext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Storage contract status enum.
type StorageStatus int32

const (
	// INVALID prevents from defaulting to a valid status.
	StorageStatus_STORAGE_INVALID StorageStatus = 0
	// CREATED indicates a newly created storage contract.
	StorageStatus_CREATED StorageStatus = 1
	// OPEN indicates a valid and effective storage contract.
	StorageStatus_OPEN StorageStatus = 2
	// AUDIT indicates an object audit is required.
	StorageStatus_AUDIT StorageStatus = 3
)

// Enum value maps for StorageStatus.
var (
	StorageStatus_name = map[int32]string{
		0: "STORAGE_INVALID",
		1: "CREATED",
		2: "OPEN",
		3: "AUDIT",
	}
	StorageStatus_value = map[string]int32{
		"STORAGE_INVALID": 0,
		"CREATED":         1,
		"OPEN":            2,
		"AUDIT":           3,
	}
)

func (x StorageStatus) Enum() *StorageStatus {
	p := new(StorageStatus)
	*p = x
	return p
}

func (x StorageStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_consensus_contract_proto_enumTypes[0].Descriptor()
}

func (StorageStatus) Type() protoreflect.EnumType {
	return &file_consensus_contract_proto_enumTypes[0]
}

func (x StorageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageStatus.Descriptor instead.
func (StorageStatus) EnumDescriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{0}
}

// Storage contract event type.
type StorageEventType int32

const (
	// INVALID prevents from defaulting to a valid status.
	StorageEventType_STORAGE_EVENT_TYPE_INVALID StorageEventType = 0
	// AUDIT_NEW is an object audit start event.
	StorageEventType_AUDIT_NEW StorageEventType = 1
	// POR_NEW is an object PoR start event.
	StorageEventType_POR_NEW StorageEventType = 2
	// AUDIT_EXPIRATION is an object audit window expiration event.
	StorageEventType_AUDIT_EXPIRATION StorageEventType = 3
	// CONTRACT_EXPIRATION is a storage contract expiration event.
	StorageEventType_CONTRACT_EXPIRATION StorageEventType = 4
)

// Enum value maps for StorageEventType.
var (
	StorageEventType_name = map[int32]string{
		0: "STORAGE_EVENT_TYPE_INVALID",
		1: "AUDIT_NEW",
		2: "POR_NEW",
		3: "AUDIT_EXPIRATION",
		4: "CONTRACT_EXPIRATION",
	}
	StorageEventType_value = map[string]int32{
		"STORAGE_EVENT_TYPE_INVALID": 0,
		"AUDIT_NEW":                  1,
		"POR_NEW":                    2,
		"AUDIT_EXPIRATION":           3,
		"CONTRACT_EXPIRATION":        4,
	}
)

func (x StorageEventType) Enum() *StorageEventType {
	p := new(StorageEventType)
	*p = x
	return p
}

func (x StorageEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_consensus_contract_proto_enumTypes[1].Descriptor()
}

func (StorageEventType) Type() protoreflect.EnumType {
	return &file_consensus_contract_proto_enumTypes[1]
}

func (x StorageEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageEventType.Descriptor instead.
func (StorageEventType) EnumDescriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{1}
}

// StorageContract defines a storage contract between an account and
// a storoage provider.
type StorageContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The root hash of the original commit transaction associated with this
	// storage contract.
	CommitTxHash []byte `protobuf:"bytes,1,opt,name=commit_tx_hash,json=commitTxHash,proto3" json:"commit_tx_hash,omitempty" ssz-size:"32"`
	// The Merkle tree root hash of the stored original data.
	ObjectHash []byte `protobuf:"bytes,2,opt,name=object_hash,json=objectHash,proto3" json:"object_hash,omitempty" ssz-size:"32"`
	// The size of the stored original data.
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// The public key of the data owner.
	Owner []byte `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty" ssz-size:"48"`
	// The public key of the storage provider.
	Depot []byte `protobuf:"bytes,5,opt,name=depot,proto3" json:"depot,omitempty" ssz-size:"48"`
	// The p2p network discovery ID used for finding depot RPC endpoint.
	DepotDiscoveryId []byte `protobuf:"bytes,6,opt,name=depot_discovery_id,json=depotDiscoveryId,proto3" json:"depot_discovery_id,omitempty" ssz-size:"32"`
	// The effective slot of the storage contract.
	Start Slot `protobuf:"varint,7,opt,name=start,proto3" json:"start,omitempty" cast-type:"Slot"`
	// Expiring slot of the storage contract.
	// The contract expires **after** the end slot.
	End Slot `protobuf:"varint,8,opt,name=end,proto3" json:"end,omitempty" cast-type:"Slot"`
	// Total storage fee that data owner paid.
	Fee uint64 `protobuf:"varint,9,opt,name=fee,proto3" json:"fee,omitempty"`
	// Storage providers collateral deposit.
	Bond uint64 `protobuf:"varint,10,opt,name=bond,proto3" json:"bond,omitempty"`
	// Status defines status of contract. Equivalent to StorageStatus. Use
	// uint32 because SSZ does not support proto enum type.
	Status uint32 `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	// End slot of the status.
	StatusExpiration Slot `protobuf:"varint,12,opt,name=status_expiration,json=statusExpiration,proto3" json:"status_expiration,omitempty" cast-type:"Slot"`
	// The public key of the selected auditor.
	Auditor []byte `protobuf:"bytes,13,opt,name=auditor,proto3" json:"auditor,omitempty" ssz-size:"48"`
	// Audit hash of the Audit struct if non-empty.
	AuditHash []byte `protobuf:"bytes,14,opt,name=audit_hash,json=auditHash,proto3" json:"audit_hash,omitempty" ssz-size:"32"`
	// PoR hash of the PoR struct if non-empty.
	PorHash []byte `protobuf:"bytes,15,opt,name=por_hash,json=porHash,proto3" json:"por_hash,omitempty" ssz-size:"32"`
}

func (x *StorageContract) Reset() {
	*x = StorageContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageContract) ProtoMessage() {}

func (x *StorageContract) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageContract.ProtoReflect.Descriptor instead.
func (*StorageContract) Descriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{0}
}

func (x *StorageContract) GetCommitTxHash() []byte {
	if x != nil {
		return x.CommitTxHash
	}
	return nil
}

func (x *StorageContract) GetObjectHash() []byte {
	if x != nil {
		return x.ObjectHash
	}
	return nil
}

func (x *StorageContract) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *StorageContract) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *StorageContract) GetDepot() []byte {
	if x != nil {
		return x.Depot
	}
	return nil
}

func (x *StorageContract) GetDepotDiscoveryId() []byte {
	if x != nil {
		return x.DepotDiscoveryId
	}
	return nil
}

func (x *StorageContract) GetStart() Slot {
	if x != nil {
		return x.Start
	}
	return Slot(0)
}

func (x *StorageContract) GetEnd() Slot {
	if x != nil {
		return x.End
	}
	return Slot(0)
}

func (x *StorageContract) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *StorageContract) GetBond() uint64 {
	if x != nil {
		return x.Bond
	}
	return 0
}

func (x *StorageContract) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StorageContract) GetStatusExpiration() Slot {
	if x != nil {
		return x.StatusExpiration
	}
	return Slot(0)
}

func (x *StorageContract) GetAuditor() []byte {
	if x != nil {
		return x.Auditor
	}
	return nil
}

func (x *StorageContract) GetAuditHash() []byte {
	if x != nil {
		return x.AuditHash
	}
	return nil
}

func (x *StorageContract) GetPorHash() []byte {
	if x != nil {
		return x.PorHash
	}
	return nil
}

// Audit defines information for encoded object data generated by audit.
type Audit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Audit start slot.
	Start Slot `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty" cast-type:"Slot"`
	// Hash of encoded data generated by audit.
	EncodedHash []byte `protobuf:"bytes,2,opt,name=encoded_hash,json=encodedHash,proto3" json:"encoded_hash,omitempty" ssz-size:"32"`
	// Size of encoded data generated by audit.
	EncodedSize uint64 `protobuf:"varint,3,opt,name=encoded_size,json=encodedSize,proto3" json:"encoded_size,omitempty"`
	// Number of blocks of encoded data.
	NumBlocks uint32 `protobuf:"varint,4,opt,name=num_blocks,json=numBlocks,proto3" json:"num_blocks,omitempty"`
	// An array of random elements used to generated block signature.
	// The count is equal to sectors_per_block.
	Rands [][]byte `protobuf:"bytes,5,rep,name=rands,proto3" json:"rands,omitempty" ssz-max:"64,96"`
}

func (x *Audit) Reset() {
	*x = Audit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audit) ProtoMessage() {}

func (x *Audit) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audit.ProtoReflect.Descriptor instead.
func (*Audit) Descriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{1}
}

func (x *Audit) GetStart() Slot {
	if x != nil {
		return x.Start
	}
	return Slot(0)
}

func (x *Audit) GetEncodedHash() []byte {
	if x != nil {
		return x.EncodedHash
	}
	return nil
}

func (x *Audit) GetEncodedSize() uint64 {
	if x != nil {
		return x.EncodedSize
	}
	return 0
}

func (x *Audit) GetNumBlocks() uint32 {
	if x != nil {
		return x.NumBlocks
	}
	return 0
}

func (x *Audit) GetRands() [][]byte {
	if x != nil {
		return x.Rands
	}
	return nil
}

type Challenge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Block index challenged.
	BlockIndex uint32 `protobuf:"varint,1,opt,name=block_index,json=blockIndex,proto3" json:"block_index,omitempty"`
	// Coefficient for the block challenge.
	Coefficient []byte `protobuf:"bytes,2,opt,name=coefficient,proto3" json:"coefficient,omitempty" ssz-size:"32"`
}

func (x *Challenge) Reset() {
	*x = Challenge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Challenge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Challenge) ProtoMessage() {}

func (x *Challenge) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Challenge.ProtoReflect.Descriptor instead.
func (*Challenge) Descriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{2}
}

func (x *Challenge) GetBlockIndex() uint32 {
	if x != nil {
		return x.BlockIndex
	}
	return 0
}

func (x *Challenge) GetCoefficient() []byte {
	if x != nil {
		return x.Coefficient
	}
	return nil
}

type PoR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PoR start slot.
	Start Slot `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty" cast-type:"Slot"`
	// The Merkle tree root hash of the stored original data.
	ObjectHash []byte `protobuf:"bytes,2,opt,name=object_hash,json=objectHash,proto3" json:"object_hash,omitempty" ssz-size:"32"`
	// Audit hash used to retrieve the particular version of audit when the
	// PoR is initiated.
	AuditHash []byte `protobuf:"bytes,3,opt,name=audit_hash,json=auditHash,proto3" json:"audit_hash,omitempty" ssz-size:"32"`
	// The public key of the last auditor.
	Auditor []byte `protobuf:"bytes,4,opt,name=auditor,proto3" json:"auditor,omitempty" ssz-size:"48"`
	// PoR challenges.
	Challenges []*Challenge `protobuf:"bytes,5,rep,name=challenges,proto3" json:"challenges,omitempty" ssz-max:"32"`
}

func (x *PoR) Reset() {
	*x = PoR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoR) ProtoMessage() {}

func (x *PoR) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoR.ProtoReflect.Descriptor instead.
func (*PoR) Descriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{3}
}

func (x *PoR) GetStart() Slot {
	if x != nil {
		return x.Start
	}
	return Slot(0)
}

func (x *PoR) GetObjectHash() []byte {
	if x != nil {
		return x.ObjectHash
	}
	return nil
}

func (x *PoR) GetAuditHash() []byte {
	if x != nil {
		return x.AuditHash
	}
	return nil
}

func (x *PoR) GetAuditor() []byte {
	if x != nil {
		return x.Auditor
	}
	return nil
}

func (x *PoR) GetChallenges() []*Challenge {
	if x != nil {
		return x.Challenges
	}
	return nil
}

type EpochStorageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*StorageEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EpochStorageEvent) Reset() {
	*x = EpochStorageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochStorageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochStorageEvent) ProtoMessage() {}

func (x *EpochStorageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochStorageEvent.ProtoReflect.Descriptor instead.
func (*EpochStorageEvent) Descriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{4}
}

func (x *EpochStorageEvent) GetEvents() []*StorageEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type StorageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event type.
	Type StorageEventType `protobuf:"varint,1,opt,name=type,proto3,enum=photon.consensus.StorageEventType" json:"type,omitempty"`
	// The root hash of the transaction (storage contract) associated with
	// this event.
	TxHash []byte `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" ssz-size:"32"`
	// The public key of the depot specified in the storage contract.
	Depot []byte `protobuf:"bytes,3,opt,name=depot,proto3" json:"depot,omitempty" ssz-size:"48"`
	// The public key of the auditor assigned for auditing task.
	Auditor []byte `protobuf:"bytes,4,opt,name=auditor,proto3" json:"auditor,omitempty" ssz-size:"48"`
}

func (x *StorageEvent) Reset() {
	*x = StorageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageEvent) ProtoMessage() {}

func (x *StorageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageEvent.ProtoReflect.Descriptor instead.
func (*StorageEvent) Descriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{5}
}

func (x *StorageEvent) GetType() StorageEventType {
	if x != nil {
		return x.Type
	}
	return StorageEventType_STORAGE_EVENT_TYPE_INVALID
}

func (x *StorageEvent) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *StorageEvent) GetDepot() []byte {
	if x != nil {
		return x.Depot
	}
	return nil
}

func (x *StorageEvent) GetAuditor() []byte {
	if x != nil {
		return x.Auditor
	}
	return nil
}

// Request for a storage contract of the given tx hash.
type GetStorageContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tx hash of the storage contract to be requested.
	TxHash []byte `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" ssz-size:"32"`
}

func (x *GetStorageContractRequest) Reset() {
	*x = GetStorageContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_contract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStorageContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorageContractRequest) ProtoMessage() {}

func (x *GetStorageContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_contract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorageContractRequest.ProtoReflect.Descriptor instead.
func (*GetStorageContractRequest) Descriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{6}
}

func (x *GetStorageContractRequest) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

// Response for the GetStorageContractRequest.
type GetStorageContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract *StorageContract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *GetStorageContractResponse) Reset() {
	*x = GetStorageContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_contract_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStorageContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorageContractResponse) ProtoMessage() {}

func (x *GetStorageContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_contract_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorageContractResponse.ProtoReflect.Descriptor instead.
func (*GetStorageContractResponse) Descriptor() ([]byte, []int) {
	return file_consensus_contract_proto_rawDescGZIP(), []int{7}
}

func (x *GetStorageContractResponse) GetContract() *StorageContract {
	if x != nil {
		return x.Contract
	}
	return nil
}

var File_consensus_contract_proto protoreflect.FileDescriptor

var file_consensus_contract_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x1a, 0x11, 0x65, 0x78,
	0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8b, 0x04, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x78,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18,
	0x02, 0x33, 0x32, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x27, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0a,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a,
	0xb5, 0x18, 0x02, 0x34, 0x38, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x05,
	0x64, 0x65, 0x70, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18,
	0x02, 0x34, 0x38, 0x52, 0x05, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x12, 0x34, 0x0a, 0x12, 0x64, 0x65,
	0x70, 0x6f, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x10,
	0x64, 0x65, 0x70, 0x6f, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x08, 0x82, 0xb5, 0x18, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x1a, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0x82,
	0xb5, 0x18, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x62, 0x6f,
	0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x11, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x52,
	0x10, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x07, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x34, 0x38, 0x52, 0x07, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52,
	0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x08, 0x70, 0x6f,
	0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5,
	0x18, 0x02, 0x33, 0x32, 0x52, 0x07, 0x70, 0x6f, 0x72, 0x48, 0x61, 0x73, 0x68, 0x22, 0xb5, 0x01,
	0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x53, 0x6c, 0x6f, 0x74,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a,
	0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0b, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0c, 0x42, 0x09, 0x92, 0xb5, 0x18, 0x05, 0x36, 0x34, 0x2c, 0x39, 0x36, 0x52, 0x05,
	0x72, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x56, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32,
	0x52, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xdc, 0x01,
	0x0a, 0x03, 0x50, 0x6f, 0x52, 0x12, 0x1e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02,
	0x33, 0x32, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x25,
	0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x07, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x34, 0x38, 0x52, 0x07,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x06, 0x92, 0xb5, 0x18, 0x02, 0x33, 0x32,
	0x52, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x11,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x36, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x0c, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x06, 0x74, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x34, 0x38, 0x52, 0x05, 0x64, 0x65, 0x70, 0x6f,
	0x74, 0x12, 0x20, 0x0a, 0x07, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x34, 0x38, 0x52, 0x07, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x22, 0x3c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x22, 0x5b, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2a, 0x46,
	0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x13, 0x0a, 0x0f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x55, 0x44, 0x49, 0x54, 0x10, 0x03, 0x2a, 0x7d, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54,
	0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55,
	0x44, 0x49, 0x54, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x4f, 0x52,
	0x5f, 0x4e, 0x45, 0x57, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x04, 0x42, 0x99, 0x01, 0x0a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x42, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2d, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x3b, 0x70, 0x62, 0x63, 0xaa,
	0x02, 0x16, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0xca, 0x02, 0x16, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x6e, 0x5c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consensus_contract_proto_rawDescOnce sync.Once
	file_consensus_contract_proto_rawDescData = file_consensus_contract_proto_rawDesc
)

func file_consensus_contract_proto_rawDescGZIP() []byte {
	file_consensus_contract_proto_rawDescOnce.Do(func() {
		file_consensus_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_contract_proto_rawDescData)
	})
	return file_consensus_contract_proto_rawDescData
}

var file_consensus_contract_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_consensus_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_consensus_contract_proto_goTypes = []interface{}{
	(StorageStatus)(0),                 // 0: photon.consensus.StorageStatus
	(StorageEventType)(0),              // 1: photon.consensus.StorageEventType
	(*StorageContract)(nil),            // 2: photon.consensus.StorageContract
	(*Audit)(nil),                      // 3: photon.consensus.Audit
	(*Challenge)(nil),                  // 4: photon.consensus.Challenge
	(*PoR)(nil),                        // 5: photon.consensus.PoR
	(*EpochStorageEvent)(nil),          // 6: photon.consensus.EpochStorageEvent
	(*StorageEvent)(nil),               // 7: photon.consensus.StorageEvent
	(*GetStorageContractRequest)(nil),  // 8: photon.consensus.GetStorageContractRequest
	(*GetStorageContractResponse)(nil), // 9: photon.consensus.GetStorageContractResponse
}
var file_consensus_contract_proto_depIdxs = []int32{
	4, // 0: photon.consensus.PoR.challenges:type_name -> photon.consensus.Challenge
	7, // 1: photon.consensus.EpochStorageEvent.events:type_name -> photon.consensus.StorageEvent
	1, // 2: photon.consensus.StorageEvent.type:type_name -> photon.consensus.StorageEventType
	2, // 3: photon.consensus.GetStorageContractResponse.contract:type_name -> photon.consensus.StorageContract
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_consensus_contract_proto_init() }
func file_consensus_contract_proto_init() {
	if File_consensus_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consensus_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageContract); i {
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
		file_consensus_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audit); i {
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
		file_consensus_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Challenge); i {
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
		file_consensus_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoR); i {
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
		file_consensus_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochStorageEvent); i {
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
		file_consensus_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageEvent); i {
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
		file_consensus_contract_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStorageContractRequest); i {
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
		file_consensus_contract_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStorageContractResponse); i {
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
			RawDescriptor: file_consensus_contract_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consensus_contract_proto_goTypes,
		DependencyIndexes: file_consensus_contract_proto_depIdxs,
		EnumInfos:         file_consensus_contract_proto_enumTypes,
		MessageInfos:      file_consensus_contract_proto_msgTypes,
	}.Build()
	File_consensus_contract_proto = out.File
	file_consensus_contract_proto_rawDesc = nil
	file_consensus_contract_proto_goTypes = nil
	file_consensus_contract_proto_depIdxs = nil
}
