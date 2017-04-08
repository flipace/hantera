package commands

import (
	"fmt"
	"path"

	"github.com/flipace/hantera/lib"
)

// InitDev : does something cool
func InitDev(arguments map[string]string) {
	config := lib.GetProductConfig(arguments["configFile"])

	fmt.Printf("Loaded config for: \"%s\" v%s\n", config.Name, config.Version)

	println("\nCloning dependencies...")

	// the ref to checkout
	refName := "refs/heads/develop"

	for key, value := range config.Dependencies {
		targetDir := path.Join(arguments["targetDir"], key)

		fmt.Printf("--| Cloning %s to %s\n", key, targetDir)

		_, err := lib.Clone(value.Repository, targetDir, refName, true)

		if err != nil {
			panic(err)
		}
	}
}
