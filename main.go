package main

import "myapp/pkg/infrastructure/config"

func main() {
	c := config.InitApplication()
	c.App.ListenAndServe()
}