package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicCounters() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for range 50 {
		wg.Go(func() {
			for range 1000 {
				ops.Add(1)
			}
		})
	}
	wg.Wait()
	fmt.Println("Atomic counter: ", ops.Load())
}
