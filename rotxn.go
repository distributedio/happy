// package sdk defined spanner sdk for tikv transaction
package sdk

type roTxn struct {
}

func beginROTxn() (*roTxn, error) {
	return nil, nil
}

func (txn *roTxn) Get(key []byte) ([]byte, error)
func (txn *roTxn) Set(key, value []byte) error
func (txn *roTxn) Delete(key []byte) error
func (txn *roTxn) Rollback() error
func (txn *roTxn) Commit() error
func (txn *roTxn) Close() error
