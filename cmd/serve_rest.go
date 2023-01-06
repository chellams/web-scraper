package main

import (
	"net/http"

	"github.com/rs/zerolog/log"

	crawl "github.com/chellams/web-scraper/internal/scrape"
)

type RestServer struct {
	scraper *crawl.Scraper
	address string
}

func NewRestServer(scraper *crawl.Scraper, address string) WebServer {
	return RestServer{
		scraper: scraper,
		address: address,
	}
}

func (r RestServer) Serve() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", r.GetUrls)

	err := http.ListenAndServe(r.address, mux)
	if err != nil {
		log.Err(err).Msg(" error in starting REST server")
	}

	log.Info().Msgf(" Listening on %s", r.address)
}

func (r RestServer) GetUrls(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	url := values.Get("url")

	results, err := r.scraper.ScrapeUrl(url)
	if err != nil {
		http.Error(writer, "error in scraping url", http.StatusInternalServerError)
		log.Err(err).Msg(" error in scraping url")
		return
	}

	var response []byte
	for _, s := range results {
		response = append(response, []byte(s+"\n")...)
	}
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Err(err).Msg(" error in writing bytes to response writer")
		return
	}
	log.Info().Msgf("URLs has been sent for url %s", url)
}
