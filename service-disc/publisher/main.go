package main

import (
	"log"
	"time"

	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
)

func main() {
	//initialize the service

	service := micro.NewService(micro.Name("example.publisher"))
	service.Init()

	//start  the broker
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker connect error: %v", err)
	}

	//publish a message every 5 seconds
	go func() {
		tim := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-tim.C:
				msg := &broker.Message{
					Header: map[string]string{"id": "1"},
					Body:   []byte("hello from the publisher!"),
				}
				if err := broker.Publish("example.topic", msg); err != nil {
					log.Printf("Error publishing: %v", err)
				}
			}
		}

	}()

	// run the service
	if err := service.Run(); err != nil {
		log.Fatalf("bhai dikkat aa rhi hai: %v", err)
	}
}
