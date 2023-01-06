package main

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/chellams/web-scraper/api/v1/proto"
)

type GRPCServer struct {
	scraperService proto.CrawlerServiceServer
	address        string
}

func NewGServer(scraper proto.CrawlerServiceServer, address string) WebServer {
	return GRPCServer{
		scraperService: scraper,
		address:        address,
	}
}

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tcp", g.address)

	if err != nil {
		log.Error().Err(err).Msg("error in listening gRPC server")
		return
	}

	gServer := grpc.NewServer()
	proto.RegisterCrawlerServiceServer(gServer, g.scraperService)

	reflection.Register(gServer)

	if err := gServer.Serve(lis); err != nil {
		log.Error().Err(err).Msg("error in starting gRPC server")
		return
	}

	log.Info().Msgf("listening on %s", g.address)
}
