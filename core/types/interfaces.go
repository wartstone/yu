package types

import (
	. "github.com/yu-org/yu/common"
	. "github.com/yu-org/yu/core/result"
)

//type IBlock interface {
//	IHeader
//	GetHeader() IHeader
//
//	GetBlockId() BlockId
//	GetTxnsHashes() []Hash
//	SetTxnsHashes(hashes []Hash)
//
//	SetHash(hash Hash)
//	SetPreHash(hash Hash)
//	SetTxnRoot(hash Hash)
//	SetStateRoot(hash Hash)
//	SetHeight(BlockNum)
//	SetTimestamp(ts uint64)
//	SetPeerID(peer.ID)
//
//	SetLeiLimit(e uint64)
//	UseLei(e uint64)
//
//	SetSignature([]byte)
//
//	Encode() ([]byte, error)
//	Decode(data []byte) (IBlock, error)
//
//	CopyFrom(other IBlock)
//}
//
//type IHeader interface {
//	GetHeight() BlockNum
//	GetHash() Hash
//	GetPrevHash() Hash
//	GetTxnRoot() Hash
//	GetStateRoot() Hash
//	GetTimestamp() uint64
//	GetPeerID() peer.ID
//	GetLeiLimit() uint64
//	GetLeiUsed() uint64
//
//	GetSignature() []byte
//}

// --------------- blockchain interface ----------------

type ConvergeType int

const (
	Longest ConvergeType = iota
	Heaviest
	Finalize
)

type IBlockChain interface {
	ConvergeType() ConvergeType

	NewEmptyBlock() *CompactBlock

	EncodeBlocks(blocks []*CompactBlock) ([]byte, error)
	DecodeBlocks(data []byte) ([]*CompactBlock, error)

	GetGenesis() (*CompactBlock, error)
	SetGenesis(b *CompactBlock) error

	AppendBlock(b *CompactBlock) error
	GetBlock(blockHash Hash) (*CompactBlock, error)
	ExistsBlock(blockHash Hash) bool
	UpdateBlock(b *CompactBlock) error

	Children(prevBlockHash Hash) ([]*CompactBlock, error)
	Finalize(blockHash Hash) error
	LastFinalized() (*CompactBlock, error)
	GetEndBlock() (*CompactBlock, error)
	GetAllBlocks() ([]*CompactBlock, error)

	GetRangeBlocks(startHeight, endHeight BlockNum) ([]*CompactBlock, error)
}

type IyuDB interface {
	GetTxn(txnHash Hash) (*SignedTxn, error)
	SetTxn(stxn *SignedTxn) error
	ExistTxn(hash Hash) bool

	Packs(block Hash, txns []Hash) error
	Pack(block, txn Hash) error

	GetAllUnpackedTxns() (txns []*SignedTxn, err error)

	GetTxns(blockHash Hash) ([]*SignedTxn, error)
	SetTxns(blockHash Hash, txns []*SignedTxn) error

	GetEvents(blockHash Hash) ([]*Event, error)
	SetEvents(events []*Event) error

	GetErrors(blockHash Hash) ([]*Error, error)
	SetError(err *Error) error
}
