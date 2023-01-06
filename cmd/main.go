package main

import (
	crawl "github.com/chellams/web-scraper/internal/scrape"
	"github.com/chellams/web-scraper/internal/service"
	"github.com/rs/zerolog"
)

func main() {

	initLog()

	var isGRPCEnabled = false

	scraper := crawl.NewScraper()
	server := service.NewScraper(scraper)

	if isGRPCEnabled {
		grpcServer := NewGServer(server)
		grpcServer.Serve("localhost:9876")
	} else {
		restServer := NewRestServer(scraper)
		restServer.Serve("localhost:9876")
	}
}

func initLog() {
	logger := zerolog.Logger{}
	level := zerolog.DebugLevel
	logger.Level(level)
}
