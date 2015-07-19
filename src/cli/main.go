package main

import (
	"fmt"
	"os"
	"strings"
	"wu"
)

func printUsage() {
	fmt.Println("NAME")
	fmt.Println("\twucli -- Weather Underground Command-line Interface\n")

	fmt.Println("USAGE")
	fmt.Println("\twucli [location/keyword]\n")

	fmt.Println("DESCRIPTION")
	fmt.Println("\tRun wucli without any arguments to check the weather of your current location (determined")
	fmt.Println("\tautomatically from your IP address).\n")

	fmt.Println("\tOptionally you can run wucli with a location argument to get weather data related to a specific")
	fmt.Println("\tcity or area. Following types of location arguments are supported as of now:\n")

	fmt.Println("\t  * US state/city [e.g. CA/San Francisco]")
	fmt.Println("\t  * US zipcode [e.g. 93106]")
	fmt.Println("\t  * country/city [e.g. Sri Lanka/Colombo]")
	fmt.Println("\t  * latitude,longitude [e.g. 37.8,-122.4]")
	fmt.Println("\t  * airport code [e.g. KJFK]\n")

	fmt.Println("\tYou can also enter a simple keyword or a phrase as a search term [e.g. Belgrade]. In this case")
	fmt.Println("\twucli will return a list of matching locations with their unique zmw codes. You may rerun wucli")
	fmt.Println("\twith a zmw code as a location argument, to indicate which location that you are interested in.\n")

	fmt.Println("\tIf the location argument or the keyword contains spaces, enclose the argument within quotes.")
	fmt.Println("\tAlternatively, you can replace the spaces with underscores (_).\n")

	fmt.Println("CREDITS")
	fmt.Println("\tDeveloped by: Hiranya Jayathilaka and Jason Clark (mithereal@gmail.com)")
	fmt.Println("\tRemote API: http://www.wunderground.com\n")
}

func main() {
	argsWithoutProg := os.Args[1:]
	var location string
	var feature string
	if len(argsWithoutProg) == 0 {
		location = "autoip"
		feature = "conditions,forecast"
	} else if len(argsWithoutProg) == 1 {
		location = argsWithoutProg[0]
		feature = "conditions,forecast"
	} else if len(argsWithoutProg) == 2 {
		location = argsWithoutProg[0]
		feature = argsWithoutProg[1]
	} else {
		printUsage()
		return
	}

	if location == "-h" || location == "help" {
		printUsage()
		return
	}

	wd, err := wu.Query(location,feature)
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

	fmt.Println("\nWeather Forecast:")
	for _, fc := range forecast.Forecastday {
		fmt.Printf("  * %s: %s\n", fc.Title, fc.Fcttext)
	}

	fmt.Printf("\n%s\n", current.ObservationTime)
}
