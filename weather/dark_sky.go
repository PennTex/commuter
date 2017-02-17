package weather

import (
	"errors"
	"strconv"
	"time"

	forecast "github.com/mlbright/forecast/v2"
)

// supplied via -ldflags in Makefile
var DARK_SKY_API_KEY = ""

type DarkSkyWeatherInfoer struct{}

func (d *DarkSkyWeatherInfoer) GetWeatherInfo(t int, latitude float64, longitude float64) (*WeatherInfo, error) {
	var info WeatherInfo
	tHr, _, _ := time.Unix(int64(t), 0).Clock()

	lat := strconv.FormatFloat(latitude, 'f', 6, 64)
	lng := strconv.FormatFloat(longitude, 'f', 6, 64)

	f, err := forecast.Get(DARK_SKY_API_KEY, lat, lng, strconv.Itoa(t), forecast.US, forecast.English)
	if err != nil {
		return nil, errors.New("Getting forecast from DarkSky")
	}

	for _, value := range f.Hourly.Data {
		hr, _, _ := time.Unix(int64(value.Time), 0).Clock()

		if hr == tHr {
			info.Summary = value.Summary
			info.Temp = value.Temperature
			info.PrecipProbability = (value.PrecipProbability * 100)
			info.Wind = value.WindSpeed
			break
		}
	}

	return &info, nil
}
