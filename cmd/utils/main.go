package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marioharper/commuter/directions"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func RemoveLocation() {

}

func AddLocation() {

}

func GetLocations(configFile string) []directions.Location {

	locations := []directions.Location{}

	fmt.Println(configFile)
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Printf("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&locations); err != nil {
		fmt.Printf("parsing config file", err.Error())
	}

	return locations
}

func GetLocationByName(locations []directions.Location, name string) int {

	for idx, location := range locations {
		if location.Name == name {
			return idx
		}
	}

	return -1
}
