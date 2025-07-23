package linkquality

import (
	"sync"
	"time"
	"tridrasil/types"
)

type PeerLink struct {
	NodeID  types.NodeID
	Metrics types.LinkMetrics
	mu      sync.Mutex
}

func (pl *PeerLink) UpdateMetrics(latency time.Duration, packetLoss float64) {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	pl.Metrics.Latency = latency
	pl.Metrics.PacketLoss = packetLoss
	pl.Metrics.LastUpdate = time.Now()
}

func (pl *PeerLink) GetMetrics() types.LinkMetrics {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	return pl.Metrics
}

func CalculatePathCost(latency time.Duration, packetLoss float64, hops int) float64 {
	alpha := 1.0
	beta := 10.0
	gamma := 0.5
	return alpha*latency.Seconds() + beta*packetLoss + gamma*float64(hops)
}
