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
var Mongo_uri string
var RedisAddr string
var RedisPass string
var Rabbitmq_uri string

func Load() {
	paths := []string{
		"./.env",
		"../.env",
		"../../.env",
		"../../../.env",
		"../../../../.env",
	}

	var err error
	for _, p := range paths {
		err = godotenv.Load(p)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Fatal("Error loading .env file!")
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

	//MONGO
	mongo_user := os.Getenv("MONGO_USER")
	mongo_password := os.Getenv("MONGO_PASSWORD")
	mongo_port := os.Getenv("MONGO_PORT")

	if mongo_user == "" || mongo_password == "" || mongo_port == "" {
		log.Fatal("Error on loading mongo data")
	}

	Mongo_uri = fmt.Sprintf("mongodb://%s:%s@localhost:%s", mongo_user, mongo_password, mongo_port)

	//REDIS
	redisPort := os.Getenv("REDIS_PORT")
	RedisPass = os.Getenv("REDIS_PASSWORD")

	if redisPort == "" {
		log.Fatal("Error on loading redis")
	}

	RedisAddr = "localhost:" + redisPort

	//RABBITMQ
	rabbit_user := os.Getenv("RABBIT_USER")
	rabbit_pass := os.Getenv("RABBIT_PASSWORD")
	rabbit_host := os.Getenv("RABBIT_HOST")
	rabbit_port := os.Getenv("RABBITMQ_PORT")

	if rabbit_host == "" || rabbit_pass == "" || rabbit_port == "" || rabbit_user == "" {
		log.Fatal("Error on loading rabbitmq")
	}

	Rabbitmq_uri = fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbit_user, rabbit_pass, rabbit_host, rabbit_port)

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
