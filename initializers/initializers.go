package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't find .env File")
	}
}
