package main

import "math"

func maxTempDiff(weatherData <-chan WeatherData) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for wd := range weatherData {
			temps := wd.Hourly.Temperature
			if len(temps) < 2 {
				continue
			}
			maxDiff := 0.00
			for i := 1; i < len(temps); i++ {
				diff := math.Abs(temps[i] - temps[i-1])
				if diff > maxDiff {
					maxDiff = diff
				}
			}
			out <- maxDiff
		}
	}()
	return out
}
