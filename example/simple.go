package main

import (
	"net/http"

	openapi "github.com/shipink/golang-openapi-generator"
)

type Request struct {
	Name string `json:"name"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	var routes []openapi.Route

	routes = append(routes, openapi.Route{
		Method:          http.MethodGet,
		Path:            "/endpoint",
		Request:         new(Request),
		SuccessResponse: new(SuccessResponse),
		ErrorResponse:   new(ErrorResponse),
	})

	openapi.Generate("My Example API", routes)
}
