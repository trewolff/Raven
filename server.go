package main

import (
	"flag"
	"fmt"
	"raven/api"
	"raven/logging"
	"raven/services"

	"github.com/labstack/echo/v4"
)

func main() {
	logging.InitLogger()
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	echoRouter := echo.New()
	databaseService := services.NewDatabaseService(services.NewDatabaseConnection())
	cacheService := services.NewCacheService(services.NewCacheConnection())
	ravenApi := api.NewRavenAPI(
		databaseService,
		cacheService,
	)
	echoRouter.POST("/api/v1/login", ravenApi.Login)
	echoRouter.POST("/api/v1/events", ravenApi.EventIngest)
	echoRouter.GET("/api/v1/events/:id", ravenApi.EventGet)
	echoRouter.Start(fmt.Sprintf("0.0.0.0:%d", *port))
}
