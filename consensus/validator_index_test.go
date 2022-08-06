package pbc_test

import (
	"testing"
	"time"

	pbc "github.com/photon-storage/photon-proto/consensus"
)

func TestValidatorIndex_Casting(t *testing.T) {
	valIdx := pbc.ValidatorIndex(42)

	t.Run("time.Duration", func(t *testing.T) {
		if uint64(time.Duration(valIdx)) != uint64(valIdx) {
			t.Error("ValidatorIndex should produce the same result with time.Duration")
		}
	})

	t.Run("floats", func(t *testing.T) {
		var x1 float32 = 42.2
		if pbc.ValidatorIndex(x1) != valIdx {
			t.Errorf("Unequal: %v = %v", pbc.ValidatorIndex(x1), valIdx)
		}

		var x2 float64 = 42.2
		if pbc.ValidatorIndex(x2) != valIdx {
			t.Errorf("Unequal: %v = %v", pbc.ValidatorIndex(x2), valIdx)
		}
	})

	t.Run("int", func(t *testing.T) {
		var x int = 42
		if pbc.ValidatorIndex(x) != valIdx {
			t.Errorf("Unequal: %v = %v", pbc.ValidatorIndex(x), valIdx)
		}
	})
}
