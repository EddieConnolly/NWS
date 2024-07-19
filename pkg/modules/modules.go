package modules

// Response from https://api.weather.gov/points/{lat},{long}
type NWSPoints struct {
	Properties NWSPointsProperties `json:"properties"`
}
type NWSPointsProperties struct {
	ForecastURL string `json:"forecastHourly"`
}

// Response from NWSPoints ForecastURL
type NWSForecast struct {
	Properties NWSForecastProperties `json:"properties"`
}
type NWSForecastProperties struct {
	Periods []NWSForecastPeriods `json:"periods"`
}
type NWSForecastPeriods struct {
	ShortForecast   string `json:"shortForecast"`
	Temperature     int    `json:"temperature"`
	TemperatureUnit string `json:"temperatureUnit"`
	StartTime       string `json:"startTime"`
}

// Response for /weather&lat={},long={}
type ShortForecast struct {
	Times []ShortForecastTime `json:"times,omitempty"`
	Error string              `json:"error,omitempty"`
}
type ShortForecastTime struct {
	LocalTime                   string `json:"localTime"`
	ShortForecast               string `json:"shortForecast"`
	TemperatureCharacterization string `json:"temperatureCharacterization"`
}
