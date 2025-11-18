package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Postgres_uri string
var Password_hash_key string
var Jwt_secret string

func Load() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error on get .env file!")
	}

	//POSTGRES
	postgre_user := os.Getenv("POSTGRES_USER")
	postgre_password := os.Getenv("POSTGRES_PASSWORD")
	postgre_host := os.Getenv("POSTGRES_HOST")
	postgre_port := os.Getenv("POSTGRES_PORT")
	postgre_db := os.Getenv("POSTGRES_DB")
	postgre_ssl := os.Getenv("POSTGRES_SSL")

	if postgre_db == "" || postgre_host == "" || postgre_password == "" || postgre_port == "" || postgre_ssl == "" || postgre_user == "" {
		log.Fatal("Error on loading postgre data")
	}

	Postgres_uri = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", postgre_user, postgre_password, postgre_host, postgre_port, postgre_db, postgre_ssl)

	//PASSWORD HASH SECRET
	Password_hash_key = os.Getenv("HASH_PASSWORD_KEY")
	if Password_hash_key == "" {
		log.Fatal("Error on loading hash password secret")
	}

	Jwt_secret = os.Getenv("JWT_SECRET")
	if Jwt_secret == "" {
		log.Fatal("Error on loading jwt secret")
	}
}
