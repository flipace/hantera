package develop

import (
	"os"
	"path"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

func doInstallDependencies(target string, name string, wg *sync.WaitGroup) {
	if _, err := os.Stat(path.Join(target, "package.json")); err == nil {
		lib.Notice("--| Found package.json for %s - running 'yarn'\n", name)

		yarnOut, yarnErr := lib.Run(true, target, false, "yarn", "install")

		lib.Notice("%s: %s", name, yarnOut.String())

		if len(yarnErr.String()) > 0 {
			println(yarnErr.String())
		}
	} else {
		lib.Notice(">> Found no package.json for %s", name)
	}

	wg.Done()
}

// InstallDependencies : installs project dependencies (tries to figure out package manager e.g. npm)
func InstallDependencies(c *cli.Context, params ...string) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")

	lib.Catchy("\nInstalling dependencies for \"%s\" v%s...\n", config.Name, config.Version)

	// create a semaphore with l
	var wg sync.WaitGroup

	if len(params) == 2 {
		wg.Add(1)
		go doInstallDependencies(params[1], params[0], &wg)
	} else {
		for key := range config.Dependencies {
			wg.Add(1)

			targetDir := path.Join(target, key)

			go doInstallDependencies(targetDir, key, &wg)
		}
	}

	wg.Wait()
}
