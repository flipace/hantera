package lib

import "flag"

// ParseArguments : parses command line flags
func ParseArguments() map[string]string {
	configFile := flag.String("config", "", "configuration file to use")
	targetDir := flag.String("target", "", "target directory for installation")

	flag.Parse()

	args := make(map[string]string)

	args["configFile"] = *configFile
	args["targetDir"] = *targetDir
	args["command"] = flag.Args()[0]

	return args
}
