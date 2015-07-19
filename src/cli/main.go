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
	
	fmt.Println("\tOptionally you can run wucli with a feature argument to get a certain data format related to a specific")
	fmt.Println("\tcity or area. Following types of feature arguments are supported as of now:\n")

	fmt.Println("\t  * Data Formats Supported are alerts almanac conditions  forecast  history hourly  satellite tide webcams yesterday ")
	
	fmt.Println("CREDITS")
	fmt.Println("\tDeveloped by: Hiranya Jayathilaka and Jason Clark (mithereal@gmail.com)")
	fmt.Println("\tRemote API: http://www.wunderground.com\n")
}


func main() {
	argsWithoutProg := os.Args[1:]
	
	var location string
	var feature string
	var format string
	//var filters string
	
	if len(argsWithoutProg) == 0 {
		location = "autoip"
		feature = "conditions"
		format = "summary"
		
	} else if len(argsWithoutProg) == 1 {
		location = argsWithoutProg[0]
		feature = "conditions"
		format = "summary"
	
	} else if len(argsWithoutProg) == 2 {
		location = argsWithoutProg[0]
		feature = argsWithoutProg[1]
		format = "summary"
	
	} else if len(argsWithoutProg) == 3 {
		location = argsWithoutProg[0]
		feature = argsWithoutProg[1]
		format = argsWithoutProg[2]
		
	} else if len(argsWithoutProg) == 4 {
		location = argsWithoutProg[0]
		feature = argsWithoutProg[1]
		format = argsWithoutProg[2]
		
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
	
	current := wd.CurrentObservation
	loc := current.DisplayLocation
	forecast := wd.Forecast.TxtForecast
	//alerts := wd.Alerts
	//almanac := wd.Almanac
	//moonphase := wd.MoonPhase
	//hourly := wd.HourlyForecast
	//satellite := wd.Satellite
	//history := wd.History
	//tide := wd.Tide
	//webcams := wd.Webcams
	results := wd.Response.Results
	
	// we need to implement different view and filter functions 
	
		if results != nil {
		fmt.Println("Locations matched:")
		for _, r := range results {
			if r.State != "" {
				fmt.Printf("  * %s, %s, %s (zmw:%s)\n", r.City, r.State, r.CountryName, r.Zmw)
			} else {
				fmt.Printf("  * %s, %s (zmw:%s)\n", r.City, r.CountryName, r.Zmw)
			}
		}
		return
	}
	
	if strings.Contains(format, "raw") {
		fmt.Println(current.Weather)
		fmt.Println(current.TemperatureString)
		fmt.Println(current.FeelslikeString)
		fmt.Println(current.WindString)
		fmt.Println(current.PrecipTodayString)
    }else{
	title := fmt.Sprintf("Location: %s (long: %s, lat: %s)", loc.Full, loc.Longitude, loc.Latitude)
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", len(title)))
	
	summary := fmt.Sprintf("Summary: %s\n", current.Weather)
	temperature := fmt.Sprintf("Temperature: %s\n", current.TemperatureString)
	feelslike := fmt.Sprintf("Feels like: %s\n", current.FeelslikeString)
	wind := fmt.Sprintf("Wind: %s\n", current.WindString)
	windchill := fmt.Sprintf("Wind chill: %s\n", current.WindchillString)
	precipitation := fmt.Sprintf("Precipitation: %s\n", current.PrecipTodayString)
	
	fmt.Println(summary)
	fmt.Println(temperature)
	fmt.Println(feelslike)
	fmt.Println(wind)
	fmt.Println(windchill)
	fmt.Println(precipitation)
	

	if strings.Contains(feature, "forecast") {
	fmt.Println("\nWeather Forecast:")
	for _, fc := range forecast.Forecastday {
		fmt.Printf("  * %s: %s\n", fc.Title, fc.Fcttext)
	}
}
}


	fmt.Printf("\n%s\n", current.ObservationTime)
}
