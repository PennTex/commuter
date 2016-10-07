package utils

import (
	"fmt"
	"os"

	"github.com/marioharper/commuter/cmd/config"
	"github.com/marioharper/commuter/directions"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func DeleteLocation(configFile string, locationName string) {
	configS := config.GetConfig(configFile)

	locationIdx := GetLocationByName(configS.Locations, locationName)

	configS.Locations[locationIdx] = configS.Locations[len(configS.Locations)-1]
	configS.Locations[len(configS.Locations)-1] = directions.Location{}
	configS.Locations = configS.Locations[:len(configS.Locations)-1]

	config.SaveConfig(configFile, configS)
}

func AddLocation(configFile string, location directions.Location) {
	configS := config.GetConfig(configFile)

	if i := GetLocationByName(configS.Locations, location.Name); i >= 0 {
		fmt.Println("You already have a location with the same name")
		os.Exit(-1)
		return
	}

	fmt.Println(GetLocationByName(configS.Locations, location.Name))

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

type Logger struct {
	Logging bool
}

func (l *Logger) Log(output string) {
	if l.Logging {
		fmt.Println(output)
	}
}
