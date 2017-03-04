package cmd

import (
	"fmt"
	"github.com/tboerger/redirects/store"
	"github.com/urfave/cli"
	"strconv"
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

	id, err := strconv.Atoi(c.Args().First())

	if err != nil {
		return fmt.Errorf("Failed to parse the ID")
	}

	return s.DeleteRedirect(id)
}
