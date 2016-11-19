package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marioharper/commuter/directions"
)

type Config struct {
	Locations []directions.Location
}

type ConfigManager struct {
	File string
	Config
}

func New(file string) ConfigManager {
	return ConfigManager{File: file, Config: getConfig(file)}
}

func (cm *ConfigManager) DeleteLocation(locationName string) {

	locationIdx := getLocationIdxByName(cm.Config.Locations, locationName)

	cm.Config.Locations[locationIdx] = cm.Config.Locations[len(cm.Config.Locations)-1]
	cm.Config.Locations[len(cm.Config.Locations)-1] = directions.Location{}
	cm.Config.Locations = cm.Config.Locations[:len(cm.Config.Locations)-1]

	cm.saveConfig()
}

func (cm *ConfigManager) AddLocation(location directions.Location) {

	if i := getLocationIdxByName(cm.Config.Locations, location.Name); i >= 0 {
		fmt.Println("You already have a location with the same name")
		os.Exit(-1)
		return
	}

	cm.Config.Locations = append(cm.Config.Locations, location)

	cm.saveConfig()
}

func (cm *ConfigManager) GetLocations() []directions.Location {
	return cm.Config.Locations
}

func (cm *ConfigManager) GetLocationByName(name string) (directions.Location, error) {

	for _, location := range cm.Config.Locations {
		if location.Name == name {
			return location, nil
		}
	}

	return directions.Location{}, fmt.Errorf("Location %s not found.", name)
}

func getLocationIdxByName(locations []directions.Location, name string) int {
	for idx, location := range locations {
		if location.Name == name {
			return idx
		}
	}

	return -1
}

func getConfig(configFile string) Config {
	var theConfig Config
	var f *os.File
	var err error

	// create if not exists
	if f, err = os.Open(configFile); err != nil {
		f, err = os.Create(configFile)
		if err != nil {
			fmt.Printf("creating config file: %s \n", err.Error())
			os.Exit(-1)
		}
	}

	jsonParser := json.NewDecoder(f)

	if fi, err := f.Stat(); err != nil {
		fmt.Printf("getting file stat: %s \n", err.Error())
	} else {
		if fi.Size() == 0 {
			return theConfig
		}
	}

	if err := jsonParser.Decode(&theConfig); err != nil {
		fmt.Printf("parsing config file: %s \n", err.Error())
		os.Exit(-1)
	}

	return theConfig
}

func (cm *ConfigManager) saveConfig() {
	// overrite current config file
	f, err := os.Create(cm.File)
	if err != nil {
		fmt.Printf("creating config file: %s \n", err.Error())
		os.Exit(-1)
	}

	// convert config to json
	configJSON, err := json.MarshalIndent(cm.Config, "", "  ")
	if err != nil {
		fmt.Printf("marshalling config file: %s \n", err.Error())
		os.Exit(-1)
	}

	// write json to config file
	_, err = f.WriteString(string(configJSON))
	if err != nil {
		fmt.Printf(err.Error())
	}

	f.Sync()
}
