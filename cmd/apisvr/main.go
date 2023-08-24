package main

import (
	"fmt"
	"net/http"

	api "github.com/AndrewDonelson/go-http2-api/pkg"
)

func handlerTestingHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	api.Server.Initialize(":8080", "../certs/")

	// Add routes here - https://localhost:8080/v1/testing/hello
	api.Server.AddRoute(&api.APIRoute{
		Version:  1,
		SubRoute: "testing",
		Name:     "hello",
		Method:   "GET",
		Handler:  handlerTestingHello,
	})

	// Start the server
	api.Server.Start()
}
