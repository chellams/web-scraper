package crawl

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type WebPageExtractor struct {
	url string
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:57.0) Gecko/20100101 Firefox/57.0",
	"Mozilla/5.0 (Linux; Android 7.0; SM-G930V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_4 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B350 Safari/8536.25",
	"Mozilla/5.0 (Linux; U; Android 5.1; locale; device Build/build) AppleWebKit/webkit (KHTML, like Gecko) Version/4.0 Chrome/chrome Safari/safari",
}

func NewWebPageExtractor(url string) *WebPageExtractor {
	return &WebPageExtractor{
		url: url,
	}
}

func (w *WebPageExtractor) ExtractWebPage(webPageUrl string, webPageResponse chan *http.Response, errorChannel chan error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	request, err := http.NewRequest("GET", webPageUrl, nil)
	if err != nil {
		log.Err(err).Msg("error in creating request")
		errorChannel <- err
	}
	request.Header.Set("User-Agent", userAgents[0])
	response, err := client.Do(request)
	if err != nil {
		log.Err(err).Msg("error in getting response from ")
		errorChannel <- err
	}

	webPageResponse <- response
}
