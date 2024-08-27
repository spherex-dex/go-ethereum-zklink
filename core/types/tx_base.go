package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BaseTx struct {
	Hash common.Hash
}

func (tx *BaseTx) hash() common.Hash {
	return tx.Hash
}
func (tx *BaseTx) copy() TxData {
	cpy := &BaseTx{
		Hash: tx.Hash,
	}
	return cpy
}
func (tx *BaseTx) txType() byte           { panic("txType called on BaseTx)") }
func (tx *BaseTx) chainID() *big.Int      { panic("chainID called on BaseTx)") }
func (tx *BaseTx) accessList() AccessList { panic("accessList called on BaseTx)") }
func (tx *BaseTx) data() []byte           { panic("data called on BaseTx)") }
func (tx *BaseTx) gas() uint64            { panic("gas called on BaseTx)") }
func (tx *BaseTx) gasPrice() *big.Int     { panic("gasPrice called on BaseTx)") }
func (tx *BaseTx) gasTipCap() *big.Int    { panic("gasTipCap called on BaseTx)") }
func (tx *BaseTx) gasFeeCap() *big.Int    { panic("gasFeeCap called on BaseTx)") }
func (tx *BaseTx) value() *big.Int        { panic("value called on BaseTx)") }
func (tx *BaseTx) nonce() uint64          { panic("nonce called on BaseTx)") }
func (tx *BaseTx) to() *common.Address    { panic("to called on BaseTx)") }
func (tx *BaseTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	panic("effectiveGasPrice called on BaseTx")
}
func (tx *BaseTx) rawSignatureValues() (v, r, s *big.Int) {
	panic("rawSignatureValues called on BaseTx")
}
func (tx *BaseTx) setSignatureValues(chainID, v, r, s *big.Int) {
	panic("setSignatureValues called on BaseTx")
}
func (tx *BaseTx) encode(*bytes.Buffer) error {
	panic("encode called on BaseTx")
}
func (tx *BaseTx) decode([]byte) error {
	panic("decode called on BaseTx)")
}
