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
		grpcServer := NewGServer(server, "localhost:9876")
		grpcServer.Serve()
	} else {
		restServer := NewRestServer(scraper, "localhost:9876")
		restServer.Serve()
	}
}

func initLog() {
	logger := zerolog.Logger{}
	level := zerolog.DebugLevel
	logger.Level(level)
}
