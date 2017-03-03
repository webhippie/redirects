package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// List provides the sub-command to list redirect patterns.
func List() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "List available redirect patterns",
		Action: func(c *cli.Context) {
			logrus.Info("List")
		},
	}
}
