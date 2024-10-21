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
	sqlService := services.NewDatabaseConnection()
	cacheService := services.NewCacheConnection()
	ravenApi := api.NewRavenAPI{
		databaseService: services.NewDatabaseService(sqlService),
		cacheService:    services.NewCacheService(cacheService),
	}
	echoRouter.POST("/ap1/v1/events", ravenApi.EventIngest)
	echoRouter.Start(fmt.Sprintf("0.0.0.0:%d", *port))
}
