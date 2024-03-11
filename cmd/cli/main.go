package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func is_venv() (bool, error) {
	// Function checks if venv directory is made

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return false, err
	}

	// Check if the "venv" folder exists in the current directory
	_, err = os.Stat(dir + "/.venv")
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	// Return false if error is not due to not existence
	return false, err

}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "setup",
				Aliases: []string{"s"},
				Usage:   "create venv and project template",
				Action: func(cCtx *cli.Context) error {

					// Check if venv folder is present
					is_venv, err := is_venv()
					if err != nil {
						fmt.Println(err)
					} else if is_venv {
						fmt.Println(err)
					}

					// Venv not found - create
					// Prepare command
					cmd := exec.Command("python", "-m", "venv", ".venv")

					// Execute and check for errors
					_, err = cmd.Output()
					if err != nil {
						return cli.Exit("Could not create new virtual environment", 82)
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

					fmt.Println(cCtx.Args())
					arguments := fmt.Sprint(cCtx.Args())
					argumentsFmt := strings.Fields(arguments[2 : len(arguments)-1])

					// Try to run program without venv
					is_venv, _ := is_venv()
					if !is_venv {
						fmt.Println("Note, venv is not present.")
					}

					// source venv
					if is_venv {
						cmd := exec.Command("source", ".venv/bin/activate")
						_, err := cmd.Output()
						if err != nil {
							return cli.Exit("Could not source virtual environment", 90)
						}
					}

					// Prepare command
					cmdArgs := append([]string{"main.py"}, argumentsFmt...)
					cmd := exec.Command("python", cmdArgs...)

					// Set up pipes for interacting with the command
					// Source: ChatGPT
					cmd.Stdin = os.Stdin
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					fmt.Print("Running main.py...\n\n")
					// Start the command
					if err := cmd.Start(); err != nil {
						fmt.Println("Error starting command:", err)
						return cli.Exit("Could not start python program", 91)
					}

					// Wait for the command to finish
					if err := cmd.Wait(); err != nil {
						fmt.Println("Error waiting for command:", err)
						return cli.Exit("Could not finish python program", 93)
					}

					// deactivate venv
					if is_venv {
						cmd := exec.Command("deactivate")
						_, err := cmd.Output()
						if err != nil {
							return cli.Exit("Could not deactivate virtual environment", 99)
						}
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
					fmt.Println("added package: ", cCtx.Args().First())
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
