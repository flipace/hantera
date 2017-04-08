package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/flipace/hantera/lib"

	git "gopkg.in/src-d/go-git.v4"
)

// UpdateDev : updates existing working copies
func UpdateDev(arguments map[string]string) {
	config := lib.GetProductConfig(arguments["configFile"])

	refName := "refs/heads/develop"

	fmt.Printf("\"%s\" v%s\n\n", config.Name, config.Version)

	println("Updating dependencies...\n")

	for key, value := range config.Dependencies {
		targetDir := path.Join(arguments["targetDir"], key)

		// check whether the dependency directory already exists
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			fmt.Printf(">> New dependency %s, cloning to %s...", key, targetDir)

			_, err := lib.Clone(
				value.Repository,
				targetDir,
				refName,
				false,
			)

			if err != nil {
				panic(err)
			}
		} else {
			fmt.Printf(">> Dependency %s, pulling...\n", key)

			err := lib.Pull(targetDir, refName, true)

			switch err {
			case git.NoErrAlreadyUpToDate:
				println(">> - Already up to date!")
			case nil:
				println(">> - Updated!")
			default:
				println(err)
			}
		}
	}

	println("\nFinished update.")
}
