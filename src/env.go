package ringcentralchatbot

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnv returns env value
// env value loaded from .env or other sources
func GetEnv(propName string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	res := os.Getenv(propName)
	return res
}
