package searcher

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
	"github.com/gnur/go-piratebay"
	"github.com/gnur/golpje/events"
	"github.com/gnur/golpje/shows"
)

// Searchresult contains everything needed to make an informed download decision
type Searchresult struct {
	Title      string
	Magnetlink string
	Vipuser    bool
	Seeders    int
	ShowID     string
}

// Start starts the searcher that periodically searches for shows
func Start(db *storm.DB, results chan Searchresult, searchInterval time.Duration) {

	for {
		shows, err := shows.All(db)
		if err != nil {
			continue
		}
		for _, show := range shows {
			events.New(db, fmt.Sprintf("Searching for new episodes of %s", show.Name), []string{show.ID})
			torrents, _ := piratebay.Search(show.Name)
			for _, torrent := range torrents {
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
