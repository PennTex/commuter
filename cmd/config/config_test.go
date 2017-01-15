package config_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/PennTex/commuter/cmd/config"
	"github.com/PennTex/commuter/cmd/utils"
	"github.com/PennTex/commuter/directions"
)

var configFile = "_fixtures/temp_config.json"

func setup(configMock config.Config) {
	// create test config file
	f, err := os.Create(configFile)
	utils.ProcessError(err, "Creating config file")

	// convert config to json
	configJSON, err := json.MarshalIndent(configMock, "", "  ")
	utils.ProcessError(err, "Marshalling config file")

	// write json to config file
	_, err = f.WriteString(string(configJSON))
	utils.ProcessError(err, "Writing to JSON file")

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

	if !reflect.DeepEqual(theConfig.File, configFile) {
		t.Errorf("Expected config file to be %q but it was %q", configFile, theConfig.File)
	}

	if !reflect.DeepEqual(theConfig.Config.Locations, locations) {
		t.Errorf("Expected config locations to be %q but they were %q", locations, theConfig.Config.Locations)
	}
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

	if !reflect.DeepEqual(theConfig.GetLocations(), locations) {
		t.Errorf("Expected locations to be %q but they were %q", locations, theConfig.GetLocations())
	}
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
	utils.ProcessError(err, "")

	if !reflect.DeepEqual(retrievedLocation, locations[0]) {
		t.Errorf("Expected location to be %q but it was %q", locations[0], retrievedLocation)
	}
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
	if !reflect.DeepEqual(theConfig.GetLocations(), expectedLocations) {
		t.Errorf("Expected locations to be %q but they were %q", expectedLocations, theConfig.GetLocations())
	}
}

func TestConfig_DeleteLocation_InvalidLocation(t *testing.T) {
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
	theConfig.DeleteLocation("i'm not real")

	if !reflect.DeepEqual(theConfig.GetLocations(), locations) {
		t.Errorf("Expected locations to be %q but they were %q", locations, theConfig.GetLocations())
	}
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
	if !reflect.DeepEqual(theConfig.Config.Locations[0], newLocation) {
		t.Errorf("Expected location to be %q but it was %q", newLocation, theConfig.Config.Locations[0])
	}
	if len(theConfig.Config.Locations) != 1 {
		t.Errorf("Expected location lenght to be 1 but it was %d", len(theConfig.Config.Locations))
	}
}
