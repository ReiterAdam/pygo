package cmd

import (
	"fmt"

	"github.com/ReiterAdam/pygo/internal/helpers"
	"github.com/urfave/cli/v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a package to the venv",
		Action: func(cCtx *cli.Context) error {

			// Check if venv is present
			is_venv, _ := helpers.IsVenv()
			if !is_venv {
				return cli.Exit("Could not detect virtual environment", 100)
			}

			// Form command to run
			cmdArgs := []string{"bash", "-c", "source .venv/bin/activate && pip install"}
			argumentsFmt := helpers.PrepareUserArguments(fmt.Sprint(cCtx.Args()))
			cmdArgs[len(cmdArgs)-1] = cmdArgs[len(cmdArgs)-1] + " " + argumentsFmt

			if err := helpers.ExecuteCommand(cmdArgs); err != nil {
				return cli.Exit("Could not install package", 102)
			}

			fmt.Println("Added package: ", cCtx.Args().First())
			return nil
		},
	}
}
