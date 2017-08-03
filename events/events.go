package events

import (
	"fmt"
	"strings"
	"time"

	"github.com/asdine/storm"
	"github.com/gnur/golpje/golpje"
	"github.com/google/uuid"
)

// Event is a local Alias of the protobuf Event
type Event struct {
	ID        string
	Timestamp int64    `storm:"index"`
	Related   []string `storm:"index"`
	Data      string
}

// All returns all events
func All(db *storm.DB) ([]Event, error) {
	var events []Event

	err := db.AllByIndex("Timestamp", &events, storm.Reverse())
	if err != nil {
		fmt.Println(err.Error())
	}
	return events, err
}

// After returns all events after the provided timestamp
func After(db *storm.DB, then time.Time) ([]Event, error) {
	var events []Event
	err := db.Range("Timestamp", then.UnixNano(), time.Now().UnixNano(), &events)
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
func New(db *storm.DB, text string, related []string) (string, error) {
	u1 := uuid.New()

	e := Event{
		ID:        u1.String(),
		Timestamp: time.Now().UnixNano(),
		Related:   related,
		Data:      text,
	}
	err := db.Save(&e)
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

// FromProto translates a protoEvent to an Event that is used internally
func FromProto(in *golpje.ProtoEvent) Event {
	return Event{
		ID:        in.ID,
		Timestamp: in.Timestamp,
		Related:   in.Related,
		Data:      in.Data,
	}
}

// ToProto converts an Event to a ProtoEvent
func (e Event) ToProto() *golpje.ProtoEvent {
	eProto := golpje.ProtoEvent{
		ID:        e.ID,
		Timestamp: e.Timestamp,
		Related:   e.Related,
		Data:      e.Data,
	}
	return &eProto
}
