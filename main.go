package main

import (
	"github.com/flipace/hantera/commands"
	"github.com/flipace/hantera/lib"
)

func main() {
	println("hantera -")

	// parse command line arguments
	arguments := lib.ParseArguments()

	// run the correct function for given cmd
	switch arguments["command"] {
	case "setup":
		commands.InitDev(arguments)
	case "update-dev":
		commands.UpdateDev(arguments)
	default:
		panic("No or unknown command " + arguments["command"] + " passed to hantera")
	}
}
