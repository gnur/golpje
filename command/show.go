package command

import (
	"flag"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/gnur/golpje/golpje"
	"github.com/spf13/viper"
)

// ShowCommand basic setup
type ShowCommand struct {
	Cfg *viper.Viper
}

// Help returns the help for this command
func (c *ShowCommand) Help() string {
	return "Super awesome help for this ShowCommand"
}

// Run actually runs the command
func (c *ShowCommand) Run(args []string) int {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	defer conn.Close()

	client := pb.NewGolpjeClient(conn)

	if len(args) == 0 {
		args = []string{"list", "-all"}
	}
	if args[0] == "add" {
		addCommand := flag.NewFlagSet("add", flag.ExitOnError)

		showName := addCommand.String("name", "none", "name of the show")
		showRegexp := addCommand.String("regexp", "none", "regexp to match episodes against")
		showSeasonal := addCommand.Bool("seasonal", true, "if show is seasonal (false for shows like the daily show)")
		showActive := addCommand.Bool("active", true, "show status")
		showMinimal := addCommand.Int64("minseason", 0, "Minimal season to download")

		addCommand.Parse(args[1:])

		req := pb.ProtoShow{
			Name:     *showName,
			Regexp:   *showRegexp,
			Seasonal: *showSeasonal,
			Active:   *showActive,
			Minimal:  *showMinimal,
		}

		id, err := client.AddShow(context.Background(), &req)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("added Show with id: ", id)
		}
	} else if args[0] == "del" {
		delCommand := flag.NewFlagSet("del", flag.ExitOnError)

		showID := delCommand.String("id", "none", "uuid of the show to delete")

		delCommand.Parse(args[1:])

		req := pb.ProtoShow{
			ID: *showID,
		}

		id, err := client.DelShow(context.Background(), &req)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("deleted Show with id: ", id)
		}

	} else if args[0] == "list" {
		listCommand := flag.NewFlagSet("list", flag.ExitOnError)

		returnActive := listCommand.Bool("active", false, "return only active shows")
		returnName := listCommand.String("name", "", "list only shows with this name")

		listCommand.Parse(args[1:])

		var req pb.ShowRequest
		req.Onlyactive = *returnActive
		req.Name = *returnName
		resp, err := client.GetShows(context.Background(), &req)

		if err == nil {
			for _, show := range resp.Shows {
				fmt.Println(show.ID, show.Name)
			}
		}
	} else if args[0] == "sync" {
		syncCommand := flag.NewFlagSet("sync", flag.ExitOnError)

		showUUID := syncCommand.String("uuid", "", "sync episodes of show with this id")

		syncCommand.Parse(args[1:])

		var req pb.SyncShowRequest
		req.ShowID = *showUUID
		resp, err := client.SyncShow(context.Background(), &req)

		if err == nil {
			fmt.Println(resp.FoundEpisodes)
			fmt.Println(resp.Success)
			fmt.Println(resp.Error)
		}
	}

	return 0
}

// Synopsis returns a show description
func (c *ShowCommand) Synopsis() string {
	return "This does shit"
}
