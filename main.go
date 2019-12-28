package main

import (
	"github.com/alexweav/discordbot-transcoder/transcoder"
	"log"
	"os"
	"time"
)

func main() {
	os.Exit(run())
}

func run() int {
	conn, err := transcoder.Connect("localhost", 5672, "guest", "guest")
	if err != nil {
		log.Fatalf("Could not connect to RabbitMQ: %s", err)
		return -1
	}
	defer conn.Close()

	status := make(chan int)
	log.Println("Connected!")

	go func() {
		time.Sleep(5 * time.Second)
		status <- 0
	}()

	return <-status
}
