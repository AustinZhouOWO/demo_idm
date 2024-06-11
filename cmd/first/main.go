package main

import (
	"net/http"

	"github.com/tendant/chi-demo/app"
)

func main() {
	myApp := app.Default()
	// myApp.Config.Metrics.Port = 9901
	myApp.R.Get("/hello", handleHello)
	myApp.Run()
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
