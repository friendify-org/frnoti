package base_consumer

import (
	"fmt"
	"main/config"

	"github.com/rabbitmq/amqp091-go"
)

type EventSubcribe map[string]chan []byte
type EventHandler func(data []byte)

func New(queueName string) EventSubcribe {

	events := EventSubcribe{}

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

	go events.messageHandler(msgs)

	return events
}

func (events EventSubcribe) messageHandler(msgs <-chan amqp091.Delivery) {
	for d := range msgs {
		message := d.Headers["message"].(string) // The message recieved <- body
		if events[message] == nil {
			fmt.Printf("[Warning] Unsubcribe message: %v\n", message)
			continue
		}
		events[message] <- d.Body
	}
}

func (events EventSubcribe) SubcribeEvent(eventName string, eventHandler EventHandler) {

	events[eventName] = make(chan []byte)
	go func() {
		for {
			eventHandler(<-events[eventName])
		}
	}()
}
