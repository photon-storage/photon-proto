package pbc

import (
	"encoding/binary"
)

// toBytes32 is a convenience method for converting a byte slice to a fix
// sized 32 byte array. This method will truncate the input if it is larger
// than 32 bytes.
func toBytes32(x []byte) [32]byte {
	var y [32]byte
	copy(y[:], x)
	return y
}

// safeCopyBytes will copy and return a non-nil byte array, otherwise it
// returns nil.
func safeCopyBytes(cp []byte) []byte {
	if cp != nil {
		copied := make([]byte, len(cp))
		copy(copied, cp)
		return copied
	}
	return nil
}

// uint64ToBytesLittleEndian conversion.
func uint64ToBytesLittleEndian(i uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, i)
	return buf
}

// uint64ToBytesBigEndian conversion.
func uint64ToBytesBigEndian(i uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

// bytesToUint64BigEndian conversion. Returns 0 if empty bytes or byte slice
// with length less than 8.
func bytesToUint64BigEndian(b []byte) uint64 {
	if len(b) < 8 { // This will panic otherwise.
		return 0
	}
	return binary.BigEndian.Uint64(b)
}
