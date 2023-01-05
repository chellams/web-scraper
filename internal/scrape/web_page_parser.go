package crawl

import (
	"errors"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

type WebPageParser struct {
	baseUrl       string
	linkExtractor LinkExtractor
}

func NewWebPageParser(baseUrl string) *WebPageParser {
	return &WebPageParser{
		baseUrl:       baseUrl,
		linkExtractor: NewLinkExtractor(baseUrl),
	}
}

func (w *WebPageParser) Parse(response *http.Response, urlChannel chan string, wg *sync.WaitGroup, urlCache *sync.Map) {
	if response == nil {
		return
	}
	tokenizer := html.NewTokenizer(response.Body)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			wg.Done()
			log.Err(errors.New("error in getting token from tokenizer"))
			return
		}

		token := tokenizer.Token()

		if isLink(tokenType, token) {
			extractedURL, isExtracted := w.linkExtractor.extractLinksFromToken(token)
			if !isExtracted {
				log.Debug().Str("unable to extract the url ", token.String())
			}

			if len(extractedURL) != 0 {
				_, ok := urlCache.Load(extractedURL)
				if !ok {
					urlChannel <- extractedURL
					urlCache.Store(extractedURL, extractedURL)
				}
			}
		}
	}
}

func isLink(tokenType html.TokenType, token html.Token) bool {
	return tokenType == html.StartTagToken && token.DataAtom.String() == "a"
}
