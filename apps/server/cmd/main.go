package main

import (
	"fmt"
	"log"

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

	rabbitMQService, err := initializeRabbitMQ(config)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ service: %v", err)
	}

	server := api.NewAPIServer(":"+port, rabbitMQService)
	server.Run()
}

func initializeRabbitMQ(cfg *config.Config) (rabbitmq.Service, error) {
	address := getRabbitMQAddress(cfg)
	amqpConfig := rabbitmq.Config{
		Address: address,
	}

	return rabbitmq.NewService(amqpConfig)
}

func getRabbitMQAddress(config *config.Config) string {
	if config.AWSRabbitMQAMQP != "" || config.Environment == "dev" {
		return "amqp://guest:guest@localhost:5672"
	}
	return config.AWSRabbitMQAMQP
}
