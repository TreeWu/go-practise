package main

import (
	"context"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"strconv"
)

const (
	KAFKA_SERVER = "192.168.123.189:9093"
	Topic        = "test"
)

func main() {
	writer := kafka.Writer{
		Addr:     kafka.TCP(KAFKA_SERVER),
		Topic:    Topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()
	for i := 0; i < 100; i++ {
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte("key" + strconv.Itoa(i)),
			Value: []byte("Value" + strconv.Itoa(i)),
		})
		if err != nil {
			log.Fatal(err)
		}
	}

}
