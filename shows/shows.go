package shows

import (
	"errors"
	"fmt"

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
	var match Show

	err := database.Conn.One("Name", name, &match)
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

// All returns all shows
func All() ([]Show, error) {
	var matches []Show
	err := database.Conn.All(&matches)
	return matches, err
}

// GetFromID retrieve a show from an UUID
func GetFromID(uuid uuid.UUID) (Show, error) {

	var match Show

	err := database.Conn.One("ID", uuid, &match)
	if err != nil {
		return match, err
	}
	return match, nil

}

// GetFromName retrieve a show from an UUID
func GetFromName(name string) (Show, error) {

	var match Show

	err := database.Conn.One("name", name, &match)
	if err != nil {
		return match, err
	}
	return match, nil

}

// Print provides a convenient way of pretty printing a show
func (s Show) Print() {
	fmt.Println("--------------")
	fmt.Println(s.ID, s.Name, s.Active)
}
