package main

import (
	"os/exec"
	"fmt"
	"strings"
)

/*
 * This will run the provided "script"
 * and return the exitcode on success.
 * Use this to determine if if we need 
 * to run a remedy.
 */
func RunScript(script string) int {
	cmdStr := strings.Split(script, " ")
	cmd := exec.Command(cmdStr[0], cmdStr[1:]...)

	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to run command: %s\n", err)
	}

	if err := cmd.Wait(); err != nil {
		if exitcode, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Exitcode: %d\n", exitcode.ExitCode())
			return exitcode.ExitCode()
		} else {
			fmt.Printf("Unable to read exitcode: %s\n", err)
		}
	}

	return 0
}
