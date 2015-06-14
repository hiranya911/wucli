package main

import (
	"fmt"
	"os"
	"strings"
	"wu"
)

func main() {
	argsWithoutProg := os.Args[1:]
	var location string
	if len(argsWithoutProg) == 0 {
		location = "autoip"
	} else if len(argsWithoutProg) == 1 {
		location = argsWithoutProg[0]
	} else {
		fmt.Println("Usage: wucli [location]")
		return
	}

	wd, err := wu.Query(location)
	if err != nil {
		fmt.Println(err)
		return
	}

	locationResults := wd.Response.Results
	if locationResults != nil {
		fmt.Println("Locations matched:")
		for _, lr := range locationResults {
			if lr.State != "" {
				fmt.Printf("  * %s, %s, %s (zmw:%s)\n", lr.City, lr.State, lr.CountryName, lr.Zmw)
			} else {
				fmt.Printf("  * %s, %s (zmw:%s)\n", lr.City, lr.CountryName, lr.Zmw)
			}
		}
		return
	}

	current := wd.CurrentObservation
	loc := current.DisplayLocation
	forecast := wd.Forecast.TxtForecast

	title := fmt.Sprintf("Location: %s (long: %s, lat: %s)", loc.Full, loc.Longitude, loc.Latitude)
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", len(title)))

	fmt.Printf("Summary: %s\n", current.Weather)
	fmt.Printf("Temperature: %s\n", current.TemperatureString)
	fmt.Printf("Feels like: %s\n", current.FeelslikeString)
	fmt.Printf("Wind: %s\n", current.WindString)
	fmt.Printf("Wind chill: %s\n", current.WindchillString)
	fmt.Printf("Precipitation: %s\n", current.PrecipTodayString)

	fmt.Println("Weather Forecast:")
	for _, fc := range forecast.Forecastday {
		fmt.Printf("  * %s: %s\n", fc.Title, fc.Fcttext)
	}

	fmt.Printf("\n%s\n", current.ObservationTime)
}
