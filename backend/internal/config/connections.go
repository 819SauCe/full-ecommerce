package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(connectionString string) {
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error on connect in database: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error on DataBase: ", err)
	}

	log.Println("Connected in database")
}
