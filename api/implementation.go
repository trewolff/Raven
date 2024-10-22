package api

import (
	"net/http"
	"raven/logging"
	"raven/services"

	"github.com/labstack/echo/v4"
)

func HandleError(err error, statusCode int, message string) error {
	logging.Logger.Error(err.Error())
	return echo.NewHTTPError(statusCode, message)
}

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
	var body map[string]interface{}
	err := ctx.Bind(&body)
	if err != nil {
		HandleError(err, http.StatusBadRequest, "Invalid body in request")
	}
	eventId, err := api.databaseService.WriteEvent(body["event_id"].(string))
	if err != nil {
		HandleError(err, http.StatusInternalServerError, "Could not write event")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"event_id": eventId})
}
