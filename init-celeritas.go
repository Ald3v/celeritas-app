package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"os"

	"github.com/ald3v/celeritas"
)

func initApplication() *application {
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

	app := &application{
		App: cel,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)

	return app

}