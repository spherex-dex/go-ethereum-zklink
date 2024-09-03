package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type UnknownTx struct {
	Nonce    uint64          // nonce of sender account
	GasPrice *big.Int        // wei per gas
	Gas      uint64          // gas limit
	To       *common.Address `rlp:"nil"` // nil means contract creation
	Value    *big.Int        // wei amount
	Data     []byte          // contract invocation input data
	V, R, S  *big.Int        // signature values
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *UnknownTx) copy() TxData {
	cpy := &UnknownTx{
		Nonce: tx.Nonce,
		To:    copyAddressPtr(tx.To),
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are initialized below.
		Value:    new(big.Int),
		GasPrice: new(big.Int),
		V:        new(big.Int),
		R:        new(big.Int),
		S:        new(big.Int),
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.
func (tx *UnknownTx) txType() byte           { return UnknownTxType }
func (tx *UnknownTx) chainID() *big.Int      { return deriveChainId(tx.V) }
func (tx *UnknownTx) accessList() AccessList { return nil }
func (tx *UnknownTx) data() []byte           { return tx.Data }
func (tx *UnknownTx) gas() uint64            { return tx.Gas }
func (tx *UnknownTx) gasPrice() *big.Int     { return tx.GasPrice }
func (tx *UnknownTx) gasTipCap() *big.Int    { return tx.GasPrice }
func (tx *UnknownTx) gasFeeCap() *big.Int    { return tx.GasPrice }
func (tx *UnknownTx) value() *big.Int        { return tx.Value }
func (tx *UnknownTx) nonce() uint64          { return tx.Nonce }
func (tx *UnknownTx) to() *common.Address    { return tx.To }

func (tx *UnknownTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	return dst.Set(tx.GasPrice)
}

func (tx *UnknownTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *UnknownTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.V, tx.R, tx.S = v, r, s
}

func (tx *UnknownTx) encode(*bytes.Buffer) error {
	panic("encode called on UnknownTx")
}

func (tx *UnknownTx) decode([]byte) error {
	panic("decode called on UnknownTx)")
}
