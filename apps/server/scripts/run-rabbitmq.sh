#!/bin/bash

container_name="rabbitmq-dev"

# Check if RabbitMQ is already running
if [ "$(docker ps -q -f name=$container_name)" ]; then
    echo "RabbitMQ is already running."
else
    # Check if the container exists but is stopped
    if [ "$(docker ps -aq -f status=exited -f name=$container_name)" ]; then
        echo "Starting existing RabbitMQ container..."
        docker start $container_name
    else
        echo "Creating and starting new RabbitMQ container..."
        docker run -d --name $container_name -p 5672:5672 -p 15672:15672 rabbitmq:management
    fi
fi
