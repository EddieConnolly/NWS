package helpers

import (
	"testing"

	"github.com/EddieConnolly/NWS/pkg/modules"
)

func TestGetTempName(t *testing.T) {

	var cases = []struct {
		name     string
		temp     int
		unit     string
		expected string
	}{
		{"Negative F", -10, "F", "Freezing"},
		{"Positive Freezing F", 5, "F", "Freezing"},
		{"Positive Cold F", 25, "F", "Cold"},
		{"Positive Cold K", 269, "K", "Cold"},
		{"Positive Cold C", -4, "C", "Cold"}, // You think this would be considered freezing right?
		{"Positive Cold Undefined (Default to F)", 25, "M", "Cold"},
		{"Positive Boiling F", 105, "F", "Boiling"},
		{"Out of Range Boiling F", 9999999, "F", "Boiling"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := getTempName(tc.temp, tc.unit)
			if res != tc.expected {
				t.Errorf("failed test %s\ngot: %s\nexp: %s", tc.name, res, tc.expected)
			}
		})
	}
}

func TestSanitizeNWSData(t *testing.T) {

	var cases = []struct {
		name                string
		data                modules.NWSForecast
		expected            modules.ShortForecast
		expectedErrorString string
	}{
		{
			"No points",
			modules.NWSForecast{
				Properties: modules.NWSForecastProperties{
					Periods: []modules.NWSForecastPeriods{},
				},
			},
			modules.ShortForecast{
				Times: []modules.ShortForecastTime{},
			},
			"NWS period data too short to forecast",
		},
		{
			"Not enough points",
			modules.NWSForecast{
				Properties: modules.NWSForecastProperties{
					Periods: []modules.NWSForecastPeriods{
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Later"},
					},
				},
			},
			modules.ShortForecast{
				Times: []modules.ShortForecastTime{},
			},
			"NWS period data too short to forecast",
		},
		{
			"Passing",
			modules.NWSForecast{
				Properties: modules.NWSForecastProperties{
					Periods: []modules.NWSForecastPeriods{
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
						{ShortForecast: "test", Temperature: 1, TemperatureUnit: "F", StartTime: "Now"},
					},
				},
			},
			modules.ShortForecast{
				Times: []modules.ShortForecastTime{
					{LocalTime: "Now", ShortForecast: "test", TemperatureCharacterization: "Freezing"},
					{LocalTime: "Now", ShortForecast: "test", TemperatureCharacterization: "Freezing"},
					{LocalTime: "Now", ShortForecast: "test", TemperatureCharacterization: "Freezing"},
					{LocalTime: "Now", ShortForecast: "test", TemperatureCharacterization: "Freezing"},
					{LocalTime: "Now", ShortForecast: "test", TemperatureCharacterization: "Freezing"},
					{LocalTime: "Now", ShortForecast: "test", TemperatureCharacterization: "Freezing"},
					{LocalTime: "Now", ShortForecast: "test", TemperatureCharacterization: "Freezing"},
					{LocalTime: "Now", ShortForecast: "test", TemperatureCharacterization: "Freezing"},
				},
			},
			"NWS period data too short to forecast",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := SanitizeNWSData(tc.data)
			if err != nil && tc.expectedErrorString != err.Error() {
				t.Errorf("failed test %s\ngot error: %s\nexp error: %s", tc.name, err.Error(), tc.expectedErrorString)

			} else {
				if len(res.Times) != len(tc.expected.Times) {
					t.Errorf("failed test %s - length mismatch\ngot length: %d\nexp length: %d", tc.name, len(res.Times), len(tc.expected.Times))
				}
			}
		})
	}

}
