package main

import (
	"fmt"
	"time"
)

type LeakyBucket struct {
	rate     int64
	capacity int64
	queue    chan struct{}
	stopCh   chan struct{}
}

func NewLeakyBucket(rate, capacity int64) *LeakyBucket {
	lb := &LeakyBucket{
		rate:     rate,
		capacity: capacity,
		queue:    make(chan struct{}, capacity),
		stopCh:   make(chan struct{}),
	}
	go lb.leak()
	return lb
}

func (lb *LeakyBucket) Allow() bool {
	select {
	case lb.queue <- struct{}{}:
		return true
	default:
		return false
	}
}

func (lb *LeakyBucket) leak() {
	//ticker to indicate rate of adding new tokens and iterating over requests if any are in a queue
	ticker := time.NewTicker(time.Second / time.Duration(lb.rate))
	defer ticker.Stop()

	for {
		select {
		case <-lb.stopCh:
			return
		case <-ticker.C:
			select {
			case <-lb.queue: //do something with request
			default: //drop it or idk
			}
		}
	}
}

func (lb *LeakyBucket) Stop() { //graceful shutdown
	close(lb.stopCh)
}

func RunLeakyBucket() {
	leakyBucket := NewLeakyBucket(5, 10)
	defer leakyBucket.Stop()

	for i := 1; i <= 50; i++ {
		if leakyBucket.Allow() {
			fmt.Println("Request ", i, " queued")
		} else {
			fmt.Println("Request ", i, " dropped")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
