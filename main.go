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

	app.Commands = []cli.Command{
		{
			Name:      "develop",
			Aliases:   []string{"d"},
			Usage:     "run a task for develop environments",
			ArgsUsage: " ",
			Subcommands: []cli.Command{
				{
					Name:        "setup",
					Usage:       "hantera develop setup ~/path/to/manifest.yml",
					UsageText:   "hantera develop setup --config ./manifest.yml --target ./project",
					Description: "sets up a project for development",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "config, c, manifest",
							Usage: "load manifest/config from `FILE.yml`",
							Value: "./manifest.yml",
						},
						cli.StringFlag{
							Name:  "target, t",
							Usage: "setup project in `PATH`",
						},
						cli.StringFlag{
							Name:  "branch, b",
							Usage: "branch to checkout",
							Value: "refs/heads/develop",
						},
						cli.BoolFlag{
							Name:  "progress, p",
							Usage: "show clone progress",
						},
					},
					Action: develop.Setup,
				},
			},
		},
	}

	app.Run(os.Args)
	/*
		// parse command line arguments
		arguments := lib.ParseArguments()

		lib.Notice("[Command] %s -> %s\n\n", arguments["rootCommand"], arguments["subCommand"])

		// run the correct function for given cmd
		switch arguments["rootCommand"] {
		case "develop":
			develop.Main(arguments)
		default:
			panic("No or unknown command " + arguments["command"] + " passed to hantera")
		}*/
}
