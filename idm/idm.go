package idm

import (
	"net/http"

	"github.com/go-chi/render"
)

type Hello struct {
	Name string `json:"name"`
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	result := Hello{
		Name: "Austin Zhou",
	}
	render.JSON(w, r, result)
}
