package episodes

import (
	"time"

	"github.com/google/uuid"
)

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
