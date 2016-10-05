package directions

import (
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var mapsAPIKey = os.Getenv("MAPS_API_KEY")

type Location struct {
	Name    string
	Address string
}

func TotalDistDur(start, end Location) (int, float64) {

	client, err := maps.NewClient(maps.WithAPIKey(mapsAPIKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:        start.Address,
		Destination:   end.Address,
		DepartureTime: strconv.FormatInt(time.Now().Unix(), 10),
	}
	resp, _, err := client.Directions(context.Background(), r)
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

	return totalDistance, totalDuration
}

func Init() {
}
