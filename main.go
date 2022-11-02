package main

import (
	"ack/screen"
	"ack/util"
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
	argMap := ArgsToMap()

	if command == "run" {
		t := screen.NewTransfer()
		t.Run()
	} else if command == "credits" {
		c := screen.NewCredits()
		c.Run()
	} else if command == "transfer" {
		bufferSize := util.Atoi(argMap["buffer_size"], 10)
		t := screen.NewTransferWithOptions(bufferSize)
		t.Run()
	} else if command == "help" {
		PrintHelp()
	}
}
