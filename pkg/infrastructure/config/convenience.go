package config

import (
	"net/http"
)

func (app *application) get( s string, h http.HandlerFunc) {
	app.App.Routes.Get(s, h)
}

func (app *application) post(s string, h http.HandlerFunc) {
	app.App.Routes.Post(s, h)
}

func (app *application) use( m ...func(http.Handler) http.Handler) {
	app.App.Routes.Use(m...)
}