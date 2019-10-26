// package sdk defined spanner sdk for tikv transection
package sdk

import "time"

// Pacemaker sent heartbeat to every perticapate nodes
type Pacemaker struct {
	cli      string
	kvnodes  []*tikvClient
	interval time.Duration
	ctx      chan int
}

// NewPacemaker gen a pacemaker client
func NewPacemaker(pdcli string) *Pacemaker {
	return &Pacemaker{}
}

// close stop heat the node
func (p *Pacemaker) close() {}

// AddNode add new node
func (p *Pacemaker) AddNode() {}

// do heartbeat
func (p *Pacemaker) heartbeat() {}
