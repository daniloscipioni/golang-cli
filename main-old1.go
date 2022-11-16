package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main2() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Value:   "default",
				Aliases: []string{"n"},
				Usage:   "Adiciona flag com o nome do usuario",
			},
			&cli.StringFlag{
				Name:    "cidade",
				Value:   "Campinas",
				Aliases: []string{"c"},
				Usage:   "Informa o nome da cidade",
			},
		},
		Action: func(ctx *cli.Context) error {
			username := ctx.String("name")
			cidade := ctx.String("cidade")

			if username != "" {
				fmt.Println("Ola", username)
			} else {
				fmt.Println("Flag --name nao foi informada", username)
			}

			if cidade != "" {
				fmt.Println("Cidade, %s", cidade)
			} else {
				fmt.Println("Flag --cidade nao foi informada", cidade)
			}
			return nil

		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
