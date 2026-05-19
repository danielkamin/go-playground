package main

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
	"poznan": {
		lat:  53.1333,
		long: 23.1643,
	},
	"lomza": {
		lat:  52.2298,
		long: 21.0118,
	},
	"zambrow": {
		lat:  50.0614,
		long: 19.9366,
	},
	"wroclaw": {
		lat:  52.2298,
		long: 21.0118,
	},
}

func main() {
	weatherChanel := generateData(locations)
	diffChanel := maxTempDiff(weatherChanel)
	sink(diffChanel)

}
