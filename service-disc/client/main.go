package main

import (
	"context"
	"fmt"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	// create a new service
	service := micro.NewService(micro.Name("ServiceB"))

	//initializing the service
	service.Init()

	// request message
	req := service.Client().NewRequest("serviceA", "Greeter.Hello", &map[string]interface{}{}, client.WithContentType("application/json"))
	// to store the response
	rsp := &map[string]interface{}{}
	// call the service
	if err := service.Client().Call(context.Background(), req, rsp); err != nil {
		log.Fatalf("Error calling service A: %v", err)
	}

	fmt.Println("Successfully called service A")
}
