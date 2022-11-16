package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main5() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "Valida diferentes recursos",
				Action: func(cCtx *cli.Context) error {
					return runAction()

				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runAction() error {
	cmd := exec.Command("node", "-v")

	output, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Printf("Node version: %v", string(output))

	return nil
}
