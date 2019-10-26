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
