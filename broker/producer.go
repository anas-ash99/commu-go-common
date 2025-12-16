package broker

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Publish implements domain.MessageProducer
func (r *RabbitMq) Publish(ctx context.Context, queueName QueueName, msg []byte, headers amqp.Table) error {

	_, err := r.channel.QueueDeclare(
		string(queueName), // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// 3. Publish the message
	err = r.channel.PublishWithContext(ctx,
		"",
		string(queueName),
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         msg,
			Headers:      headers,
			DeliveryMode: amqp.Persistent, // Save to disk
			Timestamp:    time.Now(),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
