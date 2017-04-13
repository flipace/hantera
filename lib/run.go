package lib

import (
	"bytes"
	"fmt"
	"os/exec"
)

func check(err error, outBuf bytes.Buffer, errBuf bytes.Buffer) {
	if err != nil {
		fmt.Printf(
			"An error occured.\n\nCommand output:%s\nCommand error:\n%s",
			outBuf.String(),
			errBuf.String(),
		)
	}
}

// Run : executes a process with the given parameters
func Run(progress bool, target string, handleError bool, name string, commands ...string) (bytes.Buffer, bytes.Buffer) {
	cmd := exec.Command(name, commands...)
	cmd.Dir = target

	var outBuff bytes.Buffer
	var errBuff bytes.Buffer

	if progress {
		cmd.Stdout = &outBuff
		cmd.Stderr = &errBuff
	}

	err := cmd.Start()
	if handleError {
		check(err, outBuff, errBuff)
	}

	err = cmd.Wait()
	if handleError {
		check(err, outBuff, errBuff)
	}

	return outBuff, errBuff
}
