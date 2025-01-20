package utils

import (
	"os"

	logger "gohot/pkg/helpers"

	amqp "github.com/rabbitmq/amqp091-go"
)

func GetRabbitURL() string {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	// Verificar se a variável foi carregada corretamente
	if rabbitURL == "" {
		logger.FatalError("RABBITMQ_URL environment variable not set")

		os.Exit(1)
		return ""
	}

	return rabbitURL
}

func ConnectToRabbit(url string) *amqp.Connection {
	connection, err := amqp.Dial(url)
	logger.FailOnError(err, "Failed to connect to RabbitMQ")

	return connection
}

func ConnectToChannel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()
	logger.FailOnError(err, "Failed to open a channel")

	return channel
}
