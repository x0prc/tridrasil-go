package tests

import (
	"testing"
	"tridrasil/plugins/reputation"
	"tridrasil/types"
)

func TestReputationScoring(t *testing.T) {
	rm := reputation.NewReputationManager(-10)
	node := types.NodeID("peer1")

	rm.AdjustScore(node, -5)
	if rm.GetScore(node) != -5 {
		t.Errorf("expected score -5, got %d", rm.GetScore(node))
	}

	rm.AdjustScore(node, -10)
	if !rm.IsBlacklisted(node) {
		t.Error("expected peer to be blacklisted")
	}
}
