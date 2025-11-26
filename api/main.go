package main

import (
	"kafka-go/orders"
	"log"
	"net/http"
)

func main() {
	orders.CreateTopic()
	go orders.Consumer()

	http.HandleFunc("/", orders.OrderHandler)
	http.HandleFunc("/order", orders.OrderHandler)
	log.Println("Server started on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
