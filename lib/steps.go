package lib

// ExecuteStep : executes all commands of a single step
func ExecuteStep(commands []string, workingDirectory string) {
	for _, command := range commands {
		Notice("$ %s\n", command)

		out, _ := Run(true, workingDirectory, true, command)

		outString := out.String()

		if len(outString) > 0 {
			println(outString)
		}
	}
}
