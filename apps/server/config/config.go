package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	Environment     string
	AWSRabbitMQUser string
	AWSRabbitMQPass string
	AWSRabbitMQAMQP string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:            os.Getenv("PORT"),
		Environment:     os.Getenv("ENVIRONMENT"),
		AWSRabbitMQUser: os.Getenv("AWSRABBITMQ_USER"),
		AWSRabbitMQPass: os.Getenv("AWSRABBITMQ_PASS"),
		AWSRabbitMQAMQP: os.Getenv("AWSRABBITMQ_AMQP"),
	}, nil
}
