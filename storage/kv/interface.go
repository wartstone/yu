package kv

import (
	"github.com/pkg/errors"
)

var NoKvdbType = errors.New("no kvdb type")

type KV interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
	Iter(key []byte) (Iterator, error)
}

func NewKV(kvType string, cfg *KVconf) (KV, error) {
	switch kvType {
	case "badger":
		return NewBadger(cfg.path)
	case "bolt":
		return NewBolt(cfg.path)
	case "tikv":
		return NewTiKV(cfg.path)

	default:
		return nil, NoKvdbType
	}
}

type Iterator interface {
	Valid() bool
	Next() error
	Entry() ([]byte, []byte, error)
	Close()
}
