package develop

import (
	"fmt"
	"path"
	"path/filepath"
	"sync"

	"github.com/flipace/hantera/lib"
)

// CloneRepository : clones the given repository
func CloneRepository(name string, target string, branch string, repository string, nodeps bool, wg *sync.WaitGroup) {
	workingDir, err := filepath.Abs(path.Join(target, "../"))
	check(err)

	lib.Catchy(">> Cloning %s to %s\n", name, target)

	cloneOut, cloneErr := lib.Run(
		true,
		workingDir,
		false,
		"git",
		"clone",
		repository,
		target,
	)

	if len(cloneErr.String()) > 0 {
		println(cloneErr.String())
	}

	checkOut, checkErr := lib.Run(true, target, false, "git", "checkout", branch)
	if len(checkErr.String()) > 0 {
		println(checkErr.String())
	}

	fmt.Printf("--| %s:\n%s\n%s", name, cloneOut.String(), checkOut.String())

	wg.Done()
}
