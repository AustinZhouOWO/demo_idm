package main

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/tendant/chi-demo/app"
)

func main() {
	myApp := app.Default()
	// myApp.Config.Metrics.Port = 9901
	myApp.R.Get("/hello", handleHello)
	myApp.Run()
}

type Hello struct {
	Name string
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	result := Hello{
		Name: "Austin Zhou",
	}
	render.JSON(w, r, result)
}
