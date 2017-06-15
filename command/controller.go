package command

import (
	"github.com/gnur/golpje/controller"
	"github.com/spf13/viper"
)

// ControllerCommand basic setup
type ControllerCommand struct {
	Cfg *viper.Viper
}

// Help returns the help for this command
func (c *ControllerCommand) Help() string {
	return "Super awesome help for this ControllerCommand"
}

// Run actually runs the command
func (c *ControllerCommand) Run(args []string) int {
	err := controller.Start(c.Cfg)
	if err != nil {
		return 1
	}
	return 0
}

// Synopsis returns a show description
func (c *ControllerCommand) Synopsis() string {
	return "This starts the main process that will search, download and provide the API for the other CLI commands"
}
