package main

import (
	"github.com/AustinZhouOWO/demo_idm/idm"
	"github.com/tendant/chi-demo/app"
)

func main() {
	myApp := app.Default()
	// myApp.Config.Metrics.Port = 9901
	myApp.R.Get("/hello", idm.HandleHello)
	myApp.Run()
}
