package crawl

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
)

type Scraper struct {
	webPageParser    *WebPageParser
	webPageExtractor *WebPageExtractor
}

func NewScraper() *Scraper {
	return &Scraper{
		webPageParser:    NewWebPageParser(NewLinkExtractor()),
		webPageExtractor: NewWebPageExtractor(),
	}
}

func (s *Scraper) ScrapeUrl(baseUrl string) ([]string, error) {
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
				go s.webPageExtractor.ExtractWebPage(url, responseChannel, errorChannel)
			case pageResponse := <-responseChannel:
				go s.webPageParser.Parse(pageResponse, urlChannel, wg, urlCache, baseUrl)
			case err := <-errorChannel:
				log.Err(err).Msg(" error received from channel")
			default:
				// fmt.Println(" default")
			}
		}
	}()

	wg.Wait()

	return result, nil
}
