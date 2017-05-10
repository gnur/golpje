package events

import (
	"fmt"
	"strings"
	"time"

	"github.com/asdine/storm"
	"github.com/gnur/golpje/database"
	"github.com/gnur/golpje/golpje"
	"github.com/google/uuid"
)

// Event is a local Alias of the protobuf Event
type Event golpje.Event

// All returns all events
func All() ([]Event, error) {
	var events []Event

	err := database.Conn.AllByIndex("Timestamp", &events, storm.Reverse())
	if err != nil {
		fmt.Println(err.Error())
	}
	return events, err
}

// After returns all events after the provided timestamp
func After(then time.Time) ([]Event, error) {
	var events []Event
	err := database.Conn.Range("Timestamp", then.UnixNano(), time.Now().UnixNano(), &events)
	if err != nil {
		return nil, err
	}
	var returnEvents []Event

	for _, e := range events {
		if e.Time().After(then) {
			returnEvents = append(returnEvents, e)
		}
	}

	return returnEvents, nil
}

// New saves a new event to the database
func New(text string, related []string) (string, error) {
	u1 := uuid.New()

	e := Event{
		ID:        u1.String(),
		Timestamp: time.Now().UnixNano(),
		Related:   related,
		Data:      text,
	}
	err := database.Conn.Save(&e)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return e.ID, nil
}

// Time returns the added date as a time.Time stamp
func (e Event) Time() time.Time {
	return time.Unix(0, e.Timestamp)
}

// Print provides a convenient way of pretty printing an event
func (e Event) Print() {
	fmt.Println("--------------")
	fmt.Println(e.Time().Format(time.Stamp), "  -  ", e.ID)
	fmt.Println("> ", e.Data)
	if len(e.Related) > 0 {
		fmt.Println("| Related: [", strings.Join(e.Related, "], ["), "]")
	}
}
