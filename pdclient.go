// package sdk defined spanner sdk for tikv transaction
package sdk

import (
	"os"
	"time"

	pd "github.com/pingcap/pd/client"
	"github.com/pingcap/tidb/store/tikv/oracle"
	"github.com/pingcap/tidb/store/tikv/oracle/oracles"
	"go.uber.org/zap"
)

var (
	pdClient     pd.Client
	oracleClient oracle.Oracle
)

func initPDclient(addrs []string) {
	var err error
	pdClient, err = pd.NewClient(addrs, pd.SecurityOption{
		CAPath:   "",
		CertPath: "",
		KeyPath:  "",
	})
	if err != nil {
		os.Exit(0)
	}
	lg.Info("init pd client", zap.Strings("addrs", addrs))

	oracleClient, err = oracles.NewPdOracle(pdClient, time.Second)
	if err != nil {
		os.Exit(0)
	}
}
