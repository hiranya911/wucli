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
	fmt.Println("\twucli [location/keyword] [feature] [output]\n")

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

	fmt.Println("\t  * alerts")
	fmt.Println("\t  * almanac ")
	fmt.Println("\t  * conditions ")
	fmt.Println("\t  * forecast ")
	fmt.Println("\t  * hourly ")
	fmt.Println("\t  * satellite ")

	fmt.Println("\nWucli can output data formatted in the following types\n")
	fmt.Println("\t  * summary (default)")
	fmt.Println("\t  * raw (key/value pairs)")

	fmt.Println("CREDITS")
	fmt.Println("\tDeveloped by: Hiranya Jayathilaka and Jason Clark (mithereal@gmail.com)")
	fmt.Println("\tRemote API: http://www.wunderground.com\n")
}

func main() {
	argsWithoutProg := os.Args[1:]

	var location string
	var feature string
	var output string

	if len(os.Args) < 2 {
		location = "autoip"
		feature = "conditions,forecast"
		output = "summary"

	} else if argsWithoutProg[0] == "-h" || argsWithoutProg[0] == "help" {
		printUsage()
		return
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
		location = "autoip"
		feature = "conditions,forecast"
		output = "summary"
	}

	wd, err := wu.Query(location, feature)
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
	alerts := wd.Alerts
	almanac := wd.Almanac
	moonphase := wd.MoonPhase
	hourly := wd.HourlyForecast
	satellite := wd.Satellite
	//history := wd.History
	//tide := wd.Tide
	//webcams := wd.Webcams

	switch output {
	case "raw":
		if strings.Contains(feature, "conditions") {
			fmt.Println("Summary:", current.Weather)
			fmt.Println("Temperature:", current.TemperatureString)
			fmt.Println("Feels like:", current.FeelslikeString)
			fmt.Println("Wind:", current.WindString)
			fmt.Println("Precipitation:", current.PrecipTodayString)
		} else {
			fmt.Println("Raw output only available with the conditions feature.")
		}

	default:
		title := fmt.Sprintf("Location: %s (long: %s, lat: %s)", loc.Full, loc.Longitude, loc.Latitude)

		summary := fmt.Sprintf("Summary: %s", current.Weather)
		temperature := fmt.Sprintf("Temperature: %s", current.TemperatureString)
		feelslike := fmt.Sprintf("Feels like: %s", current.FeelslikeString)
		wind := fmt.Sprintf("Wind: %s", current.WindString)
		windchill := fmt.Sprintf("Wind chill: %s", current.WindchillString)
		precipitation := fmt.Sprintf("Precipitation: %s", current.PrecipTodayString)

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
				fmt.Printf("   %s: %s expires on: %s\n", a.Date, a.Description, a.Expires)
			}
		}
		if strings.Contains(feature, "almanac") {
			fmt.Println("\nWeather almanac:")

			fmt.Printf("   Airport Code: %s\n", almanac.AirportCode)
			fmt.Printf("   High Temp: %s\n", almanac.TempHigh.Normal.F)
			fmt.Printf("   Record High Temp: %s\n", almanac.TempHigh.Record.F)
			fmt.Printf("   Date: %s\n", almanac.TempHigh.Recordyear)

			fmt.Printf("   Low Temp: %s\n", almanac.TempLow.Normal.F)
			fmt.Printf("   Record Low Temp: %s\n", almanac.TempLow.Record.F)
			fmt.Printf("   Date: %s\n", almanac.TempLow.Recordyear)
			fmt.Println()
		}

		if strings.Contains(feature, "forecast") {
			fmt.Println("\nWeather Forecast:")
			for _, fc := range forecast.Forecastday {
				fmt.Printf("  * %s: %s\n", fc.Title, fc.Fcttext)
			}
		}

		if strings.Contains(feature, "hourly") {
			fmt.Println("\nHourly Forecast:")
			for _, h := range hourly {
				fmt.Printf("  * %s: %s Temp:%s Dew:%s Uv:%s\n", h.FCTTIME.Pretty, h.Icon, h.Temp.English, h.Humidity, h.Uvi)
			}

			if strings.Contains(feature, "moonphase") {
				fmt.Println("\nMoonphase:")
				fmt.Printf("   Age: %s\n", moonphase.AgeOfMoon)
				fmt.Printf("   Time: %s:%s\n", moonphase.CurrentTime.Hour, moonphase.CurrentTime.Minute)
				fmt.Printf("   Percent Illuminated: %s\n", moonphase.PercentIlluminated)
				fmt.Printf("   Sunrise: %s:%s\n", moonphase.Sunrise.Hour, moonphase.Sunrise.Minute)
				fmt.Printf("   Sunset: %s:%s\n", moonphase.Sunset.Hour, moonphase.Sunset.Minute)
				fmt.Println()

			}
			if strings.Contains(feature, "satellite") {
				fmt.Println("\nMoonphase:")
				fmt.Printf("   Imageurl: %s\n", satellite.ImageURL)
				fmt.Printf("   Imageurl r4: %s:%s\n", satellite.ImageURLIr4)
				fmt.Printf("   Imageurl VI: %s:%s\n", satellite.ImageURLVis)
				fmt.Println()
			}
		}
	}
	if (current.ObservationTime != "") {
		fmt.Printf("\n%s\n", current.ObservationTime)
	}
}
