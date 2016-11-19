package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marioharper/commuter/directions"
)

type config struct {
	Locations []directions.Location
}

type ConfigManager struct {
	File string
	config
}

func New(file string) ConfigManager {
	return ConfigManager{File: file, config: getConfig(file)}
}

func (cm *ConfigManager) DeleteLocation(locationName string) {

	locationIdx := getLocationIdxByName(cm.config.Locations, locationName)

	cm.config.Locations[locationIdx] = cm.config.Locations[len(cm.config.Locations)-1]
	cm.config.Locations[len(cm.config.Locations)-1] = directions.Location{}
	cm.config.Locations = cm.config.Locations[:len(cm.config.Locations)-1]

	cm.saveConfig()
}

func (cm *ConfigManager) AddLocation(location directions.Location) {

	if i := getLocationIdxByName(cm.config.Locations, location.Name); i >= 0 {
		fmt.Println("You already have a location with the same name")
		os.Exit(-1)
		return
	}

	cm.config.Locations = append(cm.config.Locations, location)

	cm.saveConfig()
}

func (cm *ConfigManager) GetLocations() []directions.Location {
	return cm.config.Locations
}

func (cm *ConfigManager) GetLocationByName(name string) (directions.Location, error) {

	for _, location := range cm.config.Locations {
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

func getConfig(configFile string) config {
	var theConfig config
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
	configJSON, err := json.MarshalIndent(cm.config, "", "  ")
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
