package main

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
	"github.com/google/uuid"
)

type Show struct {
	ID            uuid.UUID
	Name          string `storm:"unique"`
	Regexp        string
	Active        bool `storm:"index"`
	Episodeidtype string
	Minimal       int
}

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

type Event struct {
	ID        uuid.UUID
	Timestamp time.Time `storm:"index"`
	Related   []string  `storm:"index"`
	Data      string
}

func main() {
	db, err := storm.Open("test.db")
	defer db.Close()
	u1, _ := uuid.NewRandom()
	fmt.Println(u1)
	if err != nil {
		fmt.Println("nu al een faal", err.Error())
	}
	var events []Event
	err = db.All(&events)
	now := time.Now()
	then := now.Add(-2 * time.Hour)
	for _, e := range events {
		if e.Timestamp.After(then) {
			fmt.Println("------------------------")
			fmt.Println(e.ID)
			fmt.Println(e.Data)
			fmt.Println(e.Timestamp)
		}
	}
	fmt.Println("jep")
}
