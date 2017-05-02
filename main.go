package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("app", "1.0")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"foo": func() (cli.Command, error) {
			return &TestCommand{
				Test: "hoi",
			}, nil
		},
	}
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(os.Stderr, err.Error())
	}

	os.Exit(exitStatus)

}
