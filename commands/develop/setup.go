package develop

import (
	"path"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// Setup : sets up development environment for a project
func Setup(c *cli.Context) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")
	if target == "" {
		target = path.Join("./", config.Name)
	}

	progress := c.Bool("progress")

	lib.Catchy("Setup \"%s\" v%s\n", config.Name, config.Version)
	lib.Notice("\nCloning dependencies...\n")

	// the ref to checkout
	refName := c.String("branch")

	for key, value := range config.Dependencies {
		targetDir := path.Join(target, key)

		lib.Notice("--| Cloning %s to %s\n", key, targetDir)

		lib.Clone(value.Repository, targetDir, refName, progress)
	}

	InstallDependencies(c)
}
