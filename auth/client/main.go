package main

import (
	"context"
	"fmt"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
)

type HelloRequest struct {
}

type HelloResponse struct {
}

func main() {
	service := micro.NewService(micro.Name("hello.client"))
	service.Init()

	//create a context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"Token": "valid-token",
	})

	req := client.NewRequest("hello", "Greeter.Hello", &HelloRequest{})
	res := &HelloResponse{}
	// call the service
	if err := service.Client().Call(ctx, req, res); err != nil {
		log.Fatalf("Error call ing hello service: %v", err)
	}

	fmt.Println("Successfully called the hello service")
}
