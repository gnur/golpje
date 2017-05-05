package main

import (
	"github.com/gnur/golpje/command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all available golpje commands
var Commands map[string]cli.CommandFactory

func init() {

	Commands = map[string]cli.CommandFactory{
		"test": func() (cli.Command, error) {
			return &command.TestCommand{
				Test: "hoi",
			}, nil
		},
		"event": func() (cli.Command, error) {
			return &command.EventCommand{
				Test: "hoi",
			}, nil
		},
	}
}
