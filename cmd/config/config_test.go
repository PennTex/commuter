package config_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/PennTex/commuter/cmd/config"
	"github.com/PennTex/commuter/cmd/utils"
	"github.com/PennTex/commuter/directions"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, theConfig.File, configFile)
	assert.Equal(t, theConfig.Config.Locations, locations)
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
	utils.ProcessError(err, "")

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
	assert.Equal(t, theConfig.GetLocations(), locations)
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
	assert.Equal(t, len(theConfig.Config.Locations), 1)
}
