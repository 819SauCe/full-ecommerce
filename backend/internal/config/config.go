package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Postgre_uri string

func Load() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error on get .env file!")
	}

	Postgre_uri = os.Getenv("POSTGRE_URI")
}
