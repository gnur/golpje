package command

import (
	"fmt"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/gnur/golpje/events"
	pb "github.com/gnur/golpje/golpje"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// EventCommand basic setup
type EventCommand struct {
	Cfg *viper.Viper
}

// Help returns the help for this command
func (c *EventCommand) Help() string {
	return "Super awesome help for this EventCommand"
}

// Run actually runs the command
func (c *EventCommand) Run(args []string) int {
	conn, err := grpc.Dial(c.Cfg.GetString("cli_address"), grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	defer conn.Close()

	client := pb.NewGolpjeClient(conn)

	if len(args) == 0 {
		args = []string{"list", "-since", "24h"}
	}
	if args[0] == "list" {
		listCommand := flag.NewFlagSet("list", flag.ExitOnError)

		returnAll := listCommand.Bool("all", false, "Return all events")
		returnSince := listCommand.Duration("since", 24*time.Hour, "period from which to return events")

		listCommand.Parse(args[1:])
		var allevents *pb.ProtoEvents

		req := pb.EventRequest{}
		if *returnAll {
			req.All = true
			allevents, err = client.GetEvents(context.Background(), &req)
			if err != nil {
				fmt.Println(err.Error())
				return 1
			}
			fmt.Println("Retrieving all events")
		} else {
			now := time.Now()
			then := now.Add(-*returnSince)
			req.All = false
			req.Since = then.UnixNano()
			fmt.Println("Retrieving events since", then)
			allevents, err = client.GetEvents(context.Background(), &req)
			if err != nil {
				fmt.Println(err.Error())
				return 1
			}
		}

		for _, e := range allevents.Events {
			events.FromProto(e).Print()
		}
	}
	return 0
}

// Synopsis returns a show description
func (c *EventCommand) Synopsis() string {
	return "This does shit"
}
