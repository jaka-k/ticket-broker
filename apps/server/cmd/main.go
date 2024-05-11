package main

import (
	"fmt"

	"github.com/jaka-k/apps/server/ticket-broker/api"
	"github.com/jaka-k/apps/server/ticket-broker/config"
	"github.com/jaka-k/apps/server/ticket-broker/rabbitmq"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return
	}

	environment := config.Environment

	if environment == "dev" {
		fmt.Println("Running in development mode")
	} else if environment == "prod" {
		fmt.Println("Running in production mode")
	}

	port := config.Port
	if port == "" {
		port = "3000"
	}

	amqpConfig := rabbitmq.RabbitMQConfig{
		Address: getRabbitMQAddress(config),
	}

	conn, ch, err := rabbitmq.NewRabbitMQSession(amqpConfig)
	if err != nil {
		fmt.Printf("Failed to setup RabbitMQ: %v", err)
		panic(err)
	}
	defer conn.Close()
	defer ch.Close()

	if err := rabbitmq.SetupQueues(ch); err != nil {
		fmt.Printf("Failed to declare queues: %v", err)
	}

	publisher := rabbitmq.NewRabbitMQPublisher(ch)

	queues := []string{"FranceQueue", "GermanyQueue", "GreatBritainQueue", "SpainQueue"}
	for _, queue := range queues {
		go rabbitmq.StartConsumer(ch, queue)
	}

	server := api.NewAPIServer(":"+port, publisher)
	server.Run()
}

func getRabbitMQAddress(config *config.Config) string {
	if config.AWSRabbitMQAMQP != "" || config.Environment == "dev" {
		return "amqp://guest:guest@localhost:5672"
	}
	return config.AWSRabbitMQAMQP
}
