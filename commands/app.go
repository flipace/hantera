package commands

import (
	"time"

	"github.com/flipace/hantera/commands/develop"
	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// GetApp : returns the command line application
func GetApp(version string, build string) *cli.App {
	app := cli.NewApp()

	app.Version = version
	app.Name = "hantera"
	app.Usage = "a tool which helps you manage projects which utilize a service oriented architecture."
	app.Compiled = time.Now()

	lib.Catchy("%s %s\n", app.Name, version)

	app.Commands = append(
		[]cli.Command{},
		develop.Commands...,
	)

	return app
}
