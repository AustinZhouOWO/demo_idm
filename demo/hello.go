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

func (handler *Handle) HandleHello(w http.ResponseWriter, r *http.Request, params HandleHelloParams) *Response {
	if params.Name == nil {
		return &Response{
			Code:        http.StatusBadRequest,
			contentType: "no name param",
		}
	}

	result := Hello{
		Name: *params.Name,
	}

	return HandleHelloJSON200Response(Trim(result.Name))
}
