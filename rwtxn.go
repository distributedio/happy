// package sdk defined spanner sdk for tikv transaction
package sdk

type rwTxn struct {
}

func beginRWTxn() (*rwTxn, error) {
	return nil, nil
}

func (txn *rwTxn) Get(key []byte) ([]byte, error)
func (txn *rwTxn) Set(key, value []byte) error
func (txn *rwTxn) Delete(key []byte) error
func (txn *rwTxn) Rollback() error
func (txn *rwTxn) Commit() error
func (txn *rwTxn) Close() error
