package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// Update provides the sub-command to update redirect patterns.
func Update() cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "Update a redirect pattern",
		Action: func(c *cli.Context) {
			logrus.Info("Update")
		},
	}
}
