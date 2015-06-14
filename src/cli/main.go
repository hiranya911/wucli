package main

import (
	"fmt"
	"os"
	"wu"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Println("Usage: wucli <location>")
		return
	}

	location := argsWithoutProg[0]
	wd, err := wu.Query(location)
	if err != nil {
		fmt.Println(err)
		return
	}

	current := wd.CurrentObservation
	fmt.Printf("Temperature: %s\n", current.TemperatureString)
}
