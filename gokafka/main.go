package main

import (
	"github.com/dup2X/duplex/gokafka/server"
)

func main() {
	s := server.New()
	s.Serve()
	select {}
}

func initAll() {
}

func waitForSignal() {}
