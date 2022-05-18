package openapi

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/swaggest/openapi-go/openapi3"
)

type OpenAPI interface {
	registerRoutes([]Route)
	operation(method, path string, request, successResponse interface{}, errorResponse interface{}) error
	save() error
}

type openapi struct {
	Reflector *openapi3.Reflector
}

type Route struct {
	Method          string
	Path            string
	Request         interface{}
	SuccessResponse interface{}
	ErrorResponse   interface{}
}

func Generate(title string, routes []Route) error {
	openAPI := New(title)
	openAPI.registerRoutes(routes)
	return openAPI.save()
}

func New(title string) OpenAPI {
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.1.0"}
	reflector.Spec.Info.
		WithTitle(title).
		WithVersion("1.0.0").
		WithDescription(fmt.Sprintf("This is documentation for %s as openapi format.", title))

	return &openapi{Reflector: &reflector}
}

func (o *openapi) registerRoutes(routes []Route) {
	for _, route := range routes {
		if err := o.operation(route.Method, route.Path, route.Request, route.SuccessResponse, route.ErrorResponse); err != nil {
			log.Printf("failed to register route: %v", err)
		}
	}
}

func (o *openapi) save() error {
	schema, err := o.Reflector.Spec.MarshalJSON()
	if err != nil {
		return err
	}

	os.WriteFile("openapi.json", schema, 0644)

	return nil
}

func (o *openapi) operation(method, path string, request, successResponse interface{}, errorResponse interface{}) error {
	createOperation := openapi3.Operation{}
	_ = o.Reflector.SetRequest(&createOperation, request, method)
	_ = o.Reflector.SetJSONResponse(&createOperation, successResponse, http.StatusOK)
	_ = o.Reflector.SetJSONResponse(&createOperation, errorResponse, http.StatusBadRequest)
	_ = o.Reflector.SetJSONResponse(&createOperation, errorResponse, http.StatusUnauthorized)
	_ = o.Reflector.SetJSONResponse(&createOperation, errorResponse, http.StatusForbidden)
	_ = o.Reflector.SetJSONResponse(&createOperation, errorResponse, http.StatusNotFound)
	_ = o.Reflector.Spec.AddOperation(method, path, createOperation)

	return nil
}
