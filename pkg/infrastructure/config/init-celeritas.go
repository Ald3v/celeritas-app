package config

import (
	"log"
	"myapp/pkg/infrastructure/handlers"
	"myapp/pkg/infrastructure/persistence"
	"os"

	"github.com/ald3v/celeritas"
)

func InitApplication() *application {
	path, err := os.Getwd()
	if err !=nil{
		log.Fatal(err)
	}

	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err !=nil{
		log.Fatal(err)
	}

	cel.AppName = "app"
	
	myHandlers := &handlers.Handlers{
		App: cel,
	}

	app := &application{
		App: cel,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	app.Models = persistence.New(app.App.DB.Pool)

	return app

}