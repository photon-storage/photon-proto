package pbc

import (
	"google.golang.org/protobuf/proto"
)

// CopyAttestation copies the provided attestation object.
func CopyAttestation(att *Attestation) *Attestation {
	if att == nil {
		return nil
	}
	return &Attestation{
		AggregationBits: safeCopyBytes(att.AggregationBits),
		Data:            CopyAttestationData(att.Data),
		Signature:       safeCopyBytes(att.Signature),
	}
}

// CopyAttestationData copies the provided AttestationData object.
func CopyAttestationData(attData *AttestationData) *AttestationData {
	if attData == nil {
		return nil
	}
	return &AttestationData{
		Slot:            attData.Slot,
		CommitteeIndex:  attData.CommitteeIndex,
		BeaconBlockRoot: safeCopyBytes(attData.BeaconBlockRoot),
		Source:          CopyCheckpoint(attData.Source),
		Target:          CopyCheckpoint(attData.Target),
	}
}

// CopyCheckpoint copies the provided checkpoint.
func CopyCheckpoint(cp *Checkpoint) *Checkpoint {
	if cp == nil {
		return nil
	}
	return &Checkpoint{
		Epoch: cp.Epoch,
		Root:  safeCopyBytes(cp.Root),
	}
}

// CopySignedBlock copies the provided SignedBlock.
func CopySignedBlock(sigBlock *SignedBlock) *SignedBlock {
	if sigBlock == nil {
		return nil
	}
	return &SignedBlock{
		Block:     CopyBlock(sigBlock.Block),
		Signature: safeCopyBytes(sigBlock.Signature),
	}
}

// CopyBlock copies the provided Block.
func CopyBlock(block *Block) *Block {
	if block == nil {
		return nil
	}
	return &Block{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    safeCopyBytes(block.ParentRoot),
		StateRoot:     safeCopyBytes(block.StateRoot),
		Body:          CopyBlockBody(block.Body),
	}
}

// CopyBlockBody copies the provided BlockBody.
func CopyBlockBody(body *BlockBody) *BlockBody {
	if body == nil {
		return nil
	}
	return &BlockBody{
		RandaoReveal:      safeCopyBytes(body.RandaoReveal),
		Graffiti:          safeCopyBytes(body.Graffiti),
		Txs:               CopySignedTxs(body.Txs),
		ProposerSlashings: CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings: CopyAttesterSlashings(body.AttesterSlashings),
		Attestations:      CopyAttestations(body.Attestations),
	}
}

// CopySignedTxs copies the provided SignedTransaction array.
func CopySignedTxs(txs []*SignedTransaction) []*SignedTransaction {
	if txs == nil {
		return nil
	}
	newTxs := make([]*SignedTransaction, len(txs))
	for i, tx := range txs {
		newTxs[i] = CopySignedTx(tx)
	}
	return newTxs
}

// CopySignedTx copies the provided SignedTransaction.
func CopySignedTx(tx *SignedTransaction) *SignedTransaction {
	if tx == nil {
		return nil
	}
	return &SignedTransaction{
		Tx:        CopyTx(tx.Tx),
		Signature: safeCopyBytes(tx.Signature),
	}
}

// CopyTxs copies the provided Transaction array.
func CopyTxs(txs []*Transaction) []*Transaction {
	if txs == nil {
		return nil
	}
	newTxs := make([]*Transaction, len(txs))
	for i, tx := range txs {
		newTxs[i] = CopyTx(tx)
	}
	return newTxs
}

// CopyTx copies the provided Transaction.
func CopyTx(tx *Transaction) *Transaction {
	if tx == nil {
		return nil
	}

	var bf *TxDataBalanceTransfer
	if tx.TxDataBalanceTransfer != nil {
		bf = &TxDataBalanceTransfer{
			To:     safeCopyBytes(tx.TxDataBalanceTransfer.To),
			Amount: tx.TxDataBalanceTransfer.Amount,
		}
	}

	var vd *TxDataValidatorDeposit
	if tx.TxDataValidatorDeposit != nil {
		vd = &TxDataValidatorDeposit{
			Amount: tx.TxDataValidatorDeposit.Amount,
		}
	}

	var ad *TxDataAuditorDeposit
	if tx.TxDataAuditorDeposit != nil {
		ad = &TxDataAuditorDeposit{
			Amount:  tx.TxDataAuditorDeposit.Amount,
			Decoder: tx.TxDataAuditorDeposit.Decoder,
		}
	}

	var oc *TxDataObjectCommit
	if tx.TxDataObjectCommit != nil {
		oc = &TxDataObjectCommit{
			Owner:            safeCopyBytes(tx.TxDataObjectCommit.Owner),
			Depot:            safeCopyBytes(tx.TxDataObjectCommit.Depot),
			DepotDiscoveryId: safeCopyBytes(tx.TxDataObjectCommit.DepotDiscoveryId),
			Hash:             safeCopyBytes(tx.TxDataObjectCommit.Hash),
			Size:             tx.TxDataObjectCommit.Size,
			EncodedHash:      safeCopyBytes(tx.TxDataObjectCommit.EncodedHash),
			EncodedSize:      tx.TxDataObjectCommit.EncodedSize,
			NumBlocks:        tx.TxDataObjectCommit.NumBlocks,
			Duration:         tx.TxDataObjectCommit.Duration,
			Fee:              tx.TxDataObjectCommit.Fee,
			Bond:             tx.TxDataObjectCommit.Bond,
			Deadline:         tx.TxDataObjectCommit.Deadline,
		}
	}

	var oa *TxDataObjectAudit
	if tx.TxDataObjectAudit != nil {
		var rands [][]byte
		for _, v := range tx.TxDataObjectAudit.Rands {
			rands = append(rands, safeCopyBytes(v))
		}
		oa = &TxDataObjectAudit{
			CommitTxHash: safeCopyBytes(tx.TxDataObjectAudit.CommitTxHash),
			Auditor:      safeCopyBytes(tx.TxDataObjectAudit.Auditor),
			Depot:        safeCopyBytes(tx.TxDataObjectAudit.Depot),
			Hash:         safeCopyBytes(tx.TxDataObjectAudit.Hash),
			Size:         tx.TxDataObjectAudit.Size,
			EncodedHash:  safeCopyBytes(tx.TxDataObjectAudit.EncodedHash),
			EncodedSize:  tx.TxDataObjectAudit.EncodedSize,
			NumBlocks:    tx.TxDataObjectAudit.NumBlocks,
			Rands:        rands,
		}
	}

	var op *TxDataObjectPoR
	if tx.TxDataObjectPoR != nil {
		var aggs [][]byte
		for _, v := range tx.TxDataObjectPoR.BlockAggs {
			aggs = append(aggs, safeCopyBytes(v))
		}

		op = &TxDataObjectPoR{
			CommitTxHash: safeCopyBytes(tx.TxDataObjectPoR.CommitTxHash),
			BlockAggs:    aggs,
			Sigma:        safeCopyBytes(tx.TxDataObjectPoR.Sigma),
		}
	}

	return &Transaction{
		Type:                   tx.Type,
		From:                   safeCopyBytes(tx.From),
		ChainId:                tx.ChainId,
		Nonce:                  tx.Nonce,
		GasPrice:               tx.GasPrice,
		GasLimit:               tx.GasLimit,
		TxDataBalanceTransfer:  bf,
		TxDataValidatorDeposit: vd,
		TxDataAuditorDeposit:   ad,
		TxDataObjectCommit:     oc,
		TxDataObjectAudit:      oa,
		TxDataObjectPoR:        op,
	}
}

// CopyProposerSlashings copies the provided ProposerSlashing array.
func CopyProposerSlashings(slashings []*ProposerSlashing) []*ProposerSlashing {
	if slashings == nil {
		return nil
	}
	newSlashings := make([]*ProposerSlashing, len(slashings))
	for i, att := range slashings {
		newSlashings[i] = CopyProposerSlashing(att)
	}
	return newSlashings
}

// CopyProposerSlashing copies the provided ProposerSlashing.
func CopyProposerSlashing(slashing *ProposerSlashing) *ProposerSlashing {
	if slashing == nil {
		return nil
	}
	return &ProposerSlashing{
		Header_1: CopySignedBlockHeader(slashing.Header_1),
		Header_2: CopySignedBlockHeader(slashing.Header_2),
	}
}

// CopySignedBlockHeader copies the provided SignedBlockHeader.
func CopySignedBlockHeader(header *SignedBlockHeader) *SignedBlockHeader {
	if header == nil {
		return nil
	}
	return &SignedBlockHeader{
		Header:    CopyBlockHeader(header.Header),
		Signature: safeCopyBytes(header.Signature),
	}
}

// CopyBlockHeader copies the provided BlockHeader.
func CopyBlockHeader(header *BlockHeader) *BlockHeader {
	if header == nil {
		return nil
	}
	parentRoot := safeCopyBytes(header.ParentRoot)
	stateRoot := safeCopyBytes(header.StateRoot)
	bodyRoot := safeCopyBytes(header.BodyRoot)
	return &BlockHeader{
		Slot:          header.Slot,
		ProposerIndex: header.ProposerIndex,
		ParentRoot:    parentRoot,
		StateRoot:     stateRoot,
		BodyRoot:      bodyRoot,
	}
}

// CopyAttesterSlashings copies the provided AttesterSlashings array.
func CopyAttesterSlashings(slashings []*AttesterSlashing) []*AttesterSlashing {
	if slashings == nil {
		return nil
	}
	newSlashings := make([]*AttesterSlashing, len(slashings))
	for i, slashing := range slashings {
		newSlashings[i] = &AttesterSlashing{
			Attestation_1: CopyIndexedAttestation(slashing.Attestation_1),
			Attestation_2: CopyIndexedAttestation(slashing.Attestation_2),
		}
	}
	return newSlashings
}

// CopyIndexedAttestation copies the provided IndexedAttestation.
func CopyIndexedAttestation(indexedAtt *IndexedAttestation) *IndexedAttestation {
	var indices []uint64
	if indexedAtt == nil {
		return nil
	} else if indexedAtt.AttestingIndices != nil {
		indices = make([]uint64, len(indexedAtt.AttestingIndices))
		copy(indices, indexedAtt.AttestingIndices)
	}
	return &IndexedAttestation{
		AttestingIndices: indices,
		Data:             CopyAttestationData(indexedAtt.Data),
		Signature:        safeCopyBytes(indexedAtt.Signature),
	}
}

// CopyAttestations copies the provided Attestation array.
func CopyAttestations(attestations []*Attestation) []*Attestation {
	if attestations == nil {
		return nil
	}
	newAttestations := make([]*Attestation, len(attestations))
	for i, att := range attestations {
		newAttestations[i] = CopyAttestation(att)
	}
	return newAttestations
}

// CopyValidator copies the provided validator.
func CopyValidator(val *Validator) *Validator {
	pubKey := make([]byte, len(val.PublicKey))
	copy(pubKey, val.PublicKey)
	return &Validator{
		PublicKey:                  pubKey,
		Slashed:                    val.Slashed,
		ActivationEligibilityEpoch: val.ActivationEligibilityEpoch,
		ActivationEpoch:            val.ActivationEpoch,
		ExitEpoch:                  val.ExitEpoch,
		WithdrawableEpoch:          val.WithdrawableEpoch,
	}
}

func (a *Account) Copy() *Account {
	return proto.Clone(a).(*Account)
}
