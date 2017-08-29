package main

import (
  "fmt"
  "github.com/urfave/cli"
  "os"
)

func main() {
  app := cli.NewApp()
  app.Name = "autograph"
  app.Version = "0.1"

  app.Action = func(c *cli.Context) error {
    fmt.Printf("Autograph")
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    os.Exit(1)
  }
}
