package main

import (
	"full-ecommerce/internal/auth"
	"full-ecommerce/internal/config"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	config.Load()
	config.Connect(config.Postgres_uri)

	http.HandleFunc("/auth/register", auth.RegisterHandler)

	http.ListenAndServe(":8080", nil)
}
