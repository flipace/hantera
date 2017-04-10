package develop

import (
	"os"
	"path"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// Clean : remove all dependencies of the project
func Clean(c *cli.Context) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")

	lib.Catchy("Cleaning dependencies of \"%s\" v%s...\n", config.Name, config.Version)

	// create a semaphore with l

	var wg sync.WaitGroup

	for key := range config.Dependencies {
		wg.Add(1)

		go func(target string, key string) {
			lib.Notice("--| Deleting %s\n", key)

			targetDir := path.Join(target, key)

			os.RemoveAll(targetDir)

			wg.Done()
		}(target, key)
	}

	wg.Wait()
}
