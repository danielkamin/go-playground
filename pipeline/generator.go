package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		Temperature []float64 `json:"temperature"`
	} `json:"hourly"`
	Hourly_units struct {
		Temperature_2m string `json:"temperature_2m"`
	} `json:"hourly_units"`
}

const openMeteoBaseUrl = "https://api.open-meteo.com/v1/forecast"

func fetchWeather(url string) (WeatherData, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Errorf("Error occured during getting data for url: %s, %+v", url, err)
		return WeatherData{}, err
	}
	weatherData := WeatherData{}
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	return weatherData, err
}
func generateData(locations map[string]Coordinates) <-chan WeatherData {
	out := make(chan WeatherData)
	go func() {
		defer close(out)
		for location, coord := range locations {
			fmt.Printf("Getting data for location: %s \n", location)
			weatherData, err := fetchWeather(fmt.Sprintf("%s?latitude=%f&longitude=%f&hourly=temperature_2m", openMeteoBaseUrl, coord.lat, coord.long))
			if err != nil {
				fmt.Errorf("Error occured during parsing data for location: %s, %+v", location, err)
			}
			fmt.Printf("data: %+v \n", weatherData)
			out <- weatherData
		}
	}()
	return out
}
