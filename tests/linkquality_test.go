package tests

import (
	"testing"
	"time"

	"tridrasil/plugins/linkquality"
)

func TestCalculatePathCost(t *testing.T) {
	latency := 100 * time.Millisecond
	packetLoss := 0.05
	hops := 2

	cost := linkquality.CalculatePathCost(latency, packetLoss, hops)
	if cost <= 0 {
		t.Errorf("expected positive cost, got %f", cost)
	}
}

func TestLinkMetricsUpdate(t *testing.T) {
	pl := &linkquality.PeerLink{}
	pl.UpdateMetrics(133*time.Millisecond, 0.1)
	metrics := pl.GetMetrics()

	if metrics.Latency != 133*time.Millisecond {
		t.Errorf("expected latency 133ms, got %v", metrics.Latency)
	}
	if metrics.PacketLoss != 0.1 {
		t.Errorf("expected packet loss 0.1, got %f", metrics.PacketLoss)
	}
}
