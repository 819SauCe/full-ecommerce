package config

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *sql.DB
var MongoClient *mongo.Client
var MongoDB *mongo.Database
var RedisClient *redis.Client
var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

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

func ConnectMongoDB(uri, dbName string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error creating MongoDB client: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	}

	log.Println("Connected to MongoDB!")

	MongoClient = client
	MongoDB = client.Database(dbName)
}

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPass,
		DB:       0,
	})

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Connected to Redis!")
}

func ConnectRabbitMQ() {
	conn, err := amqp.Dial(Rabbitmq_uri)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open channel:", err)
	}

	RabbitConn = conn
	RabbitChannel = ch

	log.Println("Connected to RabbitMQ!")
}
