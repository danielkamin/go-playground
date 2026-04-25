package main

import (
	"fmt"
	"time"
)

func timers() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	//time.Sleep(2 * time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
