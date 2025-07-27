package tests

import (
	"testing"
	"time"

	"tridrasil/plugins/ratelimit"
)

func TestRateLimitTokenBucket(t *testing.T) {
	tb := ratelimit.NewTokenBucket(2, 1)

	if !tb.Allow() || !tb.Allow() {
		t.Errorf("Allow first two requests")
	} 

	if tb.Allow() {
		t.Errorf("Block third immediate request")
	}

	time.Sleep(1 * time.Second)
	if !tb.Allow() {
		t.Errorf("Allow after refill")
	}
}
