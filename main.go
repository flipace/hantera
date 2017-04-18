package main

import (
	"os"

	"github.com/flipace/hantera/commands"
)

var (
	// Version : current version of hantera
	Version string
	// Build : build date and time of hantera
	Build string
)

func main() {
	app := commands.GetApp(Version, Build)

	app.Run(os.Args)
}
