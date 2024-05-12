package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service interface {
	PublishMessage(queueName string, message string) error
	StartConsumer(queueName string) error
}

type rabbitMQService struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewService(config Config) (Service, error) {
	conn, ch, err := NewRabbitMQSession(config)
	if err != nil {
		return nil, err
	}

	service := &rabbitMQService{conn: conn, ch: ch}
	if err := SetupExchange(ch); err != nil {
		return nil, err
	}
	queues := []string{"order.france", "order.germany", "order.greatbritain", "order.spain"}
	if err := SetupQueues(ch, queues); err != nil {
		return nil, err
	}
	return service, nil
}

func (r *rabbitMQService) PublishMessage(queueName, message string) error {
	return r.ch.PublishWithContext(
		context.TODO(),
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

func (r *rabbitMQService) StartConsumer(queueName string) error {
	return StartConsumer(r.ch, queueName)
}
