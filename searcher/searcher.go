package searcher

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
	"github.com/gnur/go-piratebay"
	"github.com/gnur/golpje/events"
	"github.com/gnur/golpje/shows"
	"github.com/prometheus/client_golang/prometheus"
)

// Searchresult contains everything needed to make an informed download decision
type Searchresult struct {
	Title      string
	Magnetlink string
	Vipuser    bool
	Seeders    int
	ShowID     string
}

// Searchmetrics contains everything related to metrics
type Searchmetrics struct {
	Enabled        bool
	Searches       prometheus.Counter
	FailedSearches prometheus.Counter
	SearchResults  prometheus.Counter
}

// Start starts the searcher that periodically searches for shows
func Start(db *storm.DB, piratebayURL string, results chan Searchresult, searchInterval time.Duration, m Searchmetrics) {
	pb := piratebay.New(piratebayURL)

	for {
		shows, err := shows.All(db)
		if err != nil {
			continue
		}
		for _, show := range shows {
			if m.Enabled {
				m.Searches.Inc()
			}
			events.New(db, fmt.Sprintf("Searching for new episodes of %s", show.Name), []string{show.ID})
			torrents, err := pb.Search(show.Name)
			if err != nil {
				events.New(db, fmt.Sprintf("Search failed for %s, %s", show.Name, err.Error()), []string{show.ID})
				if m.Enabled {
					m.FailedSearches.Inc()
				}
				continue
			}
			for _, torrent := range torrents {
				if m.Enabled {
					m.SearchResults.Inc()
				}
				results <- Searchresult{
					Title:      torrent.Title,
					Magnetlink: torrent.MagnetLink,
					Vipuser:    torrent.VIP,
					Seeders:    torrent.Seeders,
					ShowID:     show.ID,
				}
			}
		}
		time.Sleep(searchInterval)
	}
}
