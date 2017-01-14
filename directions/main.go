package directions

import (
	"errors"
	"fmt"
	"strings"
)

type Location struct {
	Name    string
	Address string
}

type CommuteInfo struct {
	TotalDistance int
	TotalDuration float64
	Lat           float64
	Lng           float64
}

type Commute struct {
	From Location
	To   Location
	Time int64
	*CommuteInfo
}

type AddressValidator interface {
	IsValidAddress(address string) bool
}

type CommuteInfoer interface {
	GetCommuteInfo(from Location, to Location, time int64) (*CommuteInfo, error)
}

func NewCommute(infoer CommuteInfoer, from Location, to Location, time int64) (*Commute, error) {
	info, err := infoer.GetCommuteInfo(from, to, time)

	if err != nil {
		return nil, errors.New("Getting directions")
	}

	return &Commute{
		From:        from,
		To:          to,
		Time:        time,
		CommuteInfo: info,
	}, nil
}

func (c *Commute) GetMapsURL() string {
	from := strings.Replace(c.From.Address, " ", "+", -1)
	to := strings.Replace(c.To.Address, " ", "+", -1)
	url := fmt.Sprintf("https://www.google.com/maps/dir/%s/%s", from, to)

	return url
}
