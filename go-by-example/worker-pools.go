package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, request chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second * 2)
		fmt.Println("worker", id, "finished job", j)
		request <- j * 2
	}
}

func workerPools() {
	numJobs := 20
	results := make(chan int, numJobs)
	jobs := make(chan int, numJobs)

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
