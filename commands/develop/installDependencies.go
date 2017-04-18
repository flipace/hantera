package develop

import (
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

func doInstallDependencies(target string, name string, wg *sync.WaitGroup) {
	if _, err := os.Stat(path.Join(target, "package.json")); err == nil {
		lib.Notice("--| Found package.json for %s - running 'yarn'\n", name)

		yarnOut, yarnErr := lib.Run(true, target, false, "yarn -s")

		o := yarnOut.String()
		e := yarnErr.String()

		if len(o) > 0 {
			lib.Notice("--| %s: %s\n", name, o)
		}

		if len(e) > 0 {
			println(e)
		}
	} else {
		lib.Notice(">> Found no package.json for %s", name)
	}

	wg.Done()
}

// CmdInstallDependencies : this is a wrapper for InstallDependencies, used by the main command triggered by cli
func CmdInstallDependencies(c *cli.Context) {
	InstallDependencies(c)

	lib.Catchy("\nDone installing dependencies!\n")
}

// InstallDependencies : installs project dependencies (tries to figure out package manager e.g. npm)
func InstallDependencies(c *cli.Context, params ...string) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")
	steps := config.Steps.Dependencies

	workingDir, _ := filepath.Abs(target)

	lib.Catchy("\nInstalling dependencies for \"%s\" v%s...\n", config.Name, config.Version)

	var wg sync.WaitGroup

	if len(steps.Pre) > 0 {
		lib.Catchy("Running 'dependencies:pre' commands...\n")

		lib.ExecuteStep(steps.Pre, workingDir)
	}

	/**
	 * It's possible to override default behavior of install dependencies via
	 * the "dependencies.override" step.
	 */
	if len(steps.Override) > 0 {
		lib.Catchy("Running 'dependencies:override' commands...\n")

		lib.ExecuteStep(steps.Override, workingDir)
	} else {
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
	}

	wg.Wait()

	if len(steps.Post) > 0 {
		lib.Catchy("Running 'dependencies:post' commands...\n")

		lib.ExecuteStep(steps.Post, workingDir)
	}
}
