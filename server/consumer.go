package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	channel *amqp.Channel
}

func NewConsumer(channel *amqp.Channel) *Consumer {
	return &Consumer{channel: channel}
}

func (c *Consumer) Consume() {
	msgs, err := c.channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Failed to register a consumer: %v\n", err)
		return
	}

	for d := range msgs {
		fmt.Printf("Received Message: %s\n", d.Body)
	}
}
