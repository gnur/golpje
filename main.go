package main

import (
	"fmt"
	"os"

	"github.com/asdine/storm"
	"github.com/gnur/golpje/database"
	"github.com/mitchellh/cli"
)

func main() {
	database.Conn, _ = storm.Open("test.db")
	defer database.Conn.Close()

	c := cli.NewCLI("app", "1.0")
	c.Args = os.Args[1:]

	c.Commands = Commands
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(os.Stderr, err.Error())
	}
	os.Exit(exitStatus)
}
