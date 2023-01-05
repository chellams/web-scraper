package crawl

import "golang.org/x/net/html"

type LinkExtractor struct {
}

func NewLinkExtractor() LinkExtractor {
	return LinkExtractor{}
}

func (l *LinkExtractor) extractLinksFromToken(token html.Token, baseUrl string) (string, bool) {
	for _, attribute := range token.Attr {
		if attribute.Key == "href" {
			link := attribute.Val
			s, err := formatUrl(link, baseUrl)
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
