package reputation

import (
	"sync"

	"/types"
)

type ReputationManager struct {
	scores             map[types.NodeID]types.ReputationScore
	blacklistThreshold types.ReputationScore
	blacklist          map[types.NodeID]struct{}
	mu                 sync.Mutex
}

func NewReputationManager(threshold types.ReputationScore) *ReputationManager {
	return &ReputationManager{
		scores:             make(map[types.NodeID]types.ReputationScore),
		blacklist:          make(map[types.NodeID]struct{}),
		blacklistThreshold: threshold,
	}
}

func (rm *ReputationManager) AdjustScore(nodeID types.NodeID, delta int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.scores[nodeID] += types.ReputationScore(delta)
	if rm.scores[nodeID] <= rm.blacklistThreshold {
		rm.blacklist[nodeID] = struct{}{}
	}
}

func (rm *ReputationManager) GetScore(nodeID types.NodeID) types.ReputationScore {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	return rm.scores[nodeID]
}

func (rm *ReputationManager) IsBlacklisted(nodeID types.NodeID) bool {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	_, exists := rm.blacklist[nodeID]
	return exists
}

