package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func IsVenv(globalVenv bool) (bool, error) {
	// Function checks if venv directory is made

	// Get directory to check
	dir, err := os.Getwd()
	if err != nil {
		return false, err
	}
	if globalVenv {
		dir = "~/.config/pygo"
	}

	// Check if the "venv" folder exists in the current directory
	_, err = os.Stat(dir + "/.venv/")
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	// Return false if error is not due to not existence
	return false, err
}

func ExecuteCommand(cmdArgs []string) error {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	// Set up pipes for interacting with the command
	// Source: ChatGPT
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return err
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command:", err)
		return err
	}
	return nil
}

func PrepareUserArguments(arguments string) string {

	argumentsFmt := strings.Fields(arguments[2 : len(arguments)-1])
	args2 := strings.Join(argumentsFmt, " ")

	return args2
}
