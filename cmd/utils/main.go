package utils

import (
	"github.com/marioharper/commuter/cmd/config"
	"github.com/marioharper/commuter/directions"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func RemoveLocation() {

}

func AddLocation(configFile string, location directions.Location) {
	configS := config.GetConfig(configFile)
	configS.Locations = append(configS.Locations, location)
	config.SaveConfig(configFile, configS)
}

func GetLocations(configFile string) []directions.Location {
	config := config.GetConfig(configFile)
	return config.Locations
}

func GetLocationByName(locations []directions.Location, name string) int {

	for idx, location := range locations {
		if location.Name == name {
			return idx
		}
	}

	return -1
}
