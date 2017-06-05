package searcher

import (
	"fmt"
	"time"

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
func Start(results chan Searchresult) {

	for {
		shows, err := shows.All()
		if err != nil {
			continue
		}
		for _, show := range shows {
			events.New(fmt.Sprintf("Searching for new episodes of %s", show.Name), []string{show.ID})
			_, torrents := gopiratebay.Search(show.Name)
			for _, torrent := range torrents {
				results <- Searchresult{
					Title:      torrent.Title,
					Magnetlink: torrent.Magnetlink,
					Vipuser:    torrent.Vipuser,
					Seeders:    torrent.Seeders,
					ShowID:     show.ID,
				}
			}
		}
		time.Sleep(5 * time.Minute)
	}
}
