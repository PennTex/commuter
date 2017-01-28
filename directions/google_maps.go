package directions

import (
	"context"
	"strconv"

	"googlemaps.github.io/maps"
)

// supplied via -ldflags in Makefile
var MAPS_API_KEY = ""

type GoogleMapsAddressValidator struct{}

func (g GoogleMapsAddressValidator) IsValidAddress(address string) (bool, error) {
	if address == "" {
		return false, nil
	}

	client, err := maps.NewClient(maps.WithAPIKey(MAPS_API_KEY))
	if err != nil {
		return false, err
	}

	a := &maps.GeocodingRequest{
		Address: address,
	}
	_, err = client.Geocode(context.Background(), a)
	if err != nil {
		return false, nil
	}
	return true, nil
}

type GoogleMapsCommuteInfoer struct{}

func (g GoogleMapsCommuteInfoer) GetCommuteInfo(from Location, to Location, travelTime int64) (*CommuteInfo, error) {
	client, err := maps.NewClient(maps.WithAPIKey(MAPS_API_KEY))
	if err != nil {
		return nil, err
	}

	r := &maps.DirectionsRequest{
		Origin:        from.Address,
		Destination:   to.Address,
		DepartureTime: strconv.FormatInt(travelTime, 10),
	}

	resp, _, err := client.Directions(context.Background(), r)
	if err != nil {
		return nil, err
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

	return &CommuteInfo{
			TotalDistance: totalDistance,
			TotalDuration: totalDuration,
			Lat:           lat,
			Lng:           lng,
		},
		nil
}
