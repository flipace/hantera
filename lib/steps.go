package lib

// ExecuteStep : executes all commands of a single step
func ExecuteStep(commands []string, workingDirectory string) {
	for _, command := range commands {
		out, _ := Run(true, workingDirectory, true, command)

		println(out.String())
	}
}
