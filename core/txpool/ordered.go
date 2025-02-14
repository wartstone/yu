package txpool

import (
	"github.com/sirupsen/logrus"
	. "github.com/yu-org/yu/common"
	. "github.com/yu-org/yu/core/types"
)

type orderedTxns struct {
	index map[Hash]int
	txns  []*SignedTxn
}

func newOrderedTxns() *orderedTxns {
	return &orderedTxns{
		index: make(map[Hash]int),
		txns:  make([]*SignedTxn, 0),
	}
}

func (ot *orderedTxns) exist(txn *SignedTxn) bool {
	_, exist := ot.index[txn.TxnHash]
	return exist
}

func (ot *orderedTxns) insert(input *SignedTxn) {
	if len(ot.txns) == 0 {
		ot.txns = []*SignedTxn{input}
		ot.index[input.TxnHash] = 0
	}
	for i, tx := range ot.txns {
		if input.Raw.Ecall.LeiPrice > tx.Raw.Ecall.LeiPrice {
			ot.txns = append(ot.txns[:i], append([]*SignedTxn{input}, ot.txns[i:]...)...)
			ot.index[input.TxnHash] = i
			return
		}
	}
}

func (ot *orderedTxns) delete(hash Hash) {
	if idx, ok := ot.index[hash]; !ok {
		return
	} else {
		logrus.Debugf("DELETE txn(%s) from txpool", hash.String())
		ot.txns = append(ot.txns[:idx], ot.txns[idx+1:]...)
		delete(ot.index, hash)
	}
}

func (ot *orderedTxns) deletes(hashes []Hash) {
	for _, hash := range hashes {
		ot.delete(hash)
	}
}

func (ot *orderedTxns) gets(numLimit uint64, filter func(txn *SignedTxn) bool) []*SignedTxn {
	txns := make([]*SignedTxn, 0)
	if numLimit > uint64(ot.len()) {
		numLimit = uint64(ot.len())
	}
	for _, txn := range ot.txns[:numLimit] {
		if filter(txn) {
			logrus.Debugf("Pack txn(%s) from Txpool", txn.TxnHash.String())
			txns = append(txns, txn)
		}
	}
	return txns
}

func (ot *orderedTxns) len() int {
	return len(ot.txns)
}
