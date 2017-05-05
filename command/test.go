package command

import "fmt"

// TestCommand basic setup
type TestCommand struct {
	Test string
}

// Help returns the help for this command
func (c *TestCommand) Help() string {
	fmt.Println("Help!")
	return ""
}

// Run actually runs the command
func (c *TestCommand) Run(args []string) int {
	restr := ""
	for _, arg := range args {
		restr = restr + " " + arg
	}
	fmt.Println(restr)

	return 0
}

// Synopsis returns a show description
func (c *TestCommand) Synopsis() string {
	return "This does shit"
}
