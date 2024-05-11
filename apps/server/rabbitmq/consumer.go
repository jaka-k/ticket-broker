package rabbitmq

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

func (c *Consumer) Consume(queueName string) error {
	msgs, err := c.channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Failed to register a consumer: %v\n", err)
		return err
	}

	for d := range msgs {
		fmt.Printf("Received Message: %s\n", d.Body)
	}
	return nil
}

func StartConsumer(ch *amqp.Channel, queue string) {
	consumer := NewConsumer(ch)

	if err := consumer.Consume(queue); err != nil {
		fmt.Printf("Failed to start consumer for queue %s: %v", queue, err)
	}

}
