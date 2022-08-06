package pbt

import (
	"errors"
	reflect "reflect"
	"strconv"
	"strings"
	"testing"

	pbc "github.com/photon-storage/photon-proto/consensus"
)

func TestSSZTagSize(t *testing.T) {
	sigSize := 96
	pubKeySize := 48
	rootSize := 32

	sizes, err := sszTagSizes(pbc.Attestation{}, "Signature")
	if err != nil {
		t.Error("unexpected error")
	}
	if sigSize != sizes[0] {
		t.Error("Unexpected signature size")
	}

	sizes, err = sszTagSizes(pbc.SignedBlock{}, "Signature")
	if err != nil {
		t.Error("unexpected error")
	}
	if sigSize != sizes[0] {
		t.Error("Unexpected signature size")
	}

	sizes, err = sszTagSizes(pbc.Checkpoint{}, "Root")
	if err != nil {
		t.Error("unexpected error")
	}
	if rootSize != sizes[0] {
		t.Error("Unexpected signature size")
	}

	sizes, err = sszTagSizes(pbc.Validator{}, "PublicKey")
	if err != nil {
		t.Error("unexpected error")
	}
	if pubKeySize != sizes[0] {
		t.Error("Unexpected signature size")
	}
}

func sszTagSizes(i interface{}, fName string) ([]int, error) {
	v := reflect.ValueOf(i)
	field, exists := v.Type().FieldByName(fName)
	if !exists {
		return nil, errors.New("wanted field does not exist")
	}
	tag, exists := field.Tag.Lookup("ssz-size")
	if !exists {
		return nil, errors.New("wanted field does not contain ssz-size tag")
	}
	start := strings.IndexRune(tag, '=')
	items := strings.Split(tag[start+1:], ",")
	sizes := make([]int, len(items))
	var err error
	for i := 0; i < len(items); i++ {
		if items[i] == "?" {
			sizes[i] = 0
			continue
		}
		sizes[i], err = strconv.Atoi(items[i])
		if err != nil {
			return nil, err
		}
	}
	return sizes, nil
}
