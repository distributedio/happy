// package sdk defined spanner sdk for tikv transaction
package sdk

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// Key give out spanner sdk key
type Key string

func (k Key) String() string { return string(k) }

// Value defined a byte slice for transaction sdk value field
type Value struct {
	Type  Type
	Value []byte
}

func (v Value) String() string {
	return fmt.Sprintf("value: %v, type: %v", string(v.Value), v.Type)
}

// warpKV to operation for pd client
func warpKV(k Key, v Value) Operation {
	return Operation{
		Type:  v.Type,
		Key:   []byte(k),
		Value: v.Value,
	}
}

// KeyStore store all keys locally
// TODO should be an interface
type KeyStore struct {
	sync.Mutex
	store map[Key]Value
}

// NewKVStore return a local in memory kv storage
func NewKVStore() *KeyStore {
	return &KeyStore{
		store: make(map[Key]Value),
	}
}

func (s *KeyStore) SetOp(k, v []byte, op Type) {
	s.Lock()
	s.store[Key(k)] = Value{
		Type:  OpPut,
		Value: v,
	}
	s.Unlock()
}

// GetAllOperations return all k-v pairs store in kvstore in operation set
func (s *KeyStore) GetAllOperations() []Operation {
	resultSet := make([]Operation, 0)
	s.Lock()
	defer s.Unlock()
	for k, v := range s.store {
		resultSet = append(resultSet, warpKV(k, v))
	}
	lg.Debug("call get all operations", zap.Int("len", len(resultSet)))
	return resultSet
}
