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
func warpKV(k Key, v Value) *Operation {
	return &Operation{
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

// Set k v in local store
func (s *KeyStore) Set(k Key, v Value) {
	lg.Debug("call set kv store", zap.Stringer("key", k), zap.Stringer("value", v))
	s.Lock()
	s.store[k] = v
	s.Unlock()
}

// Get k from local store
func (s *KeyStore) Get(k Key) Value {
	lg.Debug("call get kv store", zap.Stringer("key", k))
	s.Lock()
	defer s.Unlock()
	return s.store[k]
}

// Delete k from local store
func (s *KeyStore) Delete(k Key) {
	lg.Debug("call delete kv store", zap.Stringer("key", k))
	s.Lock()
	delete(s.store, k)
	s.Unlock()
}

// GetAllOperations return all k-v pairs store in kvstore in operation set
func (s *KeyStore) GetAllOperations() []*Operation {
	resultSet := make([]*Operation, 0)
	s.Lock()
	defer s.Unlock()
	for k, v := range s.store {
		resultSet = append(resultSet, warpKV(k, v))
	}
	lg.Debug("call get all operations", zap.Int("len", len(resultSet)))
	return resultSet
}
