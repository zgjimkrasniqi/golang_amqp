package main

import (
	"example.com/m/v2/event"
	"github.com/streadway/amqp"
	"os"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	consumer, err := event.NewConsumer(connection)
	if err != nil {
		panic(err)
	}

	consumer.Listen(os.Args[1:])

	/*
		#t1> go run main.go log.WARN log.ERROR
		#t2> go run main.go log.*
	*/
}
