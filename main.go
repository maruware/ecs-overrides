package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var Version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Version = Version

	app.Commands = []cli.Command{
		cmdCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
