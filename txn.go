// package sdk defined spanner sdk for tikv transection
package sdk

var pdCluster []string

// init gen logger and pd kv client
func init() {
	initLogger()
	initPDclient(pdCluster)
}

// Transaction defiend spanner transection interface, implemented by ro transection and rw transection
type Transaction interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
	Delete(key []byte) error
	Rollback() error
	Commit() error
	Close() error
}

// Begin start a transection with option wether it is a RO transection
func Begin(ro bool) (Transaction, error) {
	if ro {
		return beginROTxn()
	}
	return beginRWTxn()
}

// base transection op
type transection struct {
}
