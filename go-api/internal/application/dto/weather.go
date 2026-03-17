package dto

import "time"

// WeatherForecast represents a weather forecast entry.
type WeatherForecast struct {
	Date         time.Time `json:"date"`
	TemperatureC int       `json:"temperatureC"`
	TemperatureF int       `json:"temperatureF"`
	Summary      string    `json:"summary"`
}

// NewWeatherForecast creates a new WeatherForecast, computing TemperatureF from TemperatureC.
func NewWeatherForecast(date time.Time, temperatureC int, summary string) WeatherForecast {
	return WeatherForecast{
		Date:         date,
		TemperatureC: temperatureC,
		TemperatureF: 32 + int(float64(temperatureC)/0.5556),
		Summary:      summary,
	}
}
