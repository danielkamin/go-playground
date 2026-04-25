package main

import (
	"fmt"
	"time"
)

func rangeOverChannels() {
	/*
		wpisywanie są wartości do kanału
		sleep
		zamkniecie kanalu ale poc?
		sleep
		wypisywanie elementow z 2sek interwalem
	*/

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	//wpisywanie
	time.Sleep(time.Second * 2)
	//blok do zamkniecia
	close(queue)
	time.Sleep(time.Second * 2)
	//oczekiwanie po zamknieciu
	for elem := range queue {
		//wypisywanie z przerwą
		fmt.Println(elem)
		time.Sleep(time.Second * 2)
	}
}
