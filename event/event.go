package event

import "github.com/streadway/amqp"

func getExchangeName() string {
	return "logs_topic"
}

func DeclareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare("", false, false, true, false, nil)
}

func DeclareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(getExchangeName(), "topic", true, false, false, false, nil)
}
