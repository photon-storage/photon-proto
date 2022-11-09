// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: consensus/state.proto

package pbc

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/photon-storage/photon-proto/ext"
	github_com_prysmaticlabs_go_bitfield "github.com/prysmaticlabs/go-bitfield"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Versioning [1-100]
	GenesisTime       uint64 `protobuf:"varint,1,opt,name=genesis_time,json=genesisTime,proto3" json:"genesis_time,omitempty"`
	GenesisIdentifier []byte `protobuf:"bytes,2,opt,name=genesis_identifier,json=genesisIdentifier,proto3" json:"genesis_identifier,omitempty" ssz-size:"32"`
	Slot              Slot   `protobuf:"varint,3,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"Slot"`
	// Block [101-200]
	LatestBlockHeader *BlockHeader `protobuf:"bytes,101,opt,name=latest_block_header,json=latestBlockHeader,proto3" json:"latest_block_header,omitempty"`
	// Finality [201-300]
	// Spec type [4]Bitvector which means this would be a fixed size of 4 bits.
	JustificationBits           github_com_prysmaticlabs_go_bitfield.Bitvector4 `protobuf:"bytes,201,opt,name=justification_bits,json=justificationBits,proto3" json:"justification_bits,omitempty" cast-type:"github.com/prysmaticlabs/go-bitfield.Bitvector4" ssz-size:"1"`
	PreviousJustifiedCheckpoint *Checkpoint                                     `protobuf:"bytes,202,opt,name=previous_justified_checkpoint,json=previousJustifiedCheckpoint,proto3" json:"previous_justified_checkpoint,omitempty"`
	CurrentJustifiedCheckpoint  *Checkpoint                                     `protobuf:"bytes,203,opt,name=current_justified_checkpoint,json=currentJustifiedCheckpoint,proto3" json:"current_justified_checkpoint,omitempty"`
	FinalizedCheckpoint         *Checkpoint                                     `protobuf:"bytes,204,opt,name=finalized_checkpoint,json=finalizedCheckpoint,proto3" json:"finalized_checkpoint,omitempty"`
	// Validator participation & slashings [301-400]
	Balances                   []uint64 `protobuf:"varint,301,rep,packed,name=balances,proto3" json:"balances,omitempty" ssz-max:"1099511627776"`
	InactivityScores           []uint64 `protobuf:"varint,302,rep,packed,name=inactivity_scores,json=inactivityScores,proto3" json:"inactivity_scores,omitempty" ssz-max:"1099511627776"`
	PreviousEpochParticipation []byte   `protobuf:"bytes,303,opt,name=previous_epoch_participation,json=previousEpochParticipation,proto3" json:"previous_epoch_participation,omitempty" ssz-max:"1099511627776"`
	CurrentEpochParticipation  []byte   `protobuf:"bytes,304,opt,name=current_epoch_participation,json=currentEpochParticipation,proto3" json:"current_epoch_participation,omitempty" ssz-max:"1099511627776"`
	Slashings                  []uint64 `protobuf:"varint,305,rep,packed,name=slashings,proto3" json:"slashings,omitempty" ssz-max:"8192"`
	// Trie roots [401-500]
	// Account trie root hash.
	AccountTrieRoot []byte `protobuf:"bytes,401,opt,name=account_trie_root,json=accountTrieRoot,proto3" json:"account_trie_root,omitempty" ssz-size:"32"`
	// Validator trie root hash.
	ValidatorTrieRoot []byte `protobuf:"bytes,402,opt,name=validator_trie_root,json=validatorTrieRoot,proto3" json:"validator_trie_root,omitempty" ssz-size:"32"`
	// Auditor trie root hash.
	AuditorTrieRoot []byte `protobuf:"bytes,403,opt,name=auditor_trie_root,json=auditorTrieRoot,proto3" json:"auditor_trie_root,omitempty" ssz-size:"32"`
	// Storage contract trie root hash.
	StorageTrieRoot []byte `protobuf:"bytes,404,opt,name=storage_trie_root,json=storageTrieRoot,proto3" json:"storage_trie_root,omitempty" ssz-size:"32"`
	// Storage contract audit queue's root hash.
	AuditQueueRoot []byte `protobuf:"bytes,405,opt,name=audit_queue_root,json=auditQueueRoot,proto3" json:"audit_queue_root,omitempty" ssz-size:"32"`
	// Storage contract event queue's root hash.
	EventQueueRoot []byte `protobuf:"bytes,406,opt,name=event_queue_root,json=eventQueueRoot,proto3" json:"event_queue_root,omitempty" ssz-size:"32"`
	// Aggregated state [501-600]
	ReserviorBalance uint64   `protobuf:"varint,501,opt,name=reservior_balance,json=reserviorBalance,proto3" json:"reservior_balance,omitempty"`
	Storage          *Storage `protobuf:"bytes,502,opt,name=storage,proto3" json:"storage,omitempty"`
	// Randomness [601-700]
	RandaoMixes [][]byte `protobuf:"bytes,601,rep,name=randao_mixes,json=randaoMixes,proto3" json:"randao_mixes,omitempty" ssz-max:"65536,32"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_consensus_state_proto_rawDescGZIP(), []int{0}
}

func (x *State) GetGenesisTime() uint64 {
	if x != nil {
		return x.GenesisTime
	}
	return 0
}

func (x *State) GetGenesisIdentifier() []byte {
	if x != nil {
		return x.GenesisIdentifier
	}
	return nil
}

func (x *State) GetSlot() Slot {
	if x != nil {
		return x.Slot
	}
	return Slot(0)
}

func (x *State) GetLatestBlockHeader() *BlockHeader {
	if x != nil {
		return x.LatestBlockHeader
	}
	return nil
}

func (x *State) GetJustificationBits() github_com_prysmaticlabs_go_bitfield.Bitvector4 {
	if x != nil {
		return x.JustificationBits
	}
	return github_com_prysmaticlabs_go_bitfield.Bitvector4(nil)
}

func (x *State) GetPreviousJustifiedCheckpoint() *Checkpoint {
	if x != nil {
		return x.PreviousJustifiedCheckpoint
	}
	return nil
}

func (x *State) GetCurrentJustifiedCheckpoint() *Checkpoint {
	if x != nil {
		return x.CurrentJustifiedCheckpoint
	}
	return nil
}

func (x *State) GetFinalizedCheckpoint() *Checkpoint {
	if x != nil {
		return x.FinalizedCheckpoint
	}
	return nil
}

func (x *State) GetBalances() []uint64 {
	if x != nil {
		return x.Balances
	}
	return nil
}

func (x *State) GetInactivityScores() []uint64 {
	if x != nil {
		return x.InactivityScores
	}
	return nil
}

func (x *State) GetPreviousEpochParticipation() []byte {
	if x != nil {
		return x.PreviousEpochParticipation
	}
	return nil
}

func (x *State) GetCurrentEpochParticipation() []byte {
	if x != nil {
		return x.CurrentEpochParticipation
	}
	return nil
}

func (x *State) GetSlashings() []uint64 {
	if x != nil {
		return x.Slashings
	}
	return nil
}

func (x *State) GetAccountTrieRoot() []byte {
	if x != nil {
		return x.AccountTrieRoot
	}
	return nil
}

func (x *State) GetValidatorTrieRoot() []byte {
	if x != nil {
		return x.ValidatorTrieRoot
	}
	return nil
}

func (x *State) GetAuditorTrieRoot() []byte {
	if x != nil {
		return x.AuditorTrieRoot
	}
	return nil
}

func (x *State) GetStorageTrieRoot() []byte {
	if x != nil {
		return x.StorageTrieRoot
	}
	return nil
}

func (x *State) GetAuditQueueRoot() []byte {
	if x != nil {
		return x.AuditQueueRoot
	}
	return nil
}

func (x *State) GetEventQueueRoot() []byte {
	if x != nil {
		return x.EventQueueRoot
	}
	return nil
}

func (x *State) GetReserviorBalance() uint64 {
	if x != nil {
		return x.ReserviorBalance
	}
	return 0
}

func (x *State) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *State) GetRandaoMixes() [][]byte {
	if x != nil {
		return x.RandaoMixes
	}
	return nil
}

// Metadata for all active storage contracts.
type Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of active storage contracts.
	NumContracts uint64 `protobuf:"varint,1,opt,name=num_contracts,json=numContracts,proto3" json:"num_contracts,omitempty"`
	// Total storage capacity size registered by storage contracts in bytes.
	// The integer value is serialized in big endian format.
	// This tracks the object's original size instead of the encoded size.
	TotalSize []byte `protobuf:"bytes,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty" ssz-max:"32"`
	// Accumulated reward rate at current slot used for tracking storage
	// contract's reward.
	RewardRate uint64 `protobuf:"varint,3,opt,name=reward_rate,json=rewardRate,proto3" json:"reward_rate,omitempty"`
	// Aggregated fee paid out to depot per slot.
	FeeRate uint64 `protobuf:"varint,4,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
	// The latest fee rate effective slot.
	FeeRateSince Slot `protobuf:"varint,5,opt,name=fee_rate_since,json=feeRateSince,proto3" json:"fee_rate_since,omitempty" cast-type:"Slot"`
}

func (x *Storage) Reset() {
	*x = Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_consensus_state_proto_rawDescGZIP(), []int{1}
}

func (x *Storage) GetNumContracts() uint64 {
	if x != nil {
		return x.NumContracts
	}
	return 0
}

func (x *Storage) GetTotalSize() []byte {
	if x != nil {
		return x.TotalSize
	}
	return nil
}

func (x *Storage) GetRewardRate() uint64 {
	if x != nil {
		return x.RewardRate
	}
	return 0
}

func (x *Storage) GetFeeRate() uint64 {
	if x != nil {
		return x.FeeRate
	}
	return 0
}

func (x *Storage) GetFeeRateSince() Slot {
	if x != nil {
		return x.FeeRateSince
	}
	return Slot(0)
}

var File_consensus_state_proto protoreflect.FileDescriptor

var file_consensus_state_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x1a, 0x11, 0x65, 0x78, 0x74, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe6, 0x0a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x65,
	0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x12, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33,
	0x32, 0x52, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c,
	0x6f, 0x74, 0x12, 0x4d, 0x0a, 0x13, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x11,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x68, 0x0a, 0x12, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x38,
	0x8a, 0xb5, 0x18, 0x01, 0x31, 0x82, 0xb5, 0x18, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x69, 0x74, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x42, 0x69,
	0x74, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x34, 0x52, 0x11, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x74, 0x73, 0x12, 0x61, 0x0a, 0x1d, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0xca, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x1b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4a, 0x75, 0x73, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x5f,
	0x0a, 0x1c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0xcb,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x1a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4a, 0x75, 0x73, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x50, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0xcc, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x13, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0xad, 0x02,
	0x20, 0x03, 0x28, 0x04, 0x42, 0x11, 0x92, 0xb5, 0x18, 0x0d, 0x31, 0x30, 0x39, 0x39, 0x35, 0x31,
	0x31, 0x36, 0x32, 0x37, 0x37, 0x37, 0x36, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x3f, 0x0a, 0x11, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0xae, 0x02, 0x20, 0x03, 0x28, 0x04, 0x42, 0x11, 0x92,
	0xb5, 0x18, 0x0d, 0x31, 0x30, 0x39, 0x39, 0x35, 0x31, 0x31, 0x36, 0x32, 0x37, 0x37, 0x37, 0x36,
	0x52, 0x10, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x12, 0x54, 0x0a, 0x1c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x65,
	0x70, 0x6f, 0x63, 0x68, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0xaf, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0x92, 0xb5, 0x18, 0x0d, 0x31,
	0x30, 0x39, 0x39, 0x35, 0x31, 0x31, 0x36, 0x32, 0x37, 0x37, 0x37, 0x36, 0x52, 0x1a, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x1b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xb0, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11,
	0x92, 0xb5, 0x18, 0x0d, 0x31, 0x30, 0x39, 0x39, 0x35, 0x31, 0x31, 0x36, 0x32, 0x37, 0x37, 0x37,
	0x36, 0x52, 0x19, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x09,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x18, 0xb1, 0x02, 0x20, 0x03, 0x28, 0x04,
	0x42, 0x08, 0x92, 0xb5, 0x18, 0x04, 0x38, 0x31, 0x39, 0x32, 0x52, 0x09, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x33, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x74, 0x72, 0x69, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x91, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x72, 0x69, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x37, 0x0a, 0x13, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x72, 0x69, 0x65, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x92, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32,
	0x52, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x72, 0x69, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x33, 0x0a, 0x11, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x74,
	0x72, 0x69, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x93, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x54, 0x72, 0x69, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x33, 0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x94, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x72, 0x69, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x31, 0x0a,
	0x10, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x95, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32,
	0x52, 0x0e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x6f, 0x6f, 0x74,
	0x12, 0x31, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x96, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18,
	0x02, 0x33, 0x32, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x72,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0xf5, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x10, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0xf6, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x72, 0x61, 0x6e, 0x64, 0x61,
	0x6f, 0x5f, 0x6d, 0x69, 0x78, 0x65, 0x73, 0x18, 0xd9, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x0c,
	0x92, 0xb5, 0x18, 0x08, 0x36, 0x35, 0x35, 0x33, 0x36, 0x2c, 0x33, 0x32, 0x52, 0x0b, 0x72, 0x61,
	0x6e, 0x64, 0x61, 0x6f, 0x4d, 0x69, 0x78, 0x65, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x07, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6e, 0x75,
	0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06,
	0x92, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x66, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a,
	0x0e, 0x66, 0x65, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x52,
	0x0c, 0x66, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x42, 0x96, 0x01,
	0x0a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x6e, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x3b, 0x70, 0x62, 0x63, 0xaa, 0x02, 0x16, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0xca, 0x02, 0x16,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x6e, 0x5c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consensus_state_proto_rawDescOnce sync.Once
	file_consensus_state_proto_rawDescData = file_consensus_state_proto_rawDesc
)

func file_consensus_state_proto_rawDescGZIP() []byte {
	file_consensus_state_proto_rawDescOnce.Do(func() {
		file_consensus_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_state_proto_rawDescData)
	})
	return file_consensus_state_proto_rawDescData
}

var file_consensus_state_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_consensus_state_proto_goTypes = []interface{}{
	(*State)(nil),       // 0: photon.consensus.State
	(*Storage)(nil),     // 1: photon.consensus.Storage
	(*BlockHeader)(nil), // 2: photon.consensus.BlockHeader
	(*Checkpoint)(nil),  // 3: photon.consensus.Checkpoint
}
var file_consensus_state_proto_depIdxs = []int32{
	2, // 0: photon.consensus.State.latest_block_header:type_name -> photon.consensus.BlockHeader
	3, // 1: photon.consensus.State.previous_justified_checkpoint:type_name -> photon.consensus.Checkpoint
	3, // 2: photon.consensus.State.current_justified_checkpoint:type_name -> photon.consensus.Checkpoint
	3, // 3: photon.consensus.State.finalized_checkpoint:type_name -> photon.consensus.Checkpoint
	1, // 4: photon.consensus.State.storage:type_name -> photon.consensus.Storage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_consensus_state_proto_init() }
func file_consensus_state_proto_init() {
	if File_consensus_state_proto != nil {
		return
	}
	file_consensus_attestation_proto_init()
	file_consensus_block_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_consensus_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_consensus_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storage); i {
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
			RawDescriptor: file_consensus_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consensus_state_proto_goTypes,
		DependencyIndexes: file_consensus_state_proto_depIdxs,
		MessageInfos:      file_consensus_state_proto_msgTypes,
	}.Build()
	File_consensus_state_proto = out.File
	file_consensus_state_proto_rawDesc = nil
	file_consensus_state_proto_goTypes = nil
	file_consensus_state_proto_depIdxs = nil
}
