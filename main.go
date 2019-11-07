package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/urfave/cli"
)

var Version = "0.0.3"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Usage = "[name] [command...]"

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "env, e",
			Usage: "environment",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() < 2 {
			return fmt.Errorf("shortage args")
		}

		envArgs := c.StringSlice("env")
		name := c.Args().Get(0)
		command := c.Args()[1:]

		var envs []*ecs.KeyValuePair
		for _, e := range envArgs {
			r := strings.Split(e, "=")
			if len(r) != 2 {
				return fmt.Errorf("bad env '%s'", e)
			}
			envs = append(envs, &ecs.KeyValuePair{
				Name:  aws.String(r[0]),
				Value: aws.String(r[1]),
			})
		}

		co := ecs.ContainerOverride{
			Name:        aws.String(name),
			Command:     aws.StringSlice(command),
			Environment: envs,
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
