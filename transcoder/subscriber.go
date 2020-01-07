package transcoder

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
)

type Worker struct {
	Channel     *amqp.Channel
	Messages    <-chan amqp.Delivery
	shutdown    chan bool
	workerGroup sync.WaitGroup
}

type Handler func(message amqp.Delivery)

func (conn *Connection) InitializeWorkQueueSubscriber(routingKey string, handler Handler) *Worker {
	workChannel, err := conn.Connection.Channel()
	if err != nil {
		log.Fatalf("Error opening work queue subscriber channel: %v", err)
	}

	err = workChannel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalf("Error setting channel QoS: %v", err)
	}

	queue, err := workChannel.QueueDeclare(
		routingKey,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Error declaring queue: %v", err)
	}

	messages, err := workChannel.Consume(
		queue.Name,
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	worker := Worker{
		Channel:     workChannel,
		Messages:    messages,
		shutdown:    make(chan bool),
		workerGroup: sync.WaitGroup{},
	}

	go worker.consume(handler)
	return &worker
}

func (worker *Worker) consume(handler Handler) {
	for {
		select {
		case message := <-worker.Messages:
			worker.workerGroup.Add(1)
			go worker.handleDelivery(handler, message)
		case <-worker.shutdown:
			log.Println("Stopped consuming messages.")
			return
		}
	}
}

func (worker *Worker) handleDelivery(handler Handler, message amqp.Delivery) {
	defer worker.workerGroup.Done()

	handler(message)
	err := message.Ack(false)
	if err != nil {
		log.Fatalf("Failed to ack message: %v", err)
	}
}

func (worker *Worker) Close() {
	worker.shutdown <- true
	worker.workerGroup.Wait()
	worker.Channel.Close()
}
