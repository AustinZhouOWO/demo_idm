package demo

import (
	"net/http"
	"strings"
)

type Hello struct {
	Name string `json:"name"`
}

func Trim(str string) []string {
	return strings.Split(str, " ")
}

func (handler *Handle) HandleHello(w http.ResponseWriter, r *http.Request) *Response {
	result := Hello{
		Name: "Austin Zhou",
	}

	return HandleHelloJSON200Response(Trim(result.Name))
}
