package weather

type WeatherInfo struct {
	Summary           string
	PrecipProbability float64
	Temp              float64
	Wind              float64
}

type WeatherInfoer interface {
	GetWeatherInfo(t int, latitude float64, longitude float64) (*WeatherInfo, error)
}
