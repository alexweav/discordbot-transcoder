package transcoder

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Connection struct {
	Connection *amqp.Connection
}

func Connect(address string, port int, user string, pass string) (*Connection, error) {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d", user, pass, address, port)
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("Could not connect to RabbitMQ: %s", err)
		return nil, err
	}
	return &Connection{
		Connection: conn,
	}, nil
}

func (connection *Connection) Close() {
	connection.Connection.Close()
}
