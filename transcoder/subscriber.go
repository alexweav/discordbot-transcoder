package transcoder

import (
	"github.com/streadway/amqp"
	"log"
)

type Worker struct {
	Channel *amqp.Channel
	Messages <-chan amqp.Delivery
}

func InitializeWorkQueueSubscriber(conn *Connection, routingKey string) *Worker {
	workChannel, err := conn.Connection.Channel()
	if err != nil {
		log.Fatalf("Error opening work queue subscriber channel: %v", err)
	}

	err = workChannel.Qos(
		1,  // prefetch count
		0,  // prefetch size
		false,  // global
	)
	if err != nil {
		log.Fatalf("Error setting channel QoS: %v", err)
	}

	queue, err := workChannel.QueueDeclare(
		routingKey,
		false,	// durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Error declaring queue: %v", err)
	}

	messages, err := workChannel.Consume(
		queue.Name,
		"",	// consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	
	return &Worker{
		Channel: workChannel,
		Messages: messages,
	}
}

func (worker *Worker) Close() {
	// TODO: do something with unconsumed messages
	worker.Channel.Close()
}