package cmd

import (
	"fmt"

	"github.com/ReiterAdam/pygo/internal/helpers"
	"github.com/urfave/cli/v2"
)

func TestCommand() *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "run tests from project root",
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
				return cli.Exit("Venv is not present.", 94)
			}

			fmt.Println()
			var cmdArgs []string

			// Add venv sourcing
			if is_venv && globalVenv {
				cmdArgs = []string{"bash", "-c", "source ~/.pygo/.venv/bin/activate && pytest tests/"}
			} else if is_venv && !globalVenv {
				cmdArgs = []string{"bash", "-c", "source .venv/bin/activate && pytest tests/"}
			} else {
				return cli.Exit("Could not run tests", 95)
			}

			// Prepare command
			// Get command line arguments from user
			argumentsFmt := helpers.PrepareUserArguments(fmt.Sprint(cCtx.Args()))

			// Form command to run
			cmdArgs[len(cmdArgs)-1] = cmdArgs[len(cmdArgs)-1] + " " + argumentsFmt

			// Execute
			if err := helpers.ExecuteCommand(cmdArgs); err != nil {
				return cli.Exit("Could not run tests", 95)
			}

			fmt.Println("\nTests finished.")
			return nil
		},
	}
}
