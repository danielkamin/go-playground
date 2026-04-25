package main

import (
	"fmt"
	"sync"
	"time"
)

func wgWorker(id int) {
	fmt.Printf("Worker %d started job\n", id)
	time.Sleep(time.Second * 2)
	fmt.Printf("Worker %d finished job\n", id)
}
func waitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			wgWorker(i)
		})
	}
	wg.Wait()
}
