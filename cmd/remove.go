package cmd

import (
	"fmt"

	"github.com/ReiterAdam/pygo/internal/helpers"
	"github.com/urfave/cli/v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"rm"},
		Usage:   "remove a package from venv",
		Action: func(cCtx *cli.Context) error {

			// Check if venv is present
			is_venv, _ := helpers.IsVenv()
			if !is_venv {
				return cli.Exit("Could not detect virtual environment", 100)
			}

			// Form command to run
			cmdArgs := []string{"bash", "-c", "source .venv/bin/activate && pip remove"}
			argumentsFmt := helpers.PrepareUserArguments(fmt.Sprint(cCtx.Args()))
			cmdArgs[len(cmdArgs)-1] = cmdArgs[len(cmdArgs)-1] + " " + argumentsFmt

			// Execute
			if err := helpers.ExecuteCommand(cmdArgs); err != nil {
				return cli.Exit("Could not remove a package", 102)
			}

			fmt.Println("Added package: ", cCtx.Args().First())
			return nil
		},
	}
}
