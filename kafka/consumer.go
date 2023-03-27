package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{KAFKA_SERVER},
		Topic:   Topic,
	})
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		value := message.Value
		fmt.Println("consumer : " + string(value))
	}
}
