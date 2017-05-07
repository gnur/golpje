package events

import (
	"fmt"
	"strings"
	"time"

	"github.com/asdine/storm"
	"github.com/gnur/golpje/database"
	"github.com/google/uuid"
)

// Event an event consists of data and any related urns in Related
type Event struct {
	ID        uuid.UUID
	Timestamp time.Time `storm:"index"`
	Related   []string  `storm:"index"`
	Data      string
}

// All returns all events
func All() ([]Event, error) {
	db := database.Conn
	var events []Event

	err := db.AllByIndex("Timestamp", &events, storm.Reverse())
	return events, err
}

// After returns all events after the provided timestamp
func After(then time.Time) ([]Event, error) {
	events, err := All()
	if err != nil {
		return nil, err
	}
	var returnEvents []Event

	for _, e := range events {
		if e.Timestamp.After(then) {
			returnEvents = append(returnEvents, e)
		}
	}

	return returnEvents, nil
}

// New saves a new event to the database
func New(text string, related []string) (uuid.UUID, error) {
	u1 := uuid.New()

	e := Event{
		ID:        u1,
		Timestamp: time.Now(),
		Related:   related,
		Data:      text,
	}
	err := database.Conn.Save(&e)
	if err != nil {
		return [16]byte{}, err
	}
	return e.ID, nil
}

// Print provides a convenient way of pretty printing a event
func Print(e Event) {
	fmt.Println("--------------")
	fmt.Println(e.Timestamp.Format(time.Stamp), "  -  ", e.ID)
	fmt.Println("> ", e.Data)
	if len(e.Related) > 0 {
		fmt.Println("| Related: [", strings.Join(e.Related, "], ["), "]")
	}
}
