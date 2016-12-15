package directions

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

// supplied via -ldflags in Makefile
var MAPS_API_KEY = ""

type Location struct {
	Name    string
	Address string
}

type Commute struct {
	From Location
	To   Location
	Time int64
}

type CommuteInfo struct {
	TotalDistance int
	TotalDuration float64
	Lat           float64
	Lng           float64
}

func (c *Commute) GetInfo(travelTime int64) CommuteInfo {

	var info CommuteInfo

	client, err := maps.NewClient(maps.WithAPIKey(MAPS_API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:        c.From.Address,
		Destination:   c.To.Address,
		DepartureTime: strconv.FormatInt(travelTime, 10),
	}
	resp, _, err := client.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	totalDistance := 0
	totalDuration := 0.00
	lat := 0.00
	lng := 0.00
	legs := resp[0].Legs

	for _, leg := range legs {
		totalDistance += leg.Distance.Meters
		totalDuration += leg.DurationInTraffic.Minutes()
		lat += leg.StartLocation.Lat
		lng += leg.StartLocation.Lng
	}

	info.TotalDistance = totalDistance
	info.TotalDuration = totalDuration
	info.Lat = lat
	info.Lng = lng

	return info
}

func (c *Commute) GetMapsURL() string {
	from := strings.Replace(c.From.Address, " ", "+", -1)
	to := strings.Replace(c.To.Address, " ", "+", -1)
	url := fmt.Sprintf("https://www.google.com/maps/dir/%s/%s", from, to)

	return url
}
