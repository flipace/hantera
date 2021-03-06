package develop

import (
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// RunUpdate : runs the update procedure
func RunUpdate(configFile string, target string, branch string, nodeps bool, c *cli.Context) {
	config := lib.GetProductConfig(configFile)
	steps := config.Steps.Update

	workingDir, _ := filepath.Abs(target)

	if len(steps.Pre) > 0 {
		lib.Catchy("Running 'update:pre' commands...\n")

		lib.ExecuteStep(steps.Pre, workingDir)
	}

	lib.Catchy("Updating dependencies for \"%s\" v%s...\n", config.Name, config.Version)

	var wg sync.WaitGroup

	if len(steps.Override) > 0 {
		lib.Warning("Running 'update:override' commands.. (CAUTION: This overrides hanteras core behavior!).\n")

		lib.ExecuteStep(steps.Override, workingDir)
	} else {
		for key, value := range config.Dependencies {
			wg.Add(1)

			targetDir, err := filepath.Abs(path.Join(target, key))
			check(err)

			// seems like some new dependency has been added, let's clone it
			if _, err := os.Stat(targetDir); err != nil {
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
	}

	wg.Wait()

	// install/updated all dependencies for the dependencies
	if nodeps == false {
		InstallDependencies(c)
	}

	if len(steps.Post) > 0 {
		lib.Catchy("\nRunning 'update:post' commands...\n")

		lib.ExecuteStep(steps.Post, workingDir)
	}

	lib.Catchy("\nDone updating %s!\n", config.Name)
}

// Update : updates each dependencies working copy to latest version from head
func Update(c *cli.Context) {
	configFile := c.String("config")
	branch := c.String("branch")

	target := c.String("target")
	nodeps := c.Bool("no-deps")

	RunUpdate(configFile, target, branch, nodeps, c)
}
