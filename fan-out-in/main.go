package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		randN := rand.Intn(10)
		fmt.Println("Worker: ", j, " doing job for rand: ", randN)
		time.Sleep(time.Second * time.Duration(randN))
		results <- j
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)
	var workerGroup sync.WaitGroup

	for i := 0; i < 3; i++ {
		workerGroup.Go(func() {
			worker(jobs, results)
		})
	}
	go func() {
		workerGroup.Wait()
		close(results)
	}()
	var doneGroup sync.WaitGroup
	doneGroup.Add(1)
	go func() {
		defer doneGroup.Done()
		for res := range results {
			fmt.Println("Result from: ", res)
		}
	}()

	var jobsGroup sync.WaitGroup
	for j := 0; j < 6; j++ {
		jobsGroup.Go(func() {
			jobs <- j
		})
	}
	jobsGroup.Wait()
	close(jobs)

	doneGroup.Wait()
}
