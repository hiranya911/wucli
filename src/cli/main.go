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
	
	fmt.Println("\tOptionally you can run wucli with a feature argument to get a certain data output related to a specific")
	fmt.Println("\tcity or area. Following types of feature arguments are supported as of now:\n")

	fmt.Println("\t  * Data outputs Supported are alerts almanac conditions  forecast  history hourly  satellite tide webcams yesterday ")
	
	fmt.Println("CREDITS")
	fmt.Println("\tDeveloped by: Hiranya Jayathilaka and Jason Clark (mithereal@gmail.com)")
	fmt.Println("\tRemote API: http://www.wunderground.com\n")
}


func main() {
	argsWithoutProg := os.Args[1:]
	
	var location string
	var feature string
	var output string
	//var filters string
	
	if len(argsWithoutProg) == 0 {
		location = "autoip"
		feature = "conditions,forecast"
		output = "summary"
		
	} else if len(argsWithoutProg) == 1 {
		location = argsWithoutProg[0]
		feature = "conditions,forecast"
		output = "summary"
	
	} else if len(argsWithoutProg) == 2 {
		location = argsWithoutProg[0]
		feature = argsWithoutProg[1]
		output = "summary"
	
	} else if len(argsWithoutProg) == 3 {
		location = argsWithoutProg[0]
		feature = argsWithoutProg[1]
		output = argsWithoutProg[2]
		
	} else if len(argsWithoutProg) == 4 {
		location = argsWithoutProg[0]
		feature = argsWithoutProg[1]
		output = argsWithoutProg[2]
		
	} else {
		printUsage()
		return
	}

	if argsWithoutProg[0] == "-h" || argsWithoutProg[0] == "help" {
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
	alerts := wd.Alerts
	almanac := wd.Almanac
	//moonphase := wd.MoonPhase
	hourly := wd.HourlyForecast
	//satellite := wd.Satellite
	//history := wd.History
	//tide := wd.Tide
	//webcams := wd.Webcams
	
	results := wd.Response.Results
	
	
	switch output {
    case "raw": 
    if strings.Contains(feature, "conditions") {
		fmt.Println("Summary:", current.Weather)
		fmt.Println("Temperature:",current.TemperatureString)
		fmt.Println("Feels like:",current.FeelslikeString)
		fmt.Println("Wind:",current.WindString)
		fmt.Println("Precipitation:",current.PrecipTodayString)
	}
		
    default: 
    title := fmt.Sprintf("Location: %s (long: %s, lat: %s)", loc.Full, loc.Longitude, loc.Latitude)

	summary := fmt.Sprintf("Summary: %s\n", current.Weather)
	temperature := fmt.Sprintf("Temperature: %s\n", current.TemperatureString)
	feelslike := fmt.Sprintf("Feels like: %s\n", current.FeelslikeString)
	wind := fmt.Sprintf("Wind: %s\n", current.WindString)
	windchill := fmt.Sprintf("Wind chill: %s\n", current.WindchillString)
	precipitation := fmt.Sprintf("Precipitation: %s\n", current.PrecipTodayString)
	
		if strings.Contains(feature, "conditions") {
				fmt.Println(title)
	fmt.Println(strings.Repeat("=", len(title)))
	fmt.Println(summary)
	fmt.Println(temperature)
	fmt.Println(feelslike)
	fmt.Println(wind)
	fmt.Println(windchill)
	fmt.Println(precipitation)
}
		if strings.Contains(feature, "alerts") {
				fmt.Println("\nWeather Alerts:")
				for _, a := range alerts {
		fmt.Printf("  * %s: %s expires on: %s\n", a.Date, a.Description, a.Expires)
	}
}
		if strings.Contains(feature, "almanac") {
				fmt.Println("\nWeather almanac:")
				
		fmt.Printf("  * Airport Code: %s\n", almanac.AirportCode)
		fmt.Printf("  * High Temp: %s\n", almanac.TempHigh.Normal.F)
		fmt.Printf("  * Record High Temp: %s\n", almanac.TempHigh.Record.F)
		fmt.Printf("  * Date: %s\n", almanac.TempHigh.Recordyear)

		fmt.Printf("  * Low Temp: %s\n", almanac.TempLow.Normal.F)
		fmt.Printf("  * Record Low Temp: %s\n", almanac.TempLow.Record.F)
		fmt.Printf("  * Date: %s\n", almanac.TempLow.Recordyear)
	
}
	
	
		if strings.Contains(feature, "forecast") {
	fmt.Println("\nWeather Forecast:")
	for _, fc := range forecast.Forecastday {
		fmt.Printf("  * %s: %s\n", fc.Title, fc.Fcttext)
	}
}
	
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
	
	if strings.Contains(feature, "hourly") {
		fmt.Println("\nHourly Forecast:")
		for _, h := range hourly {
		fmt.Printf("  * %s: %s Temp:%s Dew:%s Uv:%s\n", h.FCTTIME.Pretty, h.Icon, h.Temp.English,h.Humidity,h.Uvi)
	}
	
    
}

	fmt.Printf("\n%s\n", current.ObservationTime)
    }
	
		

	
}
