package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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
	// Return false if error is not due to not existance
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
						return cli.Exit("Error checking for virtual environment", 80)
					} else if is_venv {
						return cli.Exit("Virtual environment already exists", 81)
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

					fmt.Println("Running main.py...")
					// Try to run program without venv
					is_venv, _ := is_venv()
					if !is_venv {
						fmt.Println("Note, venv is not present.")
						cmd := exec.Command("python", "main.py")
						// Execute and check for errors
						_, err := cmd.Output()
						if err != nil {
							fmt.Println(err)
							return cli.Exit("Could not run python program", 90)
						}
						// source venv and then run
					} else {
						// source venv
						// Check current shell
						shell := os.Getenv("SHELL")
						if shell != "" {
							fmt.Println("Current shell:", shell)
						} else {
							fmt.Println("Unable to determine current shell.")
						}
						// run app
						// deactivate venv
					}

					fmt.Println("Program finished.")
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
