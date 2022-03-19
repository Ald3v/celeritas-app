package routes

import (
	"myapp/pkg/infrastructure/config"

	"github.com/go-chi/chi/v5"
)

func AddRoutes(app * config.Application) *chi.Mux {

	get(app,"/health", app.Handlers.HealthCheck)

	return app.App.Routes
}
