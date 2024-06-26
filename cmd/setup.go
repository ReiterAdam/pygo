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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "type",
				Usage:    "type of venv (local or global)",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {

			globalVenv := false
			if cCtx.String("type") == "global" {
				globalVenv = true
			}

			// Check if venv folder is present
			is_venv, _ := helpers.IsVenv(globalVenv)
			if is_venv {
				return cli.Exit("Venv already here", 80)
			}

			// Venv not found - create
			// Prepare command
			cmdArgs := []string{"bash", "-c", "python -m venv .venv"}
			if globalVenv {
				cmdArgs = []string{"bash", "-c", "python -m venv ~/.pygo/.venv"}
			}

			// fmt.Println()

			if err := helpers.ExecuteCommand(cmdArgs); err != nil {
				return cli.Exit("Could not setup venv", 82)
			}

			cmdArgs = []string{"bash", "-c", "source .venv/bin/activate && pip install mypy pytest"}
			if globalVenv {
				cmdArgs = []string{"bash", "-c", "source ~/.pygo/.venv/bin/activate && pip install mypy pytest"}
			}

			if err := helpers.ExecuteCommand(cmdArgs); err != nil {
				return cli.Exit("Could not setup tools", 83)
			}

			cmdArgs = []string{"bash", "-c", "mkdir src tests"}
			if err := helpers.ExecuteCommand(cmdArgs); err != nil {
				return cli.Exit("Could not setup directories", 84)
			}

			cmdArgs = []string{"bash", "-c", "touch src/__init__.py src/main.py tests/__init__.py"}
			if err := helpers.ExecuteCommand(cmdArgs); err != nil {
				return cli.Exit("Could not setup starter files", 85)
			}

			// Confirm to user
			fmt.Println("Setup finished")
			return nil
		},
	}
}
