package main

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"log"
)

func main() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9093"},
		GroupID:  "consumer-group-id",
		Topic:    "test-topic",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("new message %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

}
