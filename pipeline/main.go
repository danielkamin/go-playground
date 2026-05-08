package main

import "fmt"

type Coordinates struct {
	lat  float64
	long float64
}

var locations = map[string]Coordinates{
	"bialystok": {
		lat:  53.1333,
		long: 23.1643,
	},
	"warsaw": {
		lat:  52.2298,
		long: 21.0118,
	},
	"krakow": {
		lat:  50.0614,
		long: 19.9366,
	},
}

func main() {
	chanel := generateData(locations)
	for c := range chanel {
		fmt.Println(c)
	}
}
