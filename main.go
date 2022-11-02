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
		wireSpeed := util.Atoi(argMap["wire_speed"], 944)
		processSpeed := util.Atoi(argMap["process_speed"], 3944)
		t := screen.NewTransferWithOptions(bufferSize, wireSpeed, processSpeed)
		t.Run()
	} else if command == "help" {
		PrintHelp()
	}
}
