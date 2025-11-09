package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to read .env file - " + err.Error())
	}
	log.Println(".env file is loaded")
}
