package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main4() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "Valida diferentes recursos",
				Subcommands: []*cli.Command{
					{
						Name:  "env",
						Usage: "Valida os requisitos do ambiente",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Tarefa completada: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "project",
						Usage: "Valida a estrutura de um projeto",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Tarefa completada: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Completa uma tarefa da lista",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Tarefa completada: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
