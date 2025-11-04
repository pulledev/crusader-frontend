package main

import (
	"log"
	"net/http"

	"github.com/pulledev/crusader-frontend/crusader-back/api"
	"github.com/pulledev/crusader-frontend/crusader-back/initializers"

	"github.com/getkin/kin-openapi/openapi3filter"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	initializers.LoadEnvVars()

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("Fehler beim Laden der Swagger-Spezifikation: %s", err)
	}

	server := api.Server{}

	r := http.NewServeMux()

	h := api.HandlerFromMux(server, r)

	options := &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		},
	}

	validatingHandler := middleware.OapiRequestValidatorWithOptions(swagger, options)(h)

	s := &http.Server{
		Handler: validatingHandler,
		Addr:    "0.0.0.0:8080",
	}

	log.Println("Server startet auf http://0.0.0.0:8080")
	log.Fatal(s.ListenAndServe())
}
