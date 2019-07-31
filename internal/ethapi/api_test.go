package ethapi

import (
	"math/big"
	"testing"

	"github.com/DATx-Protocol/go-DATx/common/hexutil"
	"github.com/DATx-Protocol/go-DATx/core/types"
)

func TestToTransaction(t *testing.T) {
	nonce := uint64(0)
	args := &SendTxArgs{
		Type:     types.LoginCandidate,
		Nonce:    (*hexutil.Uint64)(&nonce),
		Gas:      (*hexutil.Big)(big.NewInt(0)),
		GasPrice: (*hexutil.Big)(big.NewInt(0)),
		Value:    (*hexutil.Big)(big.NewInt(0)),
		To:       nil,
	}
	tx := args.toTransaction()
	if tx.To() != nil {
		t.Errorf("transaction receiptent nil is expected, but got %x", tx.To())
	}
}
