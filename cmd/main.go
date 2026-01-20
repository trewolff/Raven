package main

import (
	"net/http"
	"raven/internal/api"
	"raven/internal/services/cache"
	"raven/internal/services/database"
	"raven/migrations"
)

func main() {
	db := database.NewDatabaseConnection()
	err := migrations.MigrateDatabase(db)
	if err != nil {
		panic("Failed to run migrations: " + err.Error())
	}

	router := api.NewRouter()
	handlerConfig := api.HandlerConfig{
		DatabaseService: database.NewDatabaseService(db),
		CacheService:    cache.NewCacheService(cache.NewCacheConnection()),
	}
	handlers := api.NewHandlers(handlerConfig)

	router.HandleFunc("PUT /api/v1/cases/new", handlers.CreateCase)
	router.HandleFunc("GET /api/v1/cases", handlers.ListCases)
	router.HandleFunc("GET /api/v1/cases/{id}", handlers.ReadCase)

	router.HandleFunc("PUT /api/v1/events", handlers.CreateEvent)
	router.HandleFunc("GET /api/v1/events", handlers.ListEvents)
	router.HandleFunc("GET /api/v1/events/{id}", handlers.ReadEvent)

	// Start the server (omitting error handling for brevity)
	http.ListenAndServe(":8080", router)
}
