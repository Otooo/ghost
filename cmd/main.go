package main

import (
	"os"

	"gohot/internal/services"
	logger "gohot/pkg/helpers"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	loadEnv()
	rabbitURL := getRabbitURL()
	connection := connectToRabbit(rabbitURL)
	defer connection.Close()
	channel := connectToChannel(connection)
	defer channel.Close()

	services.SendMessage(channel, "Aye Sir !!!")
}

func loadEnv() {
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		logger.FatalError("Error loading .env file")
	}
}

func getRabbitURL() string {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	// Verificar se a variável foi carregada corretamente
	if rabbitURL == "" {
		logger.FatalError("RABBITMQ_URL environment variable not set")

		os.Exit(1)
		return ""
	}

	return rabbitURL
}

func connectToRabbit(url string) *amqp.Connection {
	connection, err := amqp.Dial(url)
	logger.FailOnError(err, "Failed to connect to RabbitMQ")

	return connection
}

func connectToChannel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()
	logger.FailOnError(err, "Failed to open a channel")

	return channel
}
