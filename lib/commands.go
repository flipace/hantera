package lib

// RootCommands : all available root commands
var RootCommands = map[string]string{
	"develop": "Run a command which targets development environments (such as init or update)",
}

// DevelopCommands : commands which target development environments
var DevelopCommands = map[string]string{
	"init":   "Initializes development environment for the given configuration file.",
	"update": "Updates development environment by pulling all dependencies for config",
}
