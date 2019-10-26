// package sdk defined spanner sdk for tikv transaction
package sdk

import (
	"os"

	pd "github.com/pingcap/pd/client"
	"go.uber.org/zap"
)

var pdClient pd.Client

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
}
