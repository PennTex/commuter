package directions

import (
	"log"
	"os"

	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func Init() {

	// get maps api key from environment
	mapsAPIKey := os.Getenv("MAPS_API_KEY")

	c, err := maps.NewClient(maps.WithAPIKey(mapsAPIKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:        "1835 Kramer Ln Austin TX",
		Destination:   "1500 Royal Crest Drive Austin TX",
		DepartureTime: "now",
	}
	resp, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	totalDistance := 0
	totalDuration := 0.00
	legs := resp[0].Legs

	for _, leg := range legs {
		totalDistance += leg.Distance.Meters
		totalDuration += leg.DurationInTraffic.Minutes()
	}

	pretty.Println(totalDuration)
}
