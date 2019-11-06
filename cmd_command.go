package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/mattn/go-shellwords"
	"github.com/urfave/cli"
)

func cmdCommand() cli.Command {
	return cli.Command{
		Name:  "cmd",
		Usage: "cmd [name] [command]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				return fmt.Errorf("shortage args")
			}
			name := c.Args().Get(0)
			commandS := c.Args().Get(1)

			command, err := shellwords.Parse(commandS)
			if err != nil {
				return err
			}

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
		},
	}
}
