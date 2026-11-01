package broker

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueName string

const (
	RequestCall                QueueName = "request_call"
	AnswerCall                 QueueName = "answer_call"
	NewIceCandidate            QueueName = "ice_candidate"
	EndCall                    QueueName = "end_call"
	RequestCallViaNotification QueueName = "request_call_via_notification"

	SaveMessageQueue       QueueName = "save_message_queue"
	SendMsgViaWs           QueueName = "send_msg_via_websocket"
	SendMessageQueue       QueueName = "send_message_queue"
	SendMsgViaNotification QueueName = "send_msg_via_notification"

	OnlineStatusChange QueueName = "online_status_change"
)

func (r *RabbitMq) Run(queue QueueName, processor func(headers amqp.Table, msg []byte)) {
	// Declare the queue (ensure it exists)
	_, err := r.channel.QueueDeclare(
		string(queue), true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := r.channel.Consume(
		string(queue), // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	for msg := range msgs {
		processor(msg.Headers, msg.Body)
	}
}
