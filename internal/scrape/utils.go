package crawl

import (
	"errors"
	"strings"
)

func formatUrl(url string, base string) (string, error) {

	switch {
	case strings.HasPrefix(url, "https://"):
	case strings.HasPrefix(url, "http://"):
		if strings.Contains(url, base) {
			return url, nil
		}
	case strings.HasPrefix(url, "/"):
		return base + url, nil
	}

	return "", errors.New("url doesn't contain any valid scheme")
}
