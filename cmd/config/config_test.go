package config_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/marioharper/commuter/cmd/config"
	"github.com/marioharper/commuter/directions"
	"github.com/stretchr/testify/assert"
)

var configFile = "_fixtures/temp_config.json"

func setup(configMock config.Config) {
	// create test config file
	f, err := os.Create(configFile)
	if err != nil {
		fmt.Printf("creating config file: %s \n", err.Error())
		os.Exit(-1)
	}

	// convert config to json
	configJSON, err := json.MarshalIndent(configMock, "", "  ")
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
	f.Close()
}

func TestConfig_New(t *testing.T) {
	var locations = []directions.Location{
		directions.Location{
			Name:    "work",
			Address: "1600 Amphitheatre Pkwy, Mountain View, CA 94043",
		},
		directions.Location{
			Name:    "home",
			Address: "1060 North Rengstorff Avenue, Mountain View, CA 94043",
		},
	}
	var configMock = config.Config{
		Locations: locations,
	}
	setup(configMock)

	theConfig := config.New(configFile)
	assert.Equal(t, theConfig.File, configFile)
}

func TestConfig_GetLocations(t *testing.T) {
	var locations = []directions.Location{
		directions.Location{
			Name:    "work",
			Address: "1600 Amphitheatre Pkwy, Mountain View, CA 94043",
		},
		directions.Location{
			Name:    "home",
			Address: "1060 North Rengstorff Avenue, Mountain View, CA 94043",
		},
	}
	var configMock = config.Config{
		Locations: locations,
	}
	setup(configMock)

	theConfig := config.New(configFile)
	assert.Equal(t, theConfig.GetLocations(), locations)
}

func TestConfig_GetLocationByName(t *testing.T) {
	var locations = []directions.Location{
		directions.Location{
			Name:    "work",
			Address: "1600 Amphitheatre Pkwy, Mountain View, CA 94043",
		},
		directions.Location{
			Name:    "home",
			Address: "1060 North Rengstorff Avenue, Mountain View, CA 94043",
		},
	}
	var configMock = config.Config{
		Locations: locations,
	}
	setup(configMock)

	theConfig := config.New(configFile)
	retrievedLocation, err := theConfig.GetLocationByName("work")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, retrievedLocation, locations[0])
}

func TestConfig_DeleteLocation(t *testing.T) {
	var locations = []directions.Location{
		directions.Location{
			Name:    "work",
			Address: "1600 Amphitheatre Pkwy, Mountain View, CA 94043",
		},
		directions.Location{
			Name:    "home",
			Address: "1060 North Rengstorff Avenue, Mountain View, CA 94043",
		},
	}
	var configMock = config.Config{
		Locations: locations,
	}
	setup(configMock)

	theConfig := config.New(configFile)
	theConfig.DeleteLocation("work")
	expectedLocations := []directions.Location{
		directions.Location{
			Name:    "home",
			Address: "1060 North Rengstorff Avenue, Mountain View, CA 94043",
		},
	}
	assert.Equal(t, theConfig.GetLocations(), expectedLocations)
}

func TestConfig_AddLocation(t *testing.T) {
	var locations = []directions.Location{}
	var configMock = config.Config{
		Locations: locations,
	}
	setup(configMock)

	theConfig := config.New(configFile)
	newLocation := directions.Location{
		Name:    "new",
		Address: "new location",
	}

	theConfig.AddLocation(newLocation)
	assert.Equal(t, theConfig.Config.Locations[0], newLocation)
}
