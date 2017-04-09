package main

import (
	"os"
	"time"

	"github.com/flipace/hantera/commands/develop"
	"github.com/urfave/cli"
)

var (
	// Version : current version of hantera
	Version string
	// Build : build date and time of hantera
	Build string
)

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = "hantera"
	app.Compiled = time.Now()

	app.Commands = append(
		[]cli.Command{},
		develop.Commands...,
	)

	app.Run(os.Args)
}
