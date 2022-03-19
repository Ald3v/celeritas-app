package config

import (
	"github.com/go-chi/chi/v5"
)

func (app * application) routes() *chi.Mux {

	app.get("/health", app.Handlers.HealthCheck)

	return app.App.Routes
}
