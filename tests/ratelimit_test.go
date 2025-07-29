package tests

import (
	"testing"
	"time"

	"tridrasil/plugins/ratelimit"
)

func TestRateLimitTokenBucket(t *testing.T) {
	tb := ratelimit.NewTokenBucket(2, 1)

	if !tb.Allow() {
		t.Errorf("First request should be allowed")
	}
	if !tb.Allow() {
		t.Errorf("Second request should be allowed")
	}
	if tb.Allow() {
		t.Errorf("Third immediate request should be blocked")
	}

	time.Sleep(1 * time.Second)
	if !tb.Allow() {
		t.Errorf("Request after refill should be allowed")
	}
}
