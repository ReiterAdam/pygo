package cmd

import (
	"fmt"

	"github.com/ReiterAdam/pygo/internal/helpers"
	"github.com/urfave/cli/v2"
)

func SetupCommand() *cli.Command {
	return &cli.Command{
		Name:    "setup",
		Aliases: []string{"s"},
		Usage:   "create venv",
		Action: func(cCtx *cli.Context) error {

			// Check if venv folder is present
			is_venv, _ := helpers.IsVenv()
			if is_venv {
				return cli.Exit("Venv already here", 80)
			}

			// Venv not found - create
			// Prepare command
			cmdArgs := []string{"bash", "-c", "python -m venv .venv"}
			fmt.Println(cmdArgs)

			if err := helpers.ExecuteCommand(cmdArgs); err != nil {
				return cli.Exit("Could not setup venv", 82)
			}

			// Confirm to user
			fmt.Println("Setup finished")
			return nil
		},
	}
}
