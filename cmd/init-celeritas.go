package main

import (
	"log"
	"myapp/pkg/infrastructure/config"
	"myapp/pkg/infrastructure/handlers"
	"myapp/pkg/infrastructure/persistence"
	"myapp/pkg/infrastructure/routes"
	"os"

	"github.com/ald3v/celeritas"
)

func initApplication() *config.Application {
	path, err := os.Getwd()
	if err !=nil{
		log.Fatal(err)
	}

	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err !=nil{
		log.Fatal(err)
	}

	cel.AppName = "myapp"
	
	myHandlers := &handlers.Handlers{
		App: cel,
	}

	app := &config.Application{
		App: cel,
		Handlers: myHandlers,
	}

	app.App.Routes = routes.AddRoutes(app)

	app.Models = persistence.New(app.App.DB.Pool)

	return app

}