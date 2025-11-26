package orders

import (
	"fmt"
	"os"
	"time"

	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func getBroker() string {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "localhost:29092"
	}
	return broker
}

func Producer(order Order) {
	data, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{getBroker()},
		Topic:   "orders",
	})
	defer writer.Close()

	time.Sleep(1 * time.Second)

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(order.Product),
		Value: data,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Producer sent the order to Kafka")
}
