package transcoder

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net"
	"time"
)

// The base timeout for AMQP connections.
const amqpConnectionTimeout = 30

// The number of connection attempts to make before exiting.
const numRetries = 5

// Represents a channel and connection.
type Connection struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// Establishes a new connection to a RabbitMQ instance.
func Connect(address string, port int, user, pass string) (*Connection, error) {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d", user, pass, address, port)
	conn, err := establishConnection(uri, amqpConnectionTimeout*time.Second)
	if err != nil {
		log.Fatalf("Could not connect to RabbitMQ: %s", err)
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Could not open main channel: %s", err)
		return nil, err
	}

	return &Connection{
		Connection: conn,
		Channel:    channel,
	}, nil
}

// Closes a connection.
func (connection *Connection) Close() {
	log.Println("Closing connection to RabbitMQ.")
	connection.Channel.Close()
	connection.Connection.Close()
}

// Establishes a connection to RabbitMQ.
func establishConnection(uri string, timeout time.Duration) (*amqp.Connection, error) {
	log.Println("Attempting to connect to RabbitMQ...")

	backoff := NewConstantBackoff(ConstantBackoffPolicy{
		MaxElapsedTime: 60 * time.Second,
	})

	for interval := backoff.Next(); interval != Stop; interval = backoff.Next() {
		if conn, err := dialWithTimeout(uri, timeout); err == nil {
			return conn, err
		} else {
			log.Printf("Connection failed (%v). Retrying in 10 seconds...", err)
			time.Sleep(interval)
		}
	}
	return dialWithTimeout(uri, timeout)
}

// Connects to a RabbitMQ instance at the given URI with a custom timeout.
func dialWithTimeout(uri string, timeout time.Duration) (*amqp.Connection, error) {
	config := amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
	}
	return amqp.DialConfig(uri, config)
}
