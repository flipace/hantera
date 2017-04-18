package develop

import (
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// Update : updates each dependencies working copy to latest version from head
func Update(c *cli.Context) {
	configFile := c.String("config")
	branch := c.String("branch")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")
	nodeps := c.Bool("no-deps")

	lib.Catchy("Updating dependencies for \"%s\" v%s...\n", config.Name, config.Version)

	var wg sync.WaitGroup

	for key, value := range config.Dependencies {
		wg.Add(1)

		targetDir, err := filepath.Abs(path.Join(target, key))
		check(err)

		// seems like some new dependency has been added, let's clone it
		if _, err := os.Stat(targetDir); err != nil {
			wg.Add(1)
			go CloneRepository(key, targetDir, branch, value.Repository, false, &wg)
		} else {
			lib.Notice("--| Updating %s\n", key)

			// walk through all
			go func(name string, workingDir string) {
				pullOut, pullErr := lib.Run(true, targetDir, false, "git", "pull", "-r")

				lib.Notice("--| %s: %s", name, pullOut.String())
				if len(pullErr.String()) > 0 {
					println(pullErr.String())
				}

				wg.Done()
			}(key, targetDir)
		}
	}

	wg.Wait()

	// install/updated all dependencies for the dependencies
	if nodeps == false {
		InstallDependencies(c)
	}
}
