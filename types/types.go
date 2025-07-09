package types

import (
	"time"
)

// NodeID is a unique identifier for a peer.
type NodeID string

type ReputationScore int

type LinkMetrics struct {
	Latency    time.Duration
	PacketLoss float64
	LastUpdate time.Time
}

