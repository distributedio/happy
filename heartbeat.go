// package sdk defined spanner sdk for tikv transection
package sdk

import "time"

// Pacemaker sent heartbeat to every perticapate nodes
type Pacemaker struct {
	cli      string
	nodes    []string
	interval time.Duration
	ctx      chan int
}

// NewPacemaker gen a pacemaker client
func NewPacemaker(pdcli string) *Pacemaker {
	return &Pacemaker{}
}

// Close stop heat the node
func (p *Pacemaker) Close() {}

// AddNode add new node
func (p *Pacemaker) AddNode() {}

// do heartbeat
func (p *Pacemaker) heartbeat() {
}
