package utils

import (
	logger "gohot/pkg/helpers"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		logger.FatalError("Error loading .env file")
	}
}
