package main

import (
	"ack/screen"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "run" {
		screen.Setup()
	} else if command == "help" {
		PrintHelp()
	}
}
