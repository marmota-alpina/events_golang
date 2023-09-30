package main

import (
	"events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			panic(err)
		}
	}(ch)

	err = rabbitmq.Publish(ch, "Order 002", "amq.direct")
	if err != nil {
		panic(err)
		return
	}
}
