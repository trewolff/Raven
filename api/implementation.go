package api

import (
	"net/http"
	"raven/services"

	"github.com/labstack/echo/v4"
)

type ServerInterface interface {
	EventIngest(ctx echo.Context) error
}

type RavenAPI struct {
	databaseService services.DatabaseService
	cacheService    services.CacheService
}

func NewRavenAPI(
	databaseService services.DatabaseService,
	cacheService services.CacheService,
) ServerInterface {
	return &RavenAPI{
		databaseService: databaseService,
		cacheService:    cacheService,
	}
}

func (api *RavenAPI) EventIngest(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{"data": "yep"})
}
