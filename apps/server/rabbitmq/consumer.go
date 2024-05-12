package rabbitmq

import (
	"log"

	"github.com/jaka-k/apps/server/ticket-broker/paymant"
	amqp "github.com/rabbitmq/amqp091-go"
)

func StartConsumer(ch *amqp.Channel, queueName string) error {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to register a consumer: %v\n", err)
		return err
	}

	go func() {
		for d := range msgs {
			log.Printf("Received Message: %s\n", d.Body)
			paymant.ProcessOrder("Just FOR NOW")
		}
	}()
	return nil
}
