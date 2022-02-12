package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Reciever() {
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to register consumer: %v", err))
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved a message %s\n", d.Body)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}
