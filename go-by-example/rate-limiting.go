package main

import (
	"fmt"
	"time"
)

func rateLimiting() {
	requests := make(chan int, 10)
	limiter := time.Tick(time.Millisecond * 300)

	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)
	for t := range requests {
		<-limiter
		fmt.Println("request at", t, time.Now())
	}

	burstyLimiter := make(chan time.Time, 4)
	for range 4 {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 300) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 11)
	for l := 1; l <= 11; l++ {
		burstyRequests <- l
	}
	close(burstyRequests)

	for z := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request at ", z, time.Now())
	}
}
