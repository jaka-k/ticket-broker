package main

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	config, err := LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return
	}

	if config.Environment == "dev" {
		fmt.Println("Running in development mode")
	} else if config.Environment == "prod" {
		fmt.Println("Running in production mode")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server := NewAPIServer(":" + port)
	server.Run()

	amqp_address := config.AWSRabbitMQAMQP
	if amqp_address != "" {
		amqp_address = "amqp://guest:guest@localhost:5672"
	}

	conn, err := amqp.Dial(amqp_address)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	consumer := NewConsumer(ch)
	forever := make(chan bool)

	go func() {
		consumer.Consume()
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
