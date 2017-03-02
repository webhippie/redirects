package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// Remove provides the sub-command to reove redirect patterns.
func Remove() cli.Command {
	return cli.Command{
		Name:  "remove",
		Usage: "Remove a redirect pattern",
		Action: func(c *cli.Context) {
			logrus.Info("Remove")
		},
	}
}
