package broker

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMq struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
}

func NewRabbitMQConn(url string) *RabbitMq {
	conn, err := amqp091.Dial(url)
	if err != nil {
		fmt.Printf("failed to connect to RabbitMQ: %v", err)
		return nil
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close() // Close connection if channel fails
		fmt.Printf("failed to open a channel: %v", err)
		return nil
	}

	log.Println("Successfully connected to RabbitMQ")
	return &RabbitMq{
		connection: conn,
		channel:    ch,
	}
}
