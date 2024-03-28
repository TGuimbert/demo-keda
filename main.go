package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	var leader = func(n int) string {
		return fmt.Sprintf("kafka-%v.kafka-headless.kafka.svc.cluster.local:9092", n)
	}
	groupID := "my-group"
	topic := "test"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{leader(0), leader(1), leader(2)},
		GroupID:  groupID,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		time.Sleep(1 * time.Second)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

}
