package main

import (
	"fmt"
	"log"
	"net/http"

	"go-micro.dev/v4/web"
)

func main() {
	// Create a new web service
	service := web.NewService(
		web.Name("helloworld"),
		web.Address(":8080"),
	)

	service.Init()

	// Set up a route
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Attach the handler to the service
	service.Handle("/", http.DefaultServeMux)

	// Run the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
