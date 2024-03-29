package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() error {
	app := &cli.App{
		Name:  "pygo",
		Usage: "Helps with venv management",
		Commands: []*cli.Command{
			SetupCommand(),
			RunCommand(),
			AddCommand(),
			RemoveCommand(),
			TestCommand(),
		},
	}

	return app.Run(os.Args)
}
