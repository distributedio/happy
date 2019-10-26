// package sdk defined spanner sdk for tikv transaction
package sdk

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
func Begin(ro bool) (Transaction, error) {
	if ro {
		return beginROTxn()
	}
	return beginRWTxn()
}

// base transaction op
type transaction struct {
}
