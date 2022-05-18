package main

import (
	"net/http"

	openapi "github.com/shipink/golang-openapi-generator"
)

type openAPIRoutes []openapi.Route

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

var OpenAPIRoutes = openAPIRoutes{
	openapi.Route{
		Method:          http.MethodGet,
		Path:            "/endpoint",
		Request:         new(Request),
		SuccessResponse: new(SuccessResponse),
		ErrorResponse:   new(ErrorResponse),
	},
}

func main() {
	openapi.Generate("My Example API", OpenAPIRoutes)
}
