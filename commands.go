package main

import (
	"github.com/gnur/golpje/command"
	"github.com/gnur/golpje/config"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all available golpje commands
var Commands map[string]cli.CommandFactory

func init() {

	cfg, err := config.Load()
	if err != nil {
		return
	}

	Commands = map[string]cli.CommandFactory{
		"event": func() (cli.Command, error) {
			return &command.EventCommand{
				Cfg: cfg,
			}, nil
		},
		"start": func() (cli.Command, error) {
			return &command.ControllerCommand{
				Cfg: cfg,
			}, nil
		},
		"show": func() (cli.Command, error) {
			return &command.ShowCommand{
				Cfg: cfg,
			}, nil
		},
	}
}
