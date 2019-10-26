// package sdk defined spanner sdk for tikv transection
package sdk

type roTxn struct {
}

func beginROTxn() (*roTxn, error) {
	return nil, nil
}

func (txn *roTxn) Get(key *Key) (Value, error) {
	return nil, nil
}

func (txn *roTxn) Set(key *Key, value Value) error
func (txn *roTxn) Delete(key *Key) error
func (txn *roTxn) Rollback() error
func (txn *roTxn) Commit() error
func (txn *roTxn) Close() error
