package search

import (
	torrentapi "github.com/qopher/go-torrentapi"
	log "github.com/sirupsen/logrus"
)

// Result holds all relevant info
type Result struct {
	Title      string
	Seeders    int
	MagnetLink string
}

// Engine is the interface
type Engine interface {
	Search(string) []Result
}

// New returns a new search engine
func New(engine string) *Engine {
	var a Engine
	return &a
}

// TorrentAPI searches the torrent api
func TorrentAPI(query string) ([]Result, error) {
	results := []Result{}
	api, err := torrentapi.New("golpje")
	if err != nil {
		log.WithField("error", err).Warning("creating api failed")
		return results, err
	}
	api.SearchString(query)
	api.Ranked(true).Sort("seeders").Format("json_extended").Limit(50)

	res, err := api.Search()
	if err != nil {
		log.WithField("error", err).Warning("searching failed")
		return results, err
	}
	for _, r := range res {
		results = append(results, Result{
			Title:      r.Title,
			Seeders:    r.Seeders,
			MagnetLink: r.Download,
		})
	}

	return results, nil
}
