package main

import (
	"full-ecommerce/internal/auth"
	"full-ecommerce/internal/config"
	"full-ecommerce/internal/product" // <--- adiciona isso
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	config.Load()
	config.ConnectPostgres(config.Postgres_uri)
	config.ConnectMongoDB(config.Mongo_uri, "full_ecommerce")
	config.ConnectRabbitMQ()
	config.ConnectRedis()

	//Routers
	mux := http.NewServeMux()
	auth.RegisterAuthRoutes(mux)
	product.RegisterProductRoutes(mux)

	handler := config.CORS(mux)
	http.ListenAndServe(":8080", handler)
}
