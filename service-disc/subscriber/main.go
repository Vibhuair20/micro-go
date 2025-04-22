package main

import (
	"fmt"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
)

func main() {

	// initializinf the service
	service := micro.NewService(micro.Name("example.subscriber"))
	service.Init()

	//start the broker
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker could not connect: %v", err)
	}

	// subscribe to messages
	_, err := broker.Subscribe("example.topic", func(p broker.Event) error {
		fmt.Printf("Received message: %s\n", string(p.Message().Body))
		return nil
	})

	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

	// run the service
	if err := service.Run(); err != nil {
		log.Fatalf("bhai dikkat aa rhi hai: %v", err)
	}
}
