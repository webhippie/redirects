package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// Show provides the sub-command to show redirect patterns.
func Show() cli.Command {
	return cli.Command{
		Name:  "show",
		Usage: "Show a redirect pattern",
		Action: func(c *cli.Context) {
			logrus.Info("Show")
		},
	}
}
