package weather

import (
	"log"
	"strconv"
	"time"

	forecast "github.com/mlbright/forecast/v2"
)

// supplied via -ldflags in Makefile
var DARK_SKY_API_KEY = ""

type WeatherInfo struct {
	Summary           string
	PrecipProbability float64
	Temp              float64
	Wind              float64
}

func GetInfo(idealTime int, latitude float64, longitude float64) WeatherInfo {
	var info WeatherInfo
	idealTimeHr, _, _ := time.Unix(int64(idealTime), 0).Clock()

	lat := strconv.FormatFloat(latitude, 'f', 6, 64)
	lng := strconv.FormatFloat(longitude, 'f', 6, 64)

	f, err := forecast.Get(DARK_SKY_API_KEY, lat, lng, strconv.Itoa(idealTime), forecast.US)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range f.Hourly.Data {
		hr, _, _ := time.Unix(int64(value.Time), 0).Clock()

		if err != nil {
			panic(err)
		}

		if hr == idealTimeHr {
			info.Summary = value.Summary
			info.Temp = value.Temperature
			info.PrecipProbability = (value.PrecipProbability * 100)
			info.Wind = value.WindSpeed
			break
		}
	}

	return info
}
