// package sdk defined spanner sdk for tikv transaction
package sdk

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

var pdCluster []string

// init gen logger and pd kv client
func init() {
	initLogger()
	initPDclient(pdCluster)
}

// Transaction defiend spanner transaction interface, implemented by ro transaction and rw transaction
type Transaction interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
	Delete(key []byte) error
	Rollback() error
	Commit() error
	Close() error
}

// Begin start a transaction with option wether it is a RO transaction
func Begin(ro bool, ctx context.Context) (Transaction, error) {
	if ro {
		return beginROTxn()
	}
	return beginRWTxn()
}

// base transaction op
type transaction struct {
	ctx     context.Context
	store   KeyStore
	pm      Pacemaker
	isValid bool

	sync.Mutex
	txnid uint64
}

func (txn *transaction) setTxnIDNX() error {
	var err error
	if txn.txnid == 0 {
		txn.Lock()
		defer txn.Unlock()
		txn.txnid, err = oracleClient.GetTimestamp(txn.ctx)
		if err != nil {
			lg.Error("error in call oracle get ts", zap.Error(err))
			return err
		}
	}
	txn.pm = NewPacemaker(txn.ctx, txn.txnid)
	return nil
}
