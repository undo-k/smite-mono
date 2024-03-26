package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Get("/api/god/{godName}", app.GetGodByName)
	mux.Get("/api/gods", app.GetGodList)
	mux.Put("/api/putGod/{godName}", app.PutGodByName)
	mux.Get("/api/triggerAggregator/{number}", app.triggerAggregator)

	return mux
}
