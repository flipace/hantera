package lib

import (
	"os"
	"os/exec"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Run : executes a process with the given parameters
func Run(progress bool, target string, name string, commands ...string) {
	cmd := exec.Command(name, commands...)
	cmd.Dir = target

	if progress {
		cmd.Stdout = os.Stdout
	}

	err := cmd.Start()
	check(err)

	err = cmd.Wait()
	check(err)
}
