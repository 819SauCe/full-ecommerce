package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

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
