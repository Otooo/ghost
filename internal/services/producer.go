package services

import (
	"context"
	"time"

	"gohot/configs"
	logger "gohot/pkg/helpers"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessage(channel *amqp.Channel, msg string) {
	queue, err := channel.QueueDeclare(
		configs.QueueLocName, // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // table arguments
	)
	logger.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         []byte(msg),
	}
	err = channel.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		body,
	)
	logger.FailOnError(err, "Failed to publish a message")
	logger.SentConfirm(msg)
}
