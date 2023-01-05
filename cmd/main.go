package main

import (
	"github.com/rs/zerolog"
)

func main() {

	initLog()

	var isGRPCEnabled = true

	if isGRPCEnabled {
		grpcServer := NewGServer()
		grpcServer.Serve("localhost:9876")
	}
}

func initLog() {
	logger := zerolog.Logger{}
	level := zerolog.DebugLevel
	logger.Level(level)
}
