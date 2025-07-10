package ratelimit

import (
	"sync"
	"time"
)

type TokenBucket struct {
	tokens     int
	maxTokens  int
	refillRate int // tokens per second
	lastRefill time.Time
	mu         sync.Mutex
}

func NewTokenBucket(maxTokens, refillRate int) *TokenBucket {
	return &TokenBucket{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	now := time.Now()
	elapsed := int(now.Sub(tb.lastRefill).Seconds())
	if elapsed > 0 {
		tb.tokens += elapsed * tb.refillRate
		if tb.tokens > tb.maxTokens {
			tb.tokens = tb.maxTokens
		}
		tb.lastRefill = now
	}
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

