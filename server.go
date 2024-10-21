package main

import (
	"flag"
	"fmt"
	"raven/api"
	"raven/services"

	"github.com/labstack/echo/v4"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	echoRouter := echo.New()
	databaseService := services.NewDatabaseService(services.NewDatabaseConnection())
	cacheService := services.NewCacheService(services.NewCacheConnection())
	ravenApi := api.NewRavenAPI(
		databaseService,
		cacheService,
	)
	echoRouter.POST("/ap1/v1/events", ravenApi.EventIngest)
	echoRouter.Start(fmt.Sprintf("0.0.0.0:%d", *port))
}
