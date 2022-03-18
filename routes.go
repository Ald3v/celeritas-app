package main

import (
	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {

	a.get("/health", a.Handlers.HealthCheck)

	return a.App.Routes
}
