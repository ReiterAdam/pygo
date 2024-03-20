package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() error {
	app := &cli.App{
		Name:  "Pygo",
		Usage: "Helps with venv managment",
		Commands: []*cli.Command{
			SetupCommand(),
			RunCommand(),
			AddCommand(),
			RemoveCommand(),
		},
	}

	return app.Run(os.Args)
}
