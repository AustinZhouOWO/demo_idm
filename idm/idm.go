package idm

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

type Hello struct {
	Name string `json:"name"`
}

func Trim(str string) []string {
	return strings.Split(str, " ")
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	result := Hello{
		Name: "Austin Zhou",
	}
	render.JSON(w, r, Trim(result.Name))
}
