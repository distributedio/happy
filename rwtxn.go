// package sdk defined spanner sdk for tikv transaction
package sdk

type rwTxn struct {
	transaction
}

func beginRWTxn() (*rwTxn, error) {
	return nil, nil
}

func (txn *rwTxn) Get(key []byte) ([]byte, error) {
	lg.Debug("call get in rw transection")
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
	txn.pm.AddNode(kvcli, rleader.GetId())
	v, err := kvcli.Get(txn.ctx, txn.txnid, key, txn.txnid, false)
	if err != nil {
		return nil, err
	}
	txn.store.SetOp(key, v, OpRead)
	return v, nil
}

func (txn *rwTxn) Set(key, value []byte) error {
	lg.Debug("call set in rw transection")
	// get TS if not exist, and start heartbeat with transaction nodes
	if err := txn.setTxnIDNX(); err != nil {
		return err
	}

	// locate key from PD
	_, rleader, err := pdClient.GetRegion(txn.ctx, key)
	if err != nil {
		return err
	}
	linfo, err := pdClient.GetStore(txn.ctx, rleader.GetId())
	if err != nil {
		return err
	}

	// call tikv get RO/RW with TS
	kvcli, err := NewTiKV(linfo.GetAddress())
	if err != nil {
		return err
	}
	txn.pm.AddNode(rleader.GetId(), kvcli)
	txn.store.SetOp(key, value, OpPut)
	return nil
}

func (txn *rwTxn) Delete(key []byte) error {
	lg.Debug("call set in rw transection")
	// get TS if not exist, and start heartbeat with transaction nodes
	if err := txn.setTxnIDNX(); err != nil {
		return err
	}

	// locate key from PD
	_, rleader, err := pdClient.GetRegion(txn.ctx, key)
	if err != nil {
		return err
	}
	linfo, err := pdClient.GetStore(txn.ctx, rleader.GetId())
	if err != nil {
		return err
	}

	// call tikv get RO/RW with TS
	kvcli, err := NewTiKV(linfo.GetAddress())
	if err != nil {
		return err
	}
	txn.pm.AddNode(rleader.GetId(), kvcli)
	txn.store.SetOp(key, nil, OpDelete)
	return nil
}

func (txn *rwTxn) Rollback() error {
	return ErrNotImplemented
}

func (txn *rwTxn) Commit() error {
	lg.Debug("call commit in rw transection")
	if err := txn.setTxnIDNX(); err != nil {
		return err
	}

	kvclis, kvids := txn.pm.AllNodes()
	//Commit(ctx context.Context, txnID uint64, operations []Operation, coordinatorId []byte, participants [][]byte)
	_, err := kvclis[0].Commit(txn.ctx, txn.txnid, txn.store.GetAllOperations(), kvids[0], kvids[1:])
	return err
}

func (txn *rwTxn) Close() error {
	return ErrNotImplemented
}
