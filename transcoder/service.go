package transcoder

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Service struct {
	worker *Worker
}

func Start(conn *Connection) *Service {

	worker := conn.InitializeWorkQueueSubscriber("work.transcode", handle)
	return &Service{
		worker: worker,
	}
}

func handle(message amqp.Delivery) {
	log.Printf("Got a message: %v", message.Body)
	time.Sleep(5 * time.Second)
	log.Println("Done.")
}

func (service *Service) Shutdown() {
	service.worker.Close()
}
