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
	command, err := ToTranscodeCommand(message.Body)
	if err != nil {
		log.Fatalf("Unparseable message: %v", message.Body)
		return
	}

	log.Printf("Got transcode job for file: %v", command.FileName)
	time.Sleep(5 * time.Second)
	log.Println("Done.")
}

func (service *Service) Shutdown() {
	service.worker.Close()
}
