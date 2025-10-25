package main

import (
	"backend/api"
	"backend/initializers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	initializers.LoadEnvVars()
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()

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
