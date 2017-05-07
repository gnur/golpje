package command

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/gnur/golpje/events"
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

	if len(args) == 0 {
		args = []string{"list", "since", "24h"}
	}
	if args[0] == "add" {
		addCommand := flag.NewFlagSet("add", flag.ExitOnError)

		eventText := addCommand.String("text", "dummy event", "text of the event to add")
		relatedTags := addCommand.String("related", "", "comma separated list of related arns")

		addCommand.Parse(args[1:])

		var related []string

		if *relatedTags != "" {
			related = strings.Split(*relatedTags, ",")
		}

		id, err := events.New(*eventText, related)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("added event with id: ", id)
		}

	} else if args[0] == "list" {
		listCommand := flag.NewFlagSet("list", flag.ExitOnError)

		returnAll := listCommand.Bool("all", false, "Return all events")
		returnSince := listCommand.Duration("since", 24*time.Hour, "period from which to return events")

		listCommand.Parse(args[1:])
		var allevents []events.Event

		if *returnAll {
			allevents, _ = events.All()
			fmt.Println("Retrieving all events")
		} else {
			now := time.Now()
			then := now.Add(-*returnSince)
			fmt.Println("Retrieving events since", then)
			allevents, _ = events.After(then)
		}
		for _, e := range allevents {
			e.Print()
		}
	}

	return 0
}

// Synopsis returns a show description
func (c *EventCommand) Synopsis() string {
	return "This does shit"
}
