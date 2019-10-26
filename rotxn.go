// package sdk defined spanner sdk for tikv transaction
package sdk

type roTxn struct {
	transaction
}

func beginROTxn() (*roTxn, error) {
	return nil, nil
}

func (txn *roTxn) Get(key []byte) ([]byte, error) {
	lg.Debug("call get in ro transection")
	// get TS if not exist, and start heartbeat with transaction nodes
	if err := txn.setTxnIDNX(); err != nil {
		return nil, err
	}

	// locate key from PD
	_, rleader, err := pdClient.GetRegion(txn.ctx, key)
	if err != nil {
		return nil, err
	}
	linfo, err := pdClient.GetStore(txn.ctx, rleader.GetId())
	if err != nil {
		return nil, err
	}

	// call tikv get RO/RW with TS
	kvcli, err := NewTiKV(linfo.GetAddress())
	if err != nil {
		return nil, err
	}
	return kvcli.Get(txn.ctx, txn.txnid, key, txn.txnid, true)
}

func (txn *roTxn) Set(key, value []byte) error {
	return ErrInvalid
}
func (txn *roTxn) Delete(key []byte) error {
	return ErrInvalid
}
func (txn *roTxn) Rollback() error {
	return nil
}
func (txn *roTxn) Commit() error {
	return nil
}
func (txn *roTxn) Close() error {
	return nil
}
