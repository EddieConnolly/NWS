package helpers

import (
	"errors"

	"github.com/EddieConnolly/NWS/pkg/modules"
)

var (
	// Temp names in F (Below that temp)
	temps = []struct {
		temp int
		name string
	}{{10, "Freezing"}, {30, "Cold"}, {60, "Plesant"}, {80, "Warm"}, {100, "Hot"}, {999, "Boiling"}}
)

// Convert the data from NWS into a short 8 hour forecast
func SanitizeNWSData(data modules.NWSForecast) (modules.ShortForecast, error) {
	var sanitized modules.ShortForecast

	if len(data.Properties.Periods) < 8 {
		err := errors.New("NWS period data too short to forecast")
		return sanitized, err
	}

	for _, v := range data.Properties.Periods[:8] {
		var t modules.ShortForecastTime
		t.TemperatureCharacterization = getTempName(v.Temperature, v.TemperatureUnit)
		t.LocalTime = v.StartTime
		t.ShortForecast = v.ShortForecast

		sanitized.Times = append(sanitized.Times, t)
	}

	return sanitized, nil
}

// Get a cool name for the current temp, assuming a unit of K,C or F
func getTempName(temp int, unit string) string {

	if unit == "K" {
		temp = ((temp*9)/5 - 460) // Approx Kelvin to F
	}

	if unit == "C" {
		temp = ((temp*9)/5 + 32) // Approx Kelvin to F
	}


	for _, v := range temps[:len(temps)-1] {
		if temp <= v.temp {
			return v.name
		}
	}

	return temps[len(temps)-1].name
}
