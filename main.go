package main

import (
	"fmt"
	"github.com/codem8s/autograph/generate"
	"github.com/urfave/cli"
	"log"
	"os"
)

// This file implements common CLI operations:
// * generate - generate a key and certificate pair which can be used to sign/verify
// * sign - sign a manifest
// * verify - verify a signed manifest
// * run - starts the HTTP(S) server

func main() {
	app := cli.NewApp()
	app.Name = "autograph"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate RSA key and certificate pair",
			Action: func(c *cli.Context) error {
				_, _, err := generate.GenerateKeyPair()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "sign",
			Aliases: []string{"s"},
			Usage:   "sign a manifest",
			Action: func(c *cli.Context) error {
				fmt.Println("sign is not implemented yet")
				return nil
			},
		},
		{
			Name:    "verify",
			Aliases: []string{"v"},
			Usage:   "verify a signed manifest",
			Action: func(c *cli.Context) error {
				fmt.Println("verify is not implemented yet")
				return nil
			},
		},
		{
			Name:    "run",
			Aliases: []string{"v"},
			Usage:   "starts the HTTP(S) server",
			Action: func(c *cli.Context) error {
				fmt.Println("run is not implemented yet")
				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
