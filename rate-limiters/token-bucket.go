package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens     int64
	capacity   int64
	rate       int64
	lastRefill int64
	mu         sync.Mutex
}

func NewTokenBucket(rate, capacity int64) *TokenBucket {
	return &TokenBucket{
		tokens:     capacity,
		capacity:   capacity,
		rate:       rate,
		lastRefill: time.Now().UnixNano(),
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now().UnixNano()
	elapsed := now - tb.lastRefill
	newTokens := (elapsed * tb.rate) / 1e9
	if newTokens > 0 {
		tb.lastRefill = now
		tb.tokens = min(tb.capacity, tb.tokens+newTokens)
	}
}

func (tb *TokenBucket) Take() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.refill()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func RunTokenBucket() {
	tb := NewTokenBucket(2, 20)
	for i := 0; i < 25; i++ {
		if tb.Take() {
			fmt.Println("Request: ", i, " passed!")
		} else {
			fmt.Println("Request: ", i, " blocked!")
		}
		time.Sleep(50 * time.Millisecond)
	}
}
