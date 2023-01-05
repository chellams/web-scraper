package crawl

import "golang.org/x/net/html"

type LinkExtractor struct {
	baseUrl string
}

func NewLinkExtractor(baseUrl string) LinkExtractor {
	return LinkExtractor{
		baseUrl: baseUrl,
	}
}

func (l *LinkExtractor) extractLinksFromToken(token html.Token) (string, bool) {
	for _, attribute := range token.Attr {
		if attribute.Key == "href" {
			link := attribute.Val
			s, err := formatUrl(link, l.baseUrl)
			if err != nil {
				// log.Err(err).Msg(" error in formatting URLS")
			}
			if s == "" {
				break
			}
			return s, true
		}
	}
	return "", false
}
