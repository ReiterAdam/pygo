package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "setup",
				Aliases: []string{"s"},
				Usage:   "create venv and project template",
				Action: func(cCtx *cli.Context) error {

					// Check if venv folder is present
					is_venv, _ := isVenv()
					if is_venv {
						return cli.Exit("Venv already here", 80)
					}

					// Venv not found - create
					// Prepare command
					cmdArgs := []string{"bash", "-c", "python -m venv .venv"}
					fmt.Println(cmdArgs)

					if err := executeCommand(cmdArgs); err != nil {
						return cli.Exit("Could not setup venv", 82)
					}

					// Confirm to user
					fmt.Println("Setup finished")
					return nil
				},
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "run application from current directory",
				Action: func(cCtx *cli.Context) error {

					// Check if venv is present
					is_venv, _ := isVenv()
					if !is_venv {
						fmt.Println("Note, venv is not present.")
					}

					var cmdArgs []string

					// Add venv sourcing
					if is_venv {
						cmdArgs = []string{"bash", "-c", "source .venv/bin/activate && python main.py"}
					} else {
						cmdArgs = []string{"python", "main.py"}
					}

					// Prepare command
					// Get command line arguments from user
					argumentsFmt := prepareUserArguments(fmt.Sprint(cCtx.Args()))
					// Form command to run
					cmdArgs[len(cmdArgs)-1] = cmdArgs[len(cmdArgs)-1] + " " + argumentsFmt

					// Execute
					if err := executeCommand(cmdArgs); err != nil {
						return cli.Exit("Could not run program", 92)
					}

					fmt.Println("\nProgram finished.")
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a package to the venv and requirements",
				Action: func(cCtx *cli.Context) error {

					// Check if venv is present
					is_venv, _ := isVenv()
					if !is_venv {
						return cli.Exit("Could not detect virtual environment", 100)
					}

					// Form command to run
					cmdArgs := []string{"bash", "-c", "source .venv/bin/activate && pip install"}
					argumentsFmt := prepareUserArguments(fmt.Sprint(cCtx.Args()))
					cmdArgs[len(cmdArgs)-1] = cmdArgs[len(cmdArgs)-1] + " " + argumentsFmt

					// Execute
					if err := executeCommand(cmdArgs); err != nil {
						return cli.Exit("Could not install package", 102)
					}

					fmt.Println("Added package: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func isVenv() (bool, error) {
	// Function checks if venv directory is made

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return false, err
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

func executeCommand(cmdArgs []string) error {

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

func prepareUserArguments(arguments string) string {

	argumentsFmt := strings.Fields(arguments[2 : len(arguments)-1])
	args2 := strings.Join(argumentsFmt, " ")

	return args2
}
