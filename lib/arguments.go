package lib

import (
	"flag"
	"fmt"
	"os"
)

// ParseArguments : parses command line flags
func ParseArguments() map[string]string {
	config := flag.String("config", "", "configuration file to use")
	target := flag.String("target", "", "target directory for installation")
	help := flag.Bool("help", false, "lists available commands")

	flag.Parse()

	if *help {
		Notice("Available commands:\n")

		for key, value := range RootCommands {
			fmt.Printf("  %s - %s\n", key, value)
		}

		os.Exit(0)
	}

	args := make(map[string]string)

	args["config"] = *config
	args["target"] = *target

	commands := flag.Args()

	if len(commands) <= 1 {
		Warning("No or unkown command passed to hantera. Use hantera -help")
		os.Exit(0)
	}

	// set root and sub command
	args["rootCommand"] = commands[0]
	args["subCommand"] = commands[1]

	return args
}

// CheckRequiredArguments : checks whether all arguments in requiredArguments
// are available in arguments map and prints errors for missing arguments
func CheckRequiredArguments(requiredArguments []string, arguments map[string]string) {
	var errors []string

	for _, value := range requiredArguments {
		if arguments[value] == "" {
			errors = append(errors, fmt.Sprintf("--%s param is required", value))
		}
	}

	if len(errors) > 0 {
		for _, value := range errors {
			Error(value + "\n")
		}

		println()
		os.Exit(1)
	}
}
