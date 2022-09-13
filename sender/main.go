package main

import (
	"example.com/m/v2/event"
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	emitter, err := event.NewEventEmitter(conn)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 10; i++ {
		emitter.Push(fmt.Sprintf("[%d] - %s", i, os.Args[1]), os.Args[1])
	}

	/*
		#t3> go run sender.go log.WARN
		#t3> go run sender.go log.ERROR
		#t3> go run sender.go log.INFO
	*/
}
