package orders

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func CreateTopic() {
	conn, err := kafka.Dial("tcp", getBroker())
	if err != nil {
		log.Fatal(err)
	}
	// controller for leader election
	controller, err := conn.Controller()
	if err != nil {
		log.Fatal(err)
	}
	controllerConn, err := kafka.Dial("tcp", fmt.Sprintf("%s:%d", controller.Host, controller.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer controllerConn.Close()

	topic := "orders"
	err = controllerConn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Topic created")
}
