package main

import (
	"gohot/internal/services"
	"gohot/utils"
	"math/rand"
)

func main() {
	utils.LoadEnv()
	rabbitURL := utils.GetRabbitURL()
	connection := utils.ConnectToRabbit(rabbitURL)
	defer connection.Close()
	channel := utils.ConnectToChannel(connection)
	defer channel.Close()

	services.SendMessage(channel, "Aye Sir "+string(rand.Intn(999))+"!!!")
}
