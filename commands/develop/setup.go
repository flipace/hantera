package develop

import (
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// RunSetup : runs the setup procedure
func RunSetup(configFile string, target string, branch string, nodeps bool, c *cli.Context) {
	config := lib.GetProductConfig(configFile)

	steps := config.Steps.Setup

	workingDir, _ := filepath.Abs(target)

	if len(steps.Pre) > 0 {
		lib.Notice("Running 'setup:pre' commands...\n")

		lib.ExecuteStep(steps.Pre, workingDir)
	}

	lib.Catchy("Setup \"%s\" v%s\n", config.Name, config.Version)

	var wg sync.WaitGroup

	if len(steps.Override) > 0 {
		lib.Warning("Running 'setup:override' commands... (CAUTION: This overrides hanteras core behavior!)\n")

		lib.ExecuteStep(steps.Override, workingDir)
	} else {
		// create and open .gitignore in target path - we put all dependencies into it
		gitignore, err := os.Create(path.Join(target, ".gitignore"))
		if err != nil {
			panic(err)
		}
		defer gitignore.Close()

		for key, value := range config.Dependencies {
			wg.Add(1)

			targetDir, err := filepath.Abs(path.Join(target, key))
			check(err)

			go CloneRepository(key, targetDir, branch, value.Repository, nodeps, &wg)
		}
	}

	wg.Wait()

	println("")

	if nodeps == false {
		InstallDependencies(c)
	}

	if len(steps.Post) > 0 {
		lib.Notice("Running 'setup:post' commands...\n")

		lib.ExecuteStep(steps.Post, workingDir)
	}

	lib.Catchy("\nDone with setup of %s!\n", config.Name)
}

// Setup : sets up development environment for a project
func Setup(c *cli.Context) {
	configFile := c.String("config")

	target := c.String("target")
	branch := c.String("branch")
	nodeps := c.Bool("no-deps")

	RunSetup(configFile, target, branch, nodeps, c)
}
