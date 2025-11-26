package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{getBroker()},
		Topic:   "orders",
		GroupID: "orders-group",
	})

	defer reader.Close()

	fmt.Println("Consumer listening for orders")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		o := Order{}
		err = json.Unmarshal(msg.Value, &o)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received order: %s %d\n", o.Product, o.Quantity)

	}

}
