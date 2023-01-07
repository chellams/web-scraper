package main

import (
	"github.com/chellams/web-scraper/internal/config"
	crawl "github.com/chellams/web-scraper/internal/scrape"
	"github.com/chellams/web-scraper/internal/service"
	"github.com/gosidekick/goconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	serverConfig := mustGetServerConfig()

	initLog(serverConfig.LogLevel)

	scraper := crawl.NewScraper()
	server := service.NewScraper(scraper)

	if serverConfig.IsGRPCEnabled {
		grpcServer := NewGServer(server, serverConfig.Address)
		grpcServer.Serve()
	} else {
		restServer := NewRestServer(scraper, serverConfig.Address)
		restServer.Serve()
	}
}

func initLog(logLevel string) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if logLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func mustGetServerConfig() *config.ServerConfig {
	serverConfig := &config.ServerConfig{}
	err := goconfig.Parse(serverConfig)
	if err != nil {
		log.Err(err).Msg("error in initialising server configuration")
		return nil
	}

	return serverConfig
}
