package routes

import (
	"myapp/pkg/infrastructure/config"
	"net/http"
)

func get(app *config.Application, s string, h http.HandlerFunc) {
	app.App.Routes.Get(s, h)
}

func post(app *config.Application, s string, h http.HandlerFunc) {
	app.App.Routes.Post(s, h)
}

func use(app *config.Application, m ...func(http.Handler) http.Handler) {
	app.App.Routes.Use(m...)
}