package develop

import (
	"path"
	"path/filepath"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// Update : updates each dependencies working copy to latest version from head
func Update(c *cli.Context) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")
	progress := c.Bool("progress")

	lib.Catchy("Updating dependencies for \"%s\" v%s...\n", config.Name, config.Version)

	var wg sync.WaitGroup

	for key := range config.Dependencies {
		wg.Add(1)

		lib.Notice("--| Updating %s\n", key)

		targetDir, err := filepath.Abs(path.Join(target, key))
		check(err)

		go func(workingDir string, progress bool) {

			lib.Run(progress, targetDir, "git", "pull", "-r")

			wg.Done()
		}(targetDir, progress)
	}

	wg.Wait()
}
