package command

import (
	"fmt"

	"github.com/gnur/golpje/database"
	"github.com/gnur/golpje/structs"
)

// EventCommand basic setup
type EventCommand struct {
	Test string
}

// Help returns the help for this command
func (c *EventCommand) Help() string {
	return "Super awesome help for this EventCommand"
}

// Run actually runs the command
func (c *EventCommand) Run(args []string) int {
	db := database.Conn
	var events []structs.Event

	err := db.All(&events)
	if err != nil {
		fmt.Println("oops")
	}
	for _, e := range events {
		fmt.Println("------------------------")
		fmt.Println(e.ID)
		fmt.Println(e.Data)
		fmt.Println(e.Timestamp)
	}

	return 0
}

// Synopsis returns a show description
func (c *EventCommand) Synopsis() string {
	return "This does shit"
}
