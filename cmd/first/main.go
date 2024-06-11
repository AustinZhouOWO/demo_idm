package main

import (
	"github.com/AustinZhouOWO/demo_idm/demo"
	"github.com/tendant/chi-demo/app"
)

func main() {
	myApp := app.Default()
	//myApp.Config.Metrics.Port = 9901
	handler := demo.Handle{}
	demo.Routes(myApp.R, handler)
	myApp.Run()
}
