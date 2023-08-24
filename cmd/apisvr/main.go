package main

import (
	"github.com/AndrewDonelson/go-http2-api/pkg/api"
)

func main() {
	api.Server.Initialize(":8080", "./certs")

	api.Server.AddRoute(&api.APIRoute{
		Version:  1,
		SubRoute: "subroute",
		Name:     "test",
		Method:   "GET",
		Handler:  api.TestHandler,
	})

}
