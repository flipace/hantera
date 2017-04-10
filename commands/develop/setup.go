package develop

import (
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// Setup : sets up development environment for a project
func Setup(c *cli.Context) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")
	branch := c.String("branch")
	progress := c.Bool("progress")
	nodeps := c.Bool("no-deps")

	lib.Catchy("Setup \"%s\" v%s\n", config.Name, config.Version)

	// create and open .gitignore in target path - we put all dependencies into it
	gitignore, err := os.Create(path.Join(target, ".gitignore"))
	if err != nil {
		panic(err)
	}
	defer gitignore.Close()

	var wg sync.WaitGroup

	for key, value := range config.Dependencies {
		wg.Add(1)

		workingDir, err := filepath.Abs(path.Join(target))
		check(err)

		targetDir, err := filepath.Abs(path.Join(target, key))
		check(err)

		lib.Notice("--| Cloning %s to %s\n", key, targetDir)

		go func(
			name string,
			repository string,
			target string,
			branch string,
			progress bool,
		) {
			lib.Run(progress, workingDir, "git", "clone", repository, target)
			lib.Run(progress, target, "git", "checkout", branch)

			wg.Done()
		}(key, value.Repository, targetDir, branch, progress)
	}

	wg.Wait()

	if nodeps == false {
		InstallDependencies(c)
	}
}
