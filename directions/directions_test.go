package directions_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/PennTex/commuter/cmd/config"
	"github.com/PennTex/commuter/cmd/utils"
	"github.com/PennTex/commuter/directions"
)

type Location struct {
	Name    string
	Address string
}

type Commute struct {
	From          Location
	To            Location
	Time          int64
	TotalDistance int
	TotalDuration float64
	Lat           float64
	Lng           float64
}

var configFile = "../cmd/config/_fixtures/temp_config.json"

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

func TestCreateCommute(t *testing.T) {
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
	var now = time.Now().Unix()

	theConfig := config.New(configFile)

	from, _ := theConfig.GetLocationByName("work")
	to, _ := theConfig.GetLocationByName("home")

	commute := directions.NewCommute(from, to, now)

	fmt.Println(commute)
}
