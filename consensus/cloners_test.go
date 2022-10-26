package pbc

import (
	"fmt"
	"math/rand"
	"path/filepath"
	reflect "reflect"
	"runtime"
	"strings"
	"testing"

	"google.golang.org/protobuf/proto"
)

func notEmpty(t *testing.T, obj interface{}, msgs ...interface{}) {
	_, isProto := obj.(proto.Message)
	notEmptyImpl(t, obj, isProto, []string{} /*fields*/, 0 /*stackSize*/, msgs...)
}

// notEmpty checks all fields are not zero, including pointer field references
// to other structs. This method has the option to ignore fields without struct
// tags, which is helpful for checking protobuf messages that have internal
// fields.
func notEmptyImpl(
	t *testing.T,
	obj interface{},
	ignoreFieldsWithoutTags bool,
	fields []string,
	stackSize int,
	msgs ...interface{},
) {
	var v reflect.Value
	if vo, ok := obj.(reflect.Value); ok {
		v = reflect.Indirect(vo)
	} else {
		v = reflect.Indirect(reflect.ValueOf(obj))
	}

	if len(fields) == 0 {
		fields = []string{v.Type().Name()}
	}

	parseMsg := func(defaultMsg string, msgs ...interface{}) string {
		if len(msgs) >= 1 {
			msgFormat, ok := msgs[0].(string)
			if !ok {
				return defaultMsg
			}
			return fmt.Sprintf(msgFormat, msgs[1:]...)
		}
		return defaultMsg
	}
	fail := func(fields []string) {
		m := parseMsg("", msgs...)
		errMsg := fmt.Sprintf("empty/zero field: %s", strings.Join(fields, "."))
		if len(m) > 0 {
			m = strings.Join([]string{m, errMsg}, ": ")
		} else {
			m = errMsg
		}
		_, file, line, _ := runtime.Caller(4 + stackSize)
		t.Fatalf("%s:%d %s", filepath.Base(file), line, m)
	}

	if v.Kind() != reflect.Struct {
		if v.IsZero() {
			fail(fields)
		}
		return
	}

	for i := 0; i < v.NumField(); i++ {
		if ignoreFieldsWithoutTags && len(v.Type().Field(i).Tag) == 0 {
			continue
		}
		fields := append(fields, v.Type().Field(i).Name)

		switch k := v.Field(i).Kind(); k {
		case reflect.Ptr:
			notEmptyImpl(t, v.Field(i), ignoreFieldsWithoutTags, fields, stackSize+1, msgs...)
		case reflect.Slice:
			f := v.Field(i)
			if f.Len() == 0 {
				fail(fields)
			}
			for i := 0; i < f.Len(); i++ {
				notEmptyImpl(t, f.Index(i), ignoreFieldsWithoutTags, fields, stackSize+1, msgs...)
			}
		default:
			if v.Field(i).IsZero() {
				fail(fields)
			}
		}
	}
}

func TestCopyAttestation(t *testing.T) {
	att := &Attestation{
		AggregationBits: bytes(),
		Data: &AttestationData{
			Slot:            1,
			CommitteeIndex:  2,
			BeaconBlockRoot: bytes(),
			Source: &Checkpoint{
				Epoch: 1,
				Root:  bytes(),
			},
			Target: &Checkpoint{
				Epoch: 1,
				Root:  bytes(),
			},
		},
		Signature: bytes(),
	}

	got := CopyAttestation(att)
	if !reflect.DeepEqual(got, att) {
		t.Errorf("CopyAttestation() = %v, want %v", got, att)
	}
	notEmpty(t, got, "Copied attestation has empty fields")
}
func TestCopyAttestationData(t *testing.T) {
	att := &AttestationData{
		Slot:            1,
		CommitteeIndex:  2,
		BeaconBlockRoot: bytes(),
		Source: &Checkpoint{
			Epoch: 1,
			Root:  bytes(),
		},
		Target: &Checkpoint{
			Epoch: 1,
			Root:  bytes(),
		},
	}

	got := CopyAttestationData(att)
	if !reflect.DeepEqual(got, att) {
		t.Errorf("CopyAttestationData() = %v, want %v", got, att)
	}
	notEmpty(t, got, "Copied attestation data has empty fields")
}

func TestCopyCheckpoint(t *testing.T) {
	cp := &Checkpoint{
		Epoch: 1,
		Root:  bytes(),
	}
	got := CopyCheckpoint(cp)
	if !reflect.DeepEqual(got, cp) {
		t.Errorf("CopyCheckpoint() = %v, want %v", got, cp)
	}
	notEmpty(t, got, "Copied checkpoint has empty fields")
}

func TestCopySignedBlock(t *testing.T) {
	sbb := &SignedBlock{
		Block: &Block{
			Slot:          123455,
			ProposerIndex: 55433,
			ParentRoot:    bytes(),
			StateRoot:     bytes(),
			Body: &BlockBody{
				RandaoReveal:      bytes(),
				Graffiti:          bytes(),
				Txs:               genSignedTxs(5),
				ProposerSlashings: genProposerSlashings(t, 5),
				AttesterSlashings: genAttesterSlashings(t, 5),
				Attestations:      genAttestations(t, 10),
			},
		},
		Signature: bytes(),
	}

	got := CopySignedBlock(sbb)
	if !reflect.DeepEqual(got, sbb) {
		t.Errorf("CopySignedBlock() = %v, want %v", got, sbb)
	}
	notEmpty(t, sbb, "Copied signed block has empty fields")
}

func TestCopyBlock(t *testing.T) {
	b := &Block{
		Slot:          123455,
		ProposerIndex: 55433,
		ParentRoot:    bytes(),
		StateRoot:     bytes(),
		Body: &BlockBody{
			RandaoReveal:      bytes(),
			Graffiti:          bytes(),
			Txs:               genSignedTxs(5),
			ProposerSlashings: genProposerSlashings(t, 5),
			AttesterSlashings: genAttesterSlashings(t, 5),
			Attestations:      genAttestations(t, 10),
		},
	}

	got := CopyBlock(b)
	if !reflect.DeepEqual(got, b) {
		t.Errorf("CopyBlock() = %v, want %v", got, b)
	}
	notEmpty(t, b, "Copied block has empty fields")
}

func TestCopyBlockBody(t *testing.T) {
	bb := &BlockBody{
		RandaoReveal:      bytes(),
		Graffiti:          bytes(),
		Txs:               genSignedTxs(5),
		ProposerSlashings: genProposerSlashings(t, 5),
		AttesterSlashings: genAttesterSlashings(t, 5),
		Attestations:      genAttestations(t, 10),
	}

	got := CopyBlockBody(bb)
	if !reflect.DeepEqual(got, bb) {
		t.Errorf("CopyBlockBody() = %v, want %v", got, bb)
	}
	notEmpty(t, bb, "Copied block body has empty fields")
}

func TestCopyTxs(t *testing.T) {
	txs := genTxs(10)

	got := CopyTxs(txs)
	if !reflect.DeepEqual(got, txs) {
		t.Errorf("TestCopyTxs() = %v, want %v", got, txs)
	}
	notEmpty(t, got, "Copied transactions have empty fields")
}

func TestCopySignedTxs(t *testing.T) {
	txs := genSignedTxs(10)

	got := CopySignedTxs(txs)
	if !reflect.DeepEqual(got, txs) {
		t.Errorf("TestCopySignedTxs() = %v, want %v", got, txs)
	}
	notEmpty(t, got, "Copied signed transactions have empty fields")
}

func TestCopyTx(t *testing.T) {
	tx := genTx()
	got := CopyTx(tx)
	if !reflect.DeepEqual(got, tx) {
		t.Errorf("CopyTx() = %v, want %v", got, tx)
	}
	notEmpty(t, got, "Copied transaction has empty fields")
}

func TestCopyProposerSlashings(t *testing.T) {
	ps := genProposerSlashings(t, 10)

	got := CopyProposerSlashings(ps)
	if !reflect.DeepEqual(got, ps) {
		t.Errorf("CopyProposerSlashings() = %v, want %v", got, ps)
	}
	notEmpty(t, got, "Copied proposer slashings have empty fields")
}

func TestCopyProposerSlashing(t *testing.T) {
	ps := &ProposerSlashing{
		Header_1: &SignedBlockHeader{
			Header: &BlockHeader{
				Slot:          Slot(10),
				ProposerIndex: ValidatorIndex(15),
				ParentRoot:    bytes(),
				StateRoot:     bytes(),
				BodyRoot:      bytes(),
			},
			Signature: bytesN(96),
		},
		Header_2: &SignedBlockHeader{
			Header: &BlockHeader{
				Slot:          Slot(10),
				ProposerIndex: ValidatorIndex(15),
				ParentRoot:    bytes(),
				StateRoot:     bytes(),
				BodyRoot:      bytes(),
			},
			Signature: bytesN(96),
		},
	}
	got := CopyProposerSlashing(ps)
	if !reflect.DeepEqual(got, ps) {
		t.Errorf("CopyProposerSlashing() = %v, want %v", got, ps)
	}
	notEmpty(t, got, "Copied proposer slashing has empty fields")
}

func TestCopySignedBlockHeader(t *testing.T) {
	sbh := &SignedBlockHeader{
		Header: &BlockHeader{
			Slot:          Slot(10),
			ProposerIndex: ValidatorIndex(15),
			ParentRoot:    bytes(),
			StateRoot:     bytes(),
			BodyRoot:      bytes(),
		},
		Signature: bytesN(96),
	}
	got := CopySignedBlockHeader(sbh)
	if !reflect.DeepEqual(got, sbh) {
		t.Errorf("CopySignedBlockHeader() = %v, want %v", got, sbh)
	}
	notEmpty(t, got, "Copied signed block header has empty fields")
}

func TestCopyBlockHeader(t *testing.T) {
	bh := &BlockHeader{
		Slot:          Slot(10),
		ProposerIndex: ValidatorIndex(15),
		ParentRoot:    bytes(),
		StateRoot:     bytes(),
		BodyRoot:      bytes(),
	}
	got := CopyBlockHeader(bh)
	if !reflect.DeepEqual(got, bh) {
		t.Errorf("CopyBlockHeader() = %v, want %v", got, bh)
	}
	notEmpty(t, got, "Copied block header has empty fields")
}

func TestCopyAttesterSlashings(t *testing.T) {
	as := genAttesterSlashings(t, 10)

	got := CopyAttesterSlashings(as)
	if !reflect.DeepEqual(got, as) {
		t.Errorf("CopyAttesterSlashings() = %v, want %v", got, as)
	}
	notEmpty(t, got, "Copied attester slashings have empty fields")
}

func TestCopyIndexedAttestation(t *testing.T) {
	ia := &IndexedAttestation{
		AttestingIndices: []uint64{1, 2, 3},
		Data: &AttestationData{
			Slot:            1,
			CommitteeIndex:  2,
			BeaconBlockRoot: bytes(),
			Source: &Checkpoint{
				Epoch: 1,
				Root:  bytes(),
			},
			Target: &Checkpoint{
				Epoch: 1,
				Root:  bytes(),
			},
		},
		Signature: bytes(),
	}

	got := CopyIndexedAttestation(ia)
	if !reflect.DeepEqual(got, ia) {
		t.Errorf("CopyIndexedAttestation() = %v, want %v", got, ia)
	}
	notEmpty(t, got, "Copied indexed attestation has empty fields")
}

func TestCopyAttestations(t *testing.T) {
	atts := genAttestations(t, 10)

	got := CopyAttestations(atts)
	if !reflect.DeepEqual(got, atts) {
		t.Errorf("CopyAttestations() = %v, want %v", got, atts)
	}
	notEmpty(t, got, "Copied attestations have empty fields")
}

func TestCopyValidator(t *testing.T) {
	v := &Validator{
		PublicKey:                  bytes(),
		Slashed:                    true,
		ActivationEligibilityEpoch: 14322,
		ActivationEpoch:            14325,
		ExitEpoch:                  23425,
		WithdrawableEpoch:          30000,
	}

	got := CopyValidator(v)
	if !reflect.DeepEqual(got, v) {
		t.Errorf("CopyValidator() = %v, want %v", got, v)
	}
	notEmpty(t, got, "Copied validator has empty fields")
}

func bytes() []byte {
	return bytesN(32)
}

func bytesN(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	for i := 0; i < n; i++ {
		if b[i] == 0x00 {
			b[i]++
		}
	}
	return b
}

func genAttestations(t *testing.T, num int) []*Attestation {
	atts := make([]*Attestation, num)
	for i := 0; i < num; i++ {
		atts[i] = &Attestation{
			AggregationBits: bytes(),
			Data: &AttestationData{
				Slot:            1,
				CommitteeIndex:  2,
				BeaconBlockRoot: bytes(),
				Source: &Checkpoint{
					Epoch: 1,
					Root:  bytes(),
				},
				Target: &Checkpoint{
					Epoch: 1,
					Root:  bytes(),
				},
			},
			Signature: bytes(),
		}
	}
	return atts
}

func genTx() *Transaction {
	return &Transaction{
		Type:     1,
		From:     bytesN(48),
		ChainId:  2,
		Nonce:    3,
		GasPrice: 100,
		GasLimit: 100000,
		TxDataBalanceTransfer: &TxDataBalanceTransfer{
			To:     bytesN(48),
			Amount: 100,
		},
		TxDataValidatorDeposit: &TxDataValidatorDeposit{
			Amount: 200,
		},
		TxDataAuditorDeposit: &TxDataAuditorDeposit{
			Amount:  300,
			Decoder: bytesN(512),
		},
		TxDataObjectCommit: &TxDataObjectCommit{
			Owner:            bytesN(48),
			Depot:            bytesN(48),
			DepotDiscoveryId: bytesN(32),
			Hash:             bytesN(32),
			Size:             4096,
			EncodedHash:      bytesN(32),
			EncodedSize:      4096,
			NumBlocks:        4,
			Duration:         10000,
			Fee:              5000000000,
			Pledge:           1000000000,
			Deadline:         2000000000,
		},
		TxDataObjectAudit: &TxDataObjectAudit{
			CommitTxHash: bytesN(32),
			Auditor:      bytesN(48),
			Depot:        bytesN(48),
			Hash:         bytesN(32),
			Size:         4096,
			EncodedHash:  bytesN(32),
			EncodedSize:  4096,
			NumBlocks:    4,
			Rands: [][]byte{
				bytesN(96),
				bytesN(96),
				bytesN(96),
				bytesN(96),
				bytesN(96),
				bytesN(96),
				bytesN(96),
			},
		},
		TxDataObjectChallenge: &TxDataObjectChallenge{
			Depot:        bytesN(48),
			CommitTxHash: bytesN(32),
		},
		TxDataObjectPoR: &TxDataObjectPoR{
			CommitTxHash: bytesN(32),
			BlockAggs: [][]byte{
				bytesN(32),
				bytesN(32),
				bytesN(32),
				bytesN(32),
				bytesN(32),
				bytesN(32),
				bytesN(32),
				bytesN(32),
			},
			Sigma: bytesN(96),
		},
	}
}
func genTxs(num int) []*Transaction {
	txs := make([]*Transaction, num)
	for i := 0; i < num; i++ {
		txs[i] = genTx()
	}
	return txs
}

func genSignedTxs(num int) []*SignedTransaction {
	txs := make([]*SignedTransaction, num)
	for i := 0; i < num; i++ {
		txs[i] = &SignedTransaction{
			Tx:        genTx(),
			Signature: bytesN(96),
		}
	}
	return txs
}

func genProposerSlashings(t *testing.T, num int) []*ProposerSlashing {
	ps := make([]*ProposerSlashing, num)
	for i := 0; i < num; i++ {
		ps[i] = &ProposerSlashing{
			Header_1: &SignedBlockHeader{
				Header: &BlockHeader{
					Slot:          Slot(10),
					ProposerIndex: ValidatorIndex(15),
					ParentRoot:    bytes(),
					StateRoot:     bytes(),
					BodyRoot:      bytes(),
				},
				Signature: bytesN(96),
			},
			Header_2: &SignedBlockHeader{
				Header: &BlockHeader{
					Slot:          Slot(10),
					ProposerIndex: ValidatorIndex(15),
					ParentRoot:    bytes(),
					StateRoot:     bytes(),
					BodyRoot:      bytes(),
				},
				Signature: bytesN(96),
			},
		}
	}
	return ps
}

func genAttesterSlashings(t *testing.T, num int) []*AttesterSlashing {
	as := make([]*AttesterSlashing, num)
	for i := 0; i < num; i++ {
		as[i] = &AttesterSlashing{
			Attestation_1: &IndexedAttestation{
				AttestingIndices: []uint64{1, 2, 3},
				Data: &AttestationData{
					Slot:            1,
					CommitteeIndex:  2,
					BeaconBlockRoot: bytes(),
					Source: &Checkpoint{
						Epoch: 1,
						Root:  bytes(),
					},
					Target: &Checkpoint{
						Epoch: 1,
						Root:  bytes(),
					},
				},
				Signature: bytes(),
			},
			Attestation_2: &IndexedAttestation{
				AttestingIndices: []uint64{1, 2, 3},
				Data: &AttestationData{
					Slot:            1,
					CommitteeIndex:  2,
					BeaconBlockRoot: bytes(),
					Source: &Checkpoint{
						Epoch: 1,
						Root:  bytes(),
					},
					Target: &Checkpoint{
						Epoch: 1,
						Root:  bytes(),
					},
				},
				Signature: bytes(),
			},
		}
	}
	return as
}
