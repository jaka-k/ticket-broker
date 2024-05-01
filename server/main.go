package main

import (
	"fmt"
)

func main() {

	config, err := LoadConfig()
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

	amqpConfig := RabbitMQConfig{
		Address: getRabbitMQAddress(config),
	}

	conn, ch, err := NewRabbitMQ(amqpConfig)
	if err != nil {
		fmt.Printf("Failed to setup RabbitMQ: %v", err)
		panic(err)
	}
	defer conn.Close()
	defer ch.Close()

	server := NewAPIServer(":" + port)
	server.Run()
}

func getRabbitMQAddress(config *Config) string {
	if config.AWSRabbitMQAMQP != "" || config.Environment == "dev" {
		return "amqp://guest:guest@localhost:5672"
	}
	return config.AWSRabbitMQAMQP
}
