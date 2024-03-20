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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "type",
				Usage:    "type of venv",
				Required: false,
			},
		},
		Action: func(cCtx *cli.Context) error {

			globalVenv := false
			if cCtx.String("type") == "global" {
				globalVenv = true
			}

			// Check if venv is present
			is_venv, _ := helpers.IsVenv(globalVenv)
			if !is_venv {
				fmt.Printf("Note, chosen venv is not present.")
			}

			fmt.Println()
			var cmdArgs []string

			// Add venv sourcing
			if is_venv && globalVenv {
				cmdArgs = []string{"bash", "-c", "source ~/.pygo/.venv/bin/activate && python src/main.py"}
			} else if is_venv && !globalVenv {
				cmdArgs = []string{"bash", "-c", "source .venv/bin/activate && python src/main.py"}
			} else {
				cmdArgs = []string{"bash", "-c", "python src/main.py"}
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
