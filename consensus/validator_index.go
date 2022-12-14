package pbc

import (
	"encoding/binary"
	"fmt"

	fssz "github.com/photon-storage/fastssz"
)

var _ fssz.HashRoot = (ValidatorIndex)(0)
var _ fssz.Marshaler = (*ValidatorIndex)(nil)
var _ fssz.Unmarshaler = (*ValidatorIndex)(nil)

// ValidatorIndex in eth2.
type ValidatorIndex uint64

// ValidatorIndexFromUniqueKey convert unique key bytes to validator index.
func ValidatorIndexFromUniqueKey(b []byte) ValidatorIndex {
	return ValidatorIndex(binary.BigEndian.Uint64(b))
}

// Div divides validator index by x.
func (v ValidatorIndex) Div(x uint64) ValidatorIndex {
	if x == 0 {
		panic("divbyzero")
	}
	return ValidatorIndex(uint64(v) / x)
}

// Add increases validator index by x.
func (v ValidatorIndex) Add(x uint64) ValidatorIndex {
	return ValidatorIndex(uint64(v) + x)
}

// Sub subtracts x from the validator index.
func (v ValidatorIndex) Sub(x uint64) ValidatorIndex {
	if uint64(v) < x {
		panic("underflow")
	}
	return ValidatorIndex(uint64(v) - x)
}

// Mod returns result of `validator index % x`.
func (v ValidatorIndex) Mod(x uint64) ValidatorIndex {
	return ValidatorIndex(uint64(v) % x)
}

// HashTreeRoot returns calculated hash root.
func (v ValidatorIndex) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(v)
}

// HashWithDefaultHasher hashes a HashRoot object with a Hasher from the default HasherPool.
func (v ValidatorIndex) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutUint64(uint64(v))
	return nil
}

// UnmarshalSSZ deserializes the provided bytes buffer into the validator index object.
func (v *ValidatorIndex) UnmarshalSSZ(buf []byte) error {
	if len(buf) != v.SizeSSZ() {
		return fmt.Errorf("expected buffer of length %d received %d", v.SizeSSZ(), len(buf))
	}
	*v = ValidatorIndex(fssz.UnmarshallUint64(buf))
	return nil
}

// MarshalSSZTo marshals validator index with the provided byte slice.
func (v *ValidatorIndex) MarshalSSZTo(dst []byte) ([]byte, error) {
	marshalled, err := v.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	return append(dst, marshalled...), nil
}

// MarshalSSZ marshals validator index into a serialized object.
func (v *ValidatorIndex) MarshalSSZ() ([]byte, error) {
	marshalled := fssz.MarshalUint64([]byte{}, uint64(*v))
	return marshalled, nil
}

// SizeSSZ returns the size of the serialized object.
func (v *ValidatorIndex) SizeSSZ() int {
	return 8
}

// UniqueKey convert the ValidatorIndex to the unique hash.
func (v ValidatorIndex) UniqueKey() [32]byte {
	return toBytes32(uint64ToBytesBigEndian(uint64(v)))
}
