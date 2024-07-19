package endpoints

import (
	"net/http"

	"github.com/EddieConnolly/NWS/pkg/helpers"
	"github.com/EddieConnolly/NWS/pkg/modules"
	"github.com/gin-gonic/gin"
)

func Weather(c *gin.Context) {
	var response modules.ShortForecast

	lat := c.Query("lat")
	long := c.Query("long")

	if lat == "" || long == "" {
		response.Error = `Usage - /weather?lat={latitude}&long={longitude}`
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get the URL
	url, err := helpers.GetHourlyForecastURL(lat, long)
	if err != nil {
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Get the period data
	data, err := helpers.GetHourlyForecastData(url)
	if err != nil {
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Sanatize the period data
	response, err = helpers.SanitizeNWSData(data)
	if err != nil {
		response.Error = err.Error()
		response.Times = []modules.ShortForecastTime{}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
