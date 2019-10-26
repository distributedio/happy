// package sdk defined spanner sdk for tikv transection
package sdk

// Value defined a byte slice for transection sdk value field
type Value []byte

// Transection defiend spanner transection interface, implemented by ro transection and rw transection
type Transection interface {
	Get(key *Key) (Value, error)
	Set(key *Key, value Value) error
	Delete(key *Key) error
	Rollback() error
	Commit() error
}

// Begin start a transection with option wether it is a RO transection
func Begin(ro bool) (Transection, error) {
	if ro {
		return beginROTxn()
	}
	return beginRWTxn()
}
