package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
	"gopkg.in/ini.v1"
)

func main() {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Falha ao ler arquivo de configuracao: %v", err)
		os.Exit(1)
	}

	fmt.Println("Working on env:", cfg.Section("").Key("env").String())
	fmt.Println("Checando dependencias:", cfg.Section("local-dependencies").Key("npm-packages").String())

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "Valida diferentes recursos",
				Action: func(cCtx *cli.Context) error {
					return checkAPI(cCtx)

				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// func runAction() error {
// 	cmd := exec.Command("node", "-v")

// 	output, err := cmd.Output()

// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("Node version: %v", string(output))

// 	return nil
// }

func checkAPI(ctx *cli.Context) error {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		return err
	}

	fmt.Println("The status code we got is:", resp.StatusCode)

	return nil
}
