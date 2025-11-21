package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConnectPostgres(connectionString string) {
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error on connect in Postgres: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error on Postgres: ", err)
	}

	log.Println("Connected to Postgres!")
}
