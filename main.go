// Contains entry points and provides service initialization logic.
package main

import (
	"github.com/alexweav/discordbot-transcoder/transcoder"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Entry point for the service.
func main() {
	status := run()
	log.Printf("Shutting down with status code %d.\n", status)
	os.Exit(status)
}

// Connects to RabbitMQ and initializes the service.
func run() int {
	shutdown := make(chan int)

	conn, err := transcoder.Connect("localhost", 5672, "guest", "guest")
	if err != nil {
		log.Fatalf("Could not connect to RabbitMQ: %s", err)
		return -1
	}
	defer conn.Close()

	log.Println("Connected!")

	worker := transcoder.InitializeWorkQueueSubscriber(conn, "test.queue")
	defer worker.Close()

	go catchSignals(shutdown)
	return <-shutdown
}

// Catches shutdown signals so resources can be cleaned up.
func catchSignals(shutdown chan int) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signals
	log.Printf("Received exit signal: %v", sig)
	shutdown <- 0
}
