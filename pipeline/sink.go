package main

import "fmt"

func sink(temp <-chan float64) {
	for t := range temp {
		fmt.Printf("%f \n", t)
	}
}
