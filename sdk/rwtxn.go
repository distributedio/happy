// package sdk defined spanner sdk for tikv transection
package sdk

type rwTxn struct {
}

func beginRWTxn() (*rwTxn, error) {
	return nil, nil
}

func (txn *rwTxn) Get(key *Key) (Value, error) {
	return nil, nil
}
