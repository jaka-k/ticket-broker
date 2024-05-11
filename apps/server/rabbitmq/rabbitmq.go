package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	Address string
}

func NewRabbitMQSession(cfg RabbitMQConfig) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(cfg.Address)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return nil, nil, err
	}

	return conn, ch, nil
}

func SetupQueues(ch *amqp.Channel) error {
	queues := []string{"FranceQueue", "GermanyQueue", "GreatBritainQueue", "SpainQueue"}
	for _, queueName := range queues {
		_, err := ch.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
