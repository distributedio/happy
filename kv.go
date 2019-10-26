// package sdk defined spanner sdk for tikv transection
package sdk

// Key give out spanner sdk key
type Key string

// Value defined a byte slice for transection sdk value field
type Value struct {
	Type  Type
	Value []byte
}

func warpKV(rk, op string) *Opreation {
	return &Operation{}
}

// KeyStore store all keys locally
// TODO should be an interface
type KeyStore struct {
	store map[Key]Value
}

func (s *KeyStore) Set(k Key, v Value) {
}

func (s *KeyStore) Get(k Key) Value {
	return Value{}
}

func (s *KeyStore) Delete(k Key) {
}

func (s *KeyStore) GetAllOperations() []*Operation {
	return nil
}
