package transcoder

import (
	"github.com/streadway/amqp"
	"log"
)

type Service struct{}

func Start(worker *Worker) *Service {
	go consume(worker.Messages)
	return &Service{}
}

func consume(messages <-chan amqp.Delivery) {
	for message := range messages {
		go handle(message)
		message.Ack(false)
	}
}

func handle(message amqp.Delivery) {
	log.Printf("Got a message: %v", message.Body)
}
