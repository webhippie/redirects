package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// Create provides the sub-command to create redirect patterns.
func Create() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "Create a redirect pattern",
		Action: func(c *cli.Context) {
			logrus.Info("Create")
		},
	}
}
