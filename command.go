package main

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
func (c *TestCommand) Run(a []string) int {
	fmt.Println("Run!")
	return 0
}

// Synopsis returns a show description
func (c *TestCommand) Synopsis() string {
	fmt.Println("Synopsis!")
	return "This does shit"
}
