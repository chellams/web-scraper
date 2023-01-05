package main

import (
	"net/http"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	crawl "github.com/chellams/web-scraper/internal/scrape"
)

func main() {

	initLog()

	baseUrl := "https://www.youtube.com"
	webPageExtractor := crawl.NewWebPageExtractor(baseUrl)
	webPageParser := crawl.NewWebPageParser(baseUrl)

	urlChannel := make(chan string)
	responseChannel := make(chan *http.Response)
	errorChannel := make(chan error)
	urlCache := &sync.Map{}

	var result []string

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		urlChannel <- baseUrl
		wg.Done()
	}()

	go func() {
		for {
			select {
			case url := <-urlChannel:
				wg.Add(1)
				result = append(result, url)
				go webPageExtractor.ExtractWebPage(url, responseChannel, errorChannel)
			case pageResponse := <-responseChannel:
				go webPageParser.Parse(pageResponse, urlChannel, wg, urlCache)
			case err := <-errorChannel:
				log.Err(err).Msg(" error received from channel")
			default:
				// fmt.Println(" default")
			}
		}
	}()

	wg.Wait()
}

func initLog() {
	logger := zerolog.Logger{}
	level := zerolog.DebugLevel
	logger.Level(level)
}
