package main

import (
	"gohot/internal/services"
	"gohot/utils"
)

func main() {
	utils.LoadEnv()
	rabbitURL := utils.GetRabbitURL()
	connection := utils.ConnectToRabbit(rabbitURL)
	defer connection.Close()
	channel := utils.ConnectToChannel(connection)
	defer channel.Close()

	services.ReadMessage(channel)
}
