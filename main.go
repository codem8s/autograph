package main

import (
  "fmt"
  "github.com/urfave/cli"
  "os"
  "generate"
)

func main() {
  app := cli.NewApp()
  app.Name = "autograph"
  app.Version = "0.1"

  app.Commands = []cli.Command{
    {
      Name:    "generate",
      Aliases: []string{"g"},
      Usage:   "generate a key and certificate pair",
      Action:  func(c *cli.Context) error {
        fmt.Println("generate is not implemented yet")
        generate.GenerateKeyPair()
        return nil
      },
    },
    {
      Name:    "sign",
      Aliases: []string{"s"},
      Usage:   "sign a manifest",
      Action:  func(c *cli.Context) error {
        fmt.Println("sign is not implemented yet")
        return nil
      },
    },
    {
      Name:        "verify",
      Aliases:     []string{"v"},
      Usage:       "verify a signed manifest",
      Action:  func(c *cli.Context) error {
        fmt.Println("verify is not implemented yet")
        return nil
      },
    },
    {
      Name:        "run",
      Aliases:     []string{"v"},
      Usage:       "starts the HTTP(S) server",
      Action:  func(c *cli.Context) error {
        fmt.Println("run is not implemented yet")
        return nil
      },
    },
  }

  err := app.Run(os.Args)

  if err != nil {
    os.Exit(1)
  }
}
