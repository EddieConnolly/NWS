package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/EddieConnolly/NWS/pkg/modules"
)

const (
	NWSPOINTSURL = "https://api.weather.gov/points/"
)

// Returns the forecast URL for the selected latitude and longitude
func GetHourlyForecastURL(lat string, long string) (string, error) {
	url := fmt.Sprintf("%s%s,%s", NWSPOINTSURL, lat, long)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid response from URL<%s> Status<%d>", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var pointsdata modules.NWSPoints

	err = json.Unmarshal(body, &pointsdata)
	if err != nil {
		return "", err
	}

	return pointsdata.Properties.ForecastURL, nil
}

// Returns the hourly weather data from the supplied URL
func GetHourlyForecastData(url string) (modules.NWSForecast, error) {
	var forecastData modules.NWSForecast

	resp, err := http.Get(url)
	if err != nil {
		return forecastData, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return forecastData, err
	}

	err = json.Unmarshal(body, &forecastData)

	return forecastData, err
}
