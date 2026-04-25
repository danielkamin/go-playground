package main

import (
	"fmt"
	"time"
)

func closingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	time.Sleep(time.Second * 2)
	close(jobs)
	time.Sleep(time.Second * 2)

	fmt.Println("sent all jobs")

	<-done
	time.Sleep(time.Second * 5)

	_, ok := <-jobs
	fmt.Println("received more jobs: ", ok)
}
