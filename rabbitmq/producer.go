package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Producer(msg string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(fmt.Sprintf("Failed to establish connection with amqp queue: %v", err))
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf("Failed to create amqp channel: %v", err))
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to create declare queue: %v", err))
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to publish: %v", err))
	}
}
