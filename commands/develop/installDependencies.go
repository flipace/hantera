package develop

import (
	"os"
	"os/exec"
	"path"

	"github.com/flipace/hantera/lib"
	"github.com/urfave/cli"
)

// InstallDependencies : installs project dependencies (tries to figure out package manager e.g. npm)
func InstallDependencies(c *cli.Context) {
	configFile := c.String("config")

	config := lib.GetProductConfig(configFile)

	target := c.String("target")
	if target == "" {
		target = path.Join("./", config.Name)
	}

	lib.Catchy("Installing dependencies for \"%s\" v%s...\n", config.Name, config.Version)

	for key := range config.Dependencies {
		targetDir := path.Join(target, key)

		if _, err := os.Stat(path.Join(targetDir, "package.json")); err == nil {
			lib.Notice("|-- Found package.json for %s - running 'yarn'\n", key)

			cmd := exec.Command("yarn", "install")
			cmd.Dir = targetDir
			cmd.Stdout = os.Stdout

			err = cmd.Start()

			if err != nil {
				panic(err)
			}

			err = cmd.Wait()

			if err != nil {
				panic(err)
			}
		}
	}
}
