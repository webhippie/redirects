package cmd

import (
	"github.com/tboerger/redirects/store"
	"github.com/urfave/cli"
)

// Remove provides the sub-command to reove redirect patterns.
func Remove() cli.Command {
	return cli.Command{
		Name:      "remove",
		Usage:     "Remove a redirect pattern",
		ArgsUsage: "<id>",
		Action: func(c *cli.Context) error {
			return Handle(c, handleRemove)
		},
	}
}

func handleRemove(c *cli.Context, s store.Store) error {
	if c.NArg() != 1 {
		return cli.ShowSubcommandHelp(c)
	}

	return s.DeleteRedirect(c.Args().First())
}
