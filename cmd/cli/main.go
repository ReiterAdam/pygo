package main

import (
    "fmt"
    "log"
    "os"
	"os/exec"

    "github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Commands: []*cli.Command{
            {
                Name:    "setup",
                Aliases: []string{"s"},
                Usage:   "create venv and project template",
                Action: func(cCtx *cli.Context) error {
					
					// Check if venv folder is present

					// Get the current working directory
					dir, err := os.Getwd()
					if err != nil {
						return cli.Exit("Error getting current directory.", 83)
					}

					// Check if the "venv" folder exists in the current directory
					_, err = os.Stat(dir + "/.venv")
					if err == nil {
						return cli.Exit("The 'venv' folder exists in the current directory.", 84)
					}
					if err != nil && !os.IsNotExist(err){
						return cli.Exit("Error checking for 'venv' folder", 85)
					}
					
					// Venv not found - create
					// Prepare command 
					cmd := exec.Command("python", "-m", "venv", ".venv")
					
					// Execute and check for errors
					_, err = cmd.Output()
					if err != nil {
						return cli.Exit("Could not create new virtual environment", 86)
					}
					
					// Confirm to user
                    fmt.Println("Setup finished")
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
