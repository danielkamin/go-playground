package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type WeatherData struct {
	Latitude              float64 `json:"latitude"`
	Longitude             float64 `json:"longitude"`
	Elevation             float64 `json:"elevation"`
	Generationtime_ms     float64 `json:"generationtime_ms"`
	Utc_offset_seconds    int     `json:"utc_offset_seconds"`
	Timezone              string  `json:"timezone"`
	Timezone_abbreviation string  `json:"timezone_abbreviation"`
	Hourly                struct {
		Time        []string  `json:"time"`
		Temperature []float64 `json:"temperature_2m"`
	} `json:"hourly"`
	Hourly_units struct {
		Temperature_2m string `json:"temperature_2m"`
	} `json:"hourly_units"`
}

const openMeteoBaseUrl = "https://api.open-meteo.com/v1/forecast"

func fetchWeatherWorker(jobs <-chan string, result chan<- WeatherData) {
	for url := range jobs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error occured during getting data for url: %s, %+v", url, err)
			continue
		}

		weatherData := WeatherData{}
		err = json.NewDecoder(resp.Body).Decode(&weatherData)
		if err != nil {
			fmt.Printf("Error occured during decoding data for url: %s, %+v", url, err)
			continue
		}
		resp.Body.Close()
		result <- weatherData
	}

}
func generateData(locations map[string]Coordinates) <-chan WeatherData {

	jobs := make(chan string)
	result := make(chan WeatherData)
	var workerGroup sync.WaitGroup

	for i := 0; i < 3; i++ {
		workerGroup.Go(func() {
			fetchWeatherWorker(jobs, result)
		})
	}
	go func() {
		for _, coord := range locations {
			jobs <- fmt.Sprintf("%s?latitude=%f&longitude=%f&hourly=temperature_2m", openMeteoBaseUrl, coord.lat, coord.long)
		}
		close(jobs)
	}()
	go func() {
		workerGroup.Wait()
		close(result)
	}()

	return result
}
