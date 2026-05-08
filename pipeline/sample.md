Generator:    dla listy miast → fetchuje temperaturę z API → chan CityWeather
Transformer:  filtruje tylko miasta > 15°C             → chan CityWeather  
Sink:         wypisuje wyniki / zapisuje do pliku




// Generator — nie ma wejścia, zwraca kanał
func generate(cities []City) <-chan CityWeather {
    out := make(chan CityWeather)
    go func() {
        defer close(out)
        for _, c := range cities {
            // fetch API...
            out <- result
        }
    }()
    return out
}

// Transformer — bierze kanał, zwraca kanał
func filterWarm(in <-chan CityWeather) <-chan CityWeather {
    out := make(chan CityWeather)
    go func() {
        defer close(out)
        for w := range in {
            if w.Temp > 15 {
                out <- w
            }
        }
    }()
    return out
}

// Sink — konsumuje, nic nie zwraca
func printResults(in <-chan CityWeather) {
    for w := range in {
        fmt.Printf("%s: %.1f°C\n", w.Name, w.Temp)
    }
}