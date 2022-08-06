package pbc

import (
	"testing"
)

func TestUint64ToBytes_RoundTrip(t *testing.T) {
	for i := uint64(0); i < 10000; i++ {
		b := uint64ToBytesBigEndian(i)
		if got := bytesToUint64BigEndian(b); got != i {
			t.Error("Round trip did not match original value")
		}
	}
}
