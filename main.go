package main

import (
	"os"
	"time"

	"github.com/flipace/hantera/commands/develop"
	"github.com/flipace/hantera/lib"
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
	app.Usage = "a tool which helps you manage projects which utilize a service oriented architecture."
	app.Compiled = time.Now()

	lib.Catchy("%s v%s (%s)\n", app.Name, Version, Build)

	app.Commands = append(
		[]cli.Command{},
		develop.Commands...,
	)

	app.Run(os.Args)
}
