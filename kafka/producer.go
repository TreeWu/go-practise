package main

import (
	"context"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"math"
	"strconv"
)

const (
	KAFKA_SERVER = "192.168.123.189:9093"
	Topic        = "test"
)

func Producer() {
	writer := kafka.Writer{
		Addr:     kafka.TCP(KAFKA_SERVER),
		Topic:    Topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()
	for i := 0; i < math.MaxInt; i++ {
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte("key" + strconv.Itoa(i)),
			Value: []byte("Value" + strconv.Itoa(i)),
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(" 发送  " + strconv.Itoa(i))
	}
}
