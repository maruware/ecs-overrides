package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/urfave/cli"
)

var Version = "0.0.2"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Usage = "[name] [command...]"
	app.Action = func(c *cli.Context) error {
		if c.NArg() < 2 {
			return fmt.Errorf("shortage args")
		}
		name := c.Args().Get(0)
		command := c.Args()[1:]

		co := ecs.ContainerOverride{
			Name:    aws.String(name),
			Command: aws.StringSlice(command),
		}
		cos := []*ecs.ContainerOverride{&co}
		ov := ecs.TaskOverride{
			ContainerOverrides: cos,
		}
		j, err := json.Marshal(ov)
		if err != nil {
			return err
		}
		fmt.Print(string(j))

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
