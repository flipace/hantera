package develop

import (
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// InstallDependencies : installs project dependencies (tries to figure out package manager e.g. npm)
func InstallDependencies(c *cli.Context) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")
	progress := c.Bool("progress")

	lib.Catchy("Installing dependencies for \"%s\" v%s...\n", config.Name, config.Version)

	// create a semaphore with l
	var wg sync.WaitGroup

	for key := range config.Dependencies {
		wg.Add(1)

		go func(target string, key string) {

			targetDir, err := filepath.Abs(path.Join(target, key))
			check(err)

			if _, err := os.Stat(path.Join(targetDir, "package.json")); err == nil {
				lib.Notice("|-- Found package.json for %s - running 'yarn'\n", key)

				lib.Run(progress, targetDir, "yarn", "install")
			}

			wg.Done()
		}(target, key)

	}

	wg.Wait()
}
