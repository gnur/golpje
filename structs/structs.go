package structs

import (
	"time"

	"github.com/google/uuid"
)

// Show a show
type Show struct {
	ID            uuid.UUID
	Name          string `storm:"unique"`
	Regexp        string
	Active        bool `storm:"index"`
	Episodeidtype string
	Minimal       int
}

// Episode : The basic struct for all episodes
type Episode struct {
	ID          uuid.UUID
	Title       string
	Showid      uuid.UUID `storm:"index"`
	Added       time.Time `storm:"index"`
	Magnetlink  string    `storm:"unique"`
	Episodeid   string
	Downloaded  bool `storm:"index"`
	Downloading bool `storm:"index"`
}

// Event an event consists of data and any related urns in Related
type Event struct {
	ID        uuid.UUID
	Timestamp time.Time `storm:"index"`
	Related   []string  `storm:"index"`
	Data      string
}
