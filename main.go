/*
Copyright 2017 Codem8s.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codem8s/autograph/generate"
	"github.com/codem8s/autograph/server"
	"github.com/golang/glog"
	"github.com/urfave/cli"
)

// This file implements common CLI operations:
// * generate - generate a key and certificate pair which can be used to sign/verify
// * sign - sign a manifest
// * verify - verify a signed manifest
// * run - starts the HTTP(S) server
func main() {
	flag.Parse()
	app := cli.NewApp()
	app.Name = "autograph"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate RSA key and certificate pair",
			Action: func(c *cli.Context) error {
				_, _, err := generate.KeyPair()
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
			Aliases: []string{"r"},
			Usage:   "starts the HTTP(S) server",
			Action: func(c *cli.Context) error {
				return server.Run()
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		glog.Fatal(err)
	}
}
