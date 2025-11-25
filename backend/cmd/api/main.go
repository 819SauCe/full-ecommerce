package main

import (
	auth "full-ecommerce/internal/_auth"
	banner "full-ecommerce/internal/_banner"
	cart "full-ecommerce/internal/_cart"
	product "full-ecommerce/internal/_product"
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

	//Routers
	mux := http.NewServeMux()
	auth.RegisterAuthRoutes(mux)
	product.RegisterProductRoutes(mux)
	cart.RegisterCartRoutes(mux)
	banner.RegisterBannerRoutes(mux)

	handler := config.CORS(mux)
	http.ListenAndServe(":8080", handler)
}
