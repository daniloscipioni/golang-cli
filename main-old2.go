package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main3() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "Valida diferentes recursos",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Recurso validado ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Cria um novo recurso",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Novo recurso criado: ", cCtx.Args().First())

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
