package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"rabbit/internal/order/infra/database"
	"rabbit/internal/order/usecase"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func consumeData(msg amqp.Delivery, uc usecase.CalculateFinalPriceUseCase) {
	var inputDTO usecase.OrderInputDTO
	err := json.Unmarshal(msg.Body, &inputDTO)
	if err != nil {
		panic(err)
	}
	outputDTO, err := uc.Execute(inputDTO)
	if err != nil {
		panic(err)
	}
	msg.Ack(false)
	fmt.Println(outputDTO)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	//out := make(chan amqp.Delivery) // channel
	//go rabbitmq.Consume(ch, out)    // T2

	for msg := range msgs {
		consumeData(msg, uc)
	}
}
