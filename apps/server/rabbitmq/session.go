package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Address string
}

func NewRabbitMQSession(cfg Config) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(cfg.Address)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		conn.Close() // Close connection if channel creation fails
		return nil, nil, err
	}
	return conn, ch, nil
}

func SetupQueues(ch *amqp.Channel, queues []string) error {
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
			log.Fatalf("Failed to declare a queue: %s", err)
			return err
		}
		err = ch.QueueBind(
			queueName,
			queueName, // Ensure correct routing key in publish
			"ticket-broker-exchange",
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("Failed to bind a queue: %s", err)
			return err
		}
	}
	return nil
}

func SetupExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"ticket-broker-exchange",
		"direct", // Changed type to 'direct' for more explicit routing
		true,
		false,
		false,
		false,
		nil,
	)
}
