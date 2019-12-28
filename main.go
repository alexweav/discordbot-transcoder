package main

import (
	"fmt"
	"os"
	"github.com/alexweav/discordbot-transcoder/transcoder"
)

func main() {
	os.Exit(run())
}

func run() int {
	status := make(chan int)
	fmt.Println(transcoder.Hello())
	return <-status
}
