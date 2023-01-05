package service

import (
	"context"
	"fmt"

	"github.com/chellams/web-scraper/api/v1/proto"
	crawl "github.com/chellams/web-scraper/internal/scrape"
)

type ScraperService struct {
	scraper *crawl.Scraper
}

func NewScraper(scraper *crawl.Scraper) proto.CrawlerServiceServer {
	return &ScraperService{
		scraper: scraper,
	}
}

func (s *ScraperService) GetURLs(_ context.Context, request *proto.ScraperRequest) (*proto.ScraperResponse, error) {

	baseUrl := request.GetUrl()
	urls, err := s.scraper.ScrapeUrl(baseUrl)

	if err != nil {
		fmt.Println(err)
	}

	return &proto.ScraperResponse{
		Urls: urls,
	}, nil
}
