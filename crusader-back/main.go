package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pulledev/crusader-frontend/crusader-back/api"
	"github.com/pulledev/crusader-frontend/crusader-back/initializers"
)

func main() {
	initializers.LoadEnvVars()
	//db := initializers.GetDB()

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.Server{}

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	fmt.Println(os.Getenv("STRIPE_API_KEY"))

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
