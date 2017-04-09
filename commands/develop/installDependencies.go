package develop

import (
	"os"
	"os/exec"
	"path"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// InstallDependencies : installs project dependencies (tries to figure out package manager e.g. npm)
func InstallDependencies(c *cli.Context) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")

	lib.Catchy("Installing dependencies for \"%s\" v%s...\n", config.Name, config.Version)

	// create a semaphore with l
	var wg sync.WaitGroup

	for key := range config.Dependencies {
		wg.Add(1)

		go func(target string, key string) {

			targetDir := path.Join(target, key)

			if _, err := os.Stat(path.Join(targetDir, "package.json")); err == nil {
				lib.Notice("|-- Found package.json for %s - running 'yarn'\n", key)

				cmd := exec.Command("yarn", "install")
				cmd.Dir = targetDir
				cmd.Stdout = os.Stdout

				err = cmd.Start()
				check(err)

				err = cmd.Wait()
				check(err)
			}

			wg.Done()
		}(target, key)

	}

	wg.Wait()
}
