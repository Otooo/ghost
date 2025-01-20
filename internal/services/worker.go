package services

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gohot/configs"
	logger "gohot/pkg/helpers"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ReadMessage(channel *amqp.Channel) {
	queue, err := channel.QueueDeclare(
		configs.QueueLocName, // queue
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	logger.FailOnError(err, "Failed to declare a consume queue")

	messages, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	logger.FailOnError(err, "Failed to register a consumer")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	var forever chan bool

	go func() {
		for msg := range messages {
			logger.Message(string(msg.Body), 1)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-sigchan

	log.Printf("interrupted, shutting down")
	forever <- true
}
