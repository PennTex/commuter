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

func GetConfig(configFile string) Config {
	var config Config
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
			return config
		}
	}

	if err := jsonParser.Decode(&config); err != nil {
		fmt.Printf("parsing config file: %s \n", err.Error())
		os.Exit(-1)
	}

	return config
}

func SaveConfig(configFile string, config Config) {

	// overrite current config file
	f, err := os.Create(configFile)
	if err != nil {
		fmt.Printf("creating config file: %s \n", err.Error())
		os.Exit(-1)
	}

	// convert config to json
	configJSON, err := json.MarshalIndent(config, "", "  ")
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
