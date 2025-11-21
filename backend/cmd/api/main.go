package main

import (
	"full-ecommerce/internal/auth"
	"full-ecommerce/internal/config"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	config.Load()
	config.ConnectPostgres(config.Postgres_uri)
	config.ConnectMongoDB(config.Mongo_uri, "full_ecommerce")
	config.ConnectRabbitMQ()
	config.ConnectRedis()
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/register", auth.RegisterHandler)
	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/logout", auth.LogoutHandler)
	mux.HandleFunc("/auth/me", auth.MeHandler)

	handler := config.CORS(mux)
	http.ListenAndServe(":8080", handler)
}
