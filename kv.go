// package sdk defined spanner sdk for tikv transection
package sdk

// Key give out spanner sdk key
type Key struct {
	optype string
}

func WarpKey(rk, op string) *Key {
	return &Key{}
}

// Value defined a byte slice for transection sdk value field
type Value []byte

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

func (s *KeyStore) Getall() ([]Key, []Value) {
	return nil, nil
}
