// package sdk defined spanner sdk for tikv transaction
package sdk

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	transactionHeartBeatInterval = 1 // 1 second  per heartbeat
)

// Pacemaker sent heartbeat to every perticapate nodes
type Pacemaker struct {
	txnid   uint64
	kvnodes []*tikvClient
	tick    <-chan time.Time
	ctx     context.Context

	sync.RWMutex
}

// NewPacemaker gen a pacemaker client with context, txnid
func NewPacemaker(c context.Context, txnid uint64) *Pacemaker {
	lg.Debug("new pacemaker with txnid", zap.Uint64("txnid", txnid))
	return &Pacemaker{
		ctx:     c,
		kvnodes: make([]*tikvClient, 0, 0xF),
		tick:    time.Tick(time.Duration(transactionHeartBeatInterval) * time.Second),
	}
}

// AddNode add new node
func (p *Pacemaker) AddNode(cli *tikvClient) {
	lg.Debug("add tikv client to txn", zap.Uint64("txnid", p.txnid))
	p.Lock()
	p.kvnodes = append(p.kvnodes, cli)
	p.Unlock()
}

// do heartbeat
func (p *Pacemaker) heartbeat() {
	for {
		select {
		case <-p.ctx.Done():
			lg.Debug("cancel heartbeat goroutine", zap.Uint64("txnid", p.txnid))
			return
		case <-p.tick:
			p.Lock()
			for i, _ := range p.kvnodes {
				// ttl shoud larger than hb interval, currently 2 times
				lg.Debug("do heartbeat", zap.Uint64("txnid", p.txnid))
				p.kvnodes[i].HeartBeat(p.ctx, p.txnid, uint64(2*transactionHeartBeatInterval))
			}
			p.Unlock()
		}
	}
}
