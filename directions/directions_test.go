package directions_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/PennTex/commuter/directions"
)

type FakeCommuteInfoer struct {
	Info *directions.CommuteInfo
	Err  error
}

func (f FakeCommuteInfoer) GetCommuteInfo(from directions.Location, to directions.Location, time int64) (*directions.CommuteInfo, error) {
	if f.Err != nil {
		return nil, f.Err
	}

	return f.Info, nil
}
func TestCommute_New(t *testing.T) {
	cases := []struct {
		f               FakeCommuteInfoer
		from            *directions.Location
		to              *directions.Location
		time            int64
		expectedCommute *directions.Commute
	}{
		{
			f: FakeCommuteInfoer{
				Info: &directions.CommuteInfo{
					TotalDistance: 0,
					TotalDuration: 0,
					Lat:           0,
					Lng:           0,
				},
				Err: nil,
			},
			from: &directions.Location{},
			to:   &directions.Location{},
			time: 0,
			expectedCommute: &directions.Commute{
				From: directions.Location{},
				To:   directions.Location{},
				Time: 0,
				CommuteInfo: &directions.CommuteInfo{
					TotalDistance: 0,
					TotalDuration: 0,
					Lat:           0,
					Lng:           0,
				},
			},
		},
		{
			f: FakeCommuteInfoer{
				Info: &directions.CommuteInfo{},
				Err:  errors.New("Error getting commute info"),
			},
			from:            &directions.Location{},
			to:              &directions.Location{},
			time:            0,
			expectedCommute: nil,
		},
	}

	for _, c := range cases {
		commute, err := directions.NewCommute(c.f, *c.from, *c.to, c.time)

		if c.f.Err != nil {
			if err == nil {
				t.Errorf("The code did not error")
			}
		} else {
			if !reflect.DeepEqual(commute, c.expectedCommute) {
				t.Errorf("Expected commute to be %q but it was %q", c.expectedCommute, commute)
			}
		}
	}
}

func TestCommute_GetMapsURL(t *testing.T) {
	commute := directions.Commute{}
	commute.From = directions.Location{
		Address: "from-address",
	}
	commute.To = directions.Location{
		Address: "to-address",
	}

	mapsURL := commute.GetMapsURL()
	expectedURL := "https://www.google.com/maps/dir/from-address/to-address"

	if mapsURL != expectedURL {
		t.Errorf("Expected URL to be %s but it was %s", expectedURL, mapsURL)
	}
}
