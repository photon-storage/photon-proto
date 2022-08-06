package pbc_test

import (
	"testing"

	pbc "github.com/photon-storage/photon-proto/consensus"
)

func TestCommitteeIndex_Casting(t *testing.T) {
	committeeIdx := pbc.CommitteeIndex(42)

	t.Run("floats", func(t *testing.T) {
		var x1 float32 = 42.2
		if pbc.CommitteeIndex(x1) != committeeIdx {
			t.Errorf("Unequal: %v = %v", pbc.CommitteeIndex(x1), committeeIdx)
		}

		var x2 float64 = 42.2
		if pbc.CommitteeIndex(x2) != committeeIdx {
			t.Errorf("Unequal: %v = %v", pbc.CommitteeIndex(x2), committeeIdx)
		}
	})

	t.Run("int", func(t *testing.T) {
		var x int = 42
		if pbc.CommitteeIndex(x) != committeeIdx {
			t.Errorf("Unequal: %v = %v", pbc.CommitteeIndex(x), committeeIdx)
		}
	})
}
