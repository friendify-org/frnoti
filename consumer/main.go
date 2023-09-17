package base_consumer

import (
	"fmt"
	"main/config"

	"github.com/rabbitmq/amqp091-go"
)

type EventSubcribe chan []byte
type EventHandler func(data []byte)

func New(queueName string) EventSubcribe {

	event := make(EventSubcribe)

	msgs, err := config.RabbitChannel.Consume(
		queueName,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		fmt.Printf("err when declare queue: %v\n", err)
	}

	go event.messageHandler(msgs)

	return event
}

func (event EventSubcribe) messageHandler(msgs <-chan amqp091.Delivery) {
	for d := range msgs {
		event <- d.Body
	}
}

func (event EventSubcribe) SubcribeEvent(eventHandler EventHandler) {

	go func() {
		for {
			eventHandler(<-event)
		}
	}()
}
