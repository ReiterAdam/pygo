package cmd

import (
	"fmt"

	"github.com/ReiterAdam/pygo/internal/helpers"
	"github.com/urfave/cli/v2"
)

func RunCommand() *cli.Command {
	return &cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "run main.py from current directory",
		Action: func(cCtx *cli.Context) error {

			globalVenv := false

			// Check if venv is present
			is_venv, _ := helpers.IsVenv(globalVenv)
			if !is_venv {
				fmt.Printf("Note, venv is not present.\n\n")
			}

			var cmdArgs []string

			// Add venv sourcing
			if is_venv {
				cmdArgs = []string{"bash", "-c", "source .venv/bin/activate && python main.py"}
			} else {
				cmdArgs = []string{"bash", "-c", "python main.py"}
			}

			// Prepare command
			// Get command line arguments from user
			argumentsFmt := helpers.PrepareUserArguments(fmt.Sprint(cCtx.Args()))

			// Form command to run
			cmdArgs[len(cmdArgs)-1] = cmdArgs[len(cmdArgs)-1] + " " + argumentsFmt

			// Execute
			if err := helpers.ExecuteCommand(cmdArgs); err != nil {
				return cli.Exit("Could not run program", 92)
			}

			fmt.Println("\nProgram finished.")
			return nil
		},
	}
}
