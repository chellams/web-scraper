package main

import (
	"github.com/chellams/web-scraper/api/v1/proto"
	crawl "github.com/chellams/web-scraper/internal/scrape"
	"github.com/chellams/web-scraper/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GServer struct {
}

func NewGServer() WebServer {
	return GServer{}
}

func (g GServer) Serve(address string) {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := service.NewScraper(crawl.NewScraper())

	gServer := grpc.NewServer()
	proto.RegisterCrawlerServiceServer(gServer, server)

	reflection.Register(gServer)

	if err := gServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
