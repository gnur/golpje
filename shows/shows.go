package shows

import (
	"errors"

	"github.com/asdine/storm"
	"github.com/gnur/golpje/database"
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

// New creates a new show
func New(name, regexp, episodeidtype string, active bool, minimal int) (uuid.UUID, error) {
	db := database.Conn

	var match Show

	err := db.One("Name", name, &match)
	if err != storm.ErrNotFound {
		return match.ID, errors.New("Show with this name already exists")
	}

	u1 := uuid.New()

	s := Show{
		ID:            u1,
		Name:          name,
		Regexp:        regexp,
		Episodeidtype: episodeidtype,
		Active:        active,
		Minimal:       minimal,
	}
	err = database.Conn.Save(&s)
	return s.ID, nil
}
