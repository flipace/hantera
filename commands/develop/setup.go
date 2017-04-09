package develop

import (
	"os"
	"path"
	"sync"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// Setup : sets up development environment for a project
func Setup(c *cli.Context) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")
	refName := c.String("branch")
	progress := c.Bool("progress")
	nodeps := c.Bool("no-deps")

	lib.Catchy("Setup \"%s\" v%s\n", config.Name, config.Version)

	println(nodeps)
	// create and open .gitignore in target path - we put all dependencies into it
	gitignore, err := os.Create(path.Join(target, ".gitignore"))
	if err != nil {
		panic(err)
	}
	defer gitignore.Close()

	var wg sync.WaitGroup

	for key, value := range config.Dependencies {
		wg.Add(1)

		targetDir := path.Join(target, key)

		lib.Notice("--| Cloning %s to %s\n", key, targetDir)

		go func(
			name string,
			repository string,
			target string,
			refName string,
			progress bool,
		) {
			lib.Clone(repository, target, refName, progress)

			gitignore.WriteString(name + "\n")
			gitignore.Sync()

			wg.Done()
		}(key, value.Repository, targetDir, refName, progress)
	}

	wg.Wait()

	if nodeps == false {
		InstallDependencies(c)
	}
}
