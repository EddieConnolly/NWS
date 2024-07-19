package main

import (
	"github.com/EddieConnolly/NWS/pkg/endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/weather", endpoints.Weather)

	router.Run(":8080")
}
