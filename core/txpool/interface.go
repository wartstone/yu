package txpool

import (
	. "github.com/yu-org/yu/common"
	. "github.com/yu-org/yu/core/types"
)

type ItxPool interface {
	// PoolSize return pool size of txpool
	PoolSize() uint64
	// txpool with the check-functions
	WithBaseChecks(checkFns []TxnCheck) ItxPool
	WithTripodChecks(checkFns []TxnCheck) ItxPool
	// base check txn
	BaseCheck(*SignedTxn) error
	TripodsCheck(stxn *SignedTxn) error
	// use for SyncTxns
	NecessaryCheck(stxn *SignedTxn) error
	// insert into txpool
	Insert(txn *SignedTxn) error
	// batch insert into txpool
	BatchInsert(txns SignedTxns) []error
	// package some txns to send to tripods
	Pack(blockHash Hash, numLimit uint64) ([]*SignedTxn, error)

	PackFor(blockHash Hash, numLimit uint64, filter func(txn *SignedTxn) bool) ([]*SignedTxn, error)

	// get unpacked txn
	GetTxn(hash Hash) (*SignedTxn, error)

	Packed(block Hash, txns []Hash) error
	// remove txns after execute all tripods
	Reset() error
}
