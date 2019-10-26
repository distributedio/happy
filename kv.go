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

func NewKVStore() *KeyStore {
}

// KeyStore store all keys locally
// TODO should be an interface
type KeyStore struct {
	sync.Mutex
	store map[Key]Value
}

func (s *KeyStore) Set(k Key, v Value) {
	s.Lock()
	s.store[k] = v
	s.Unlock()
}

func (s *KeyStore) Get(k Key) Value {
	s.Lock()
	defer s.Unlock()
	return s.store[k]
}

func (s *KeyStore) Delete(k Key) {
	lg.Debug("call delete in kv store", zap.Stringer("key", k))
	s.Lock()
	delete(s.store, k)
	s.Unlock()
}

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
