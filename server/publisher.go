package main

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	channel *amqp.Channel
}

func NewRabbitMQPublisher(ch *amqp.Channel) *RabbitMQPublisher {
	return &RabbitMQPublisher{channel: ch}
}

func (r *RabbitMQPublisher) PublishMessage(queue string, msg []byte) error {
	return r.channel.PublishWithContext(context.TODO(), "", queue, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	})
}
