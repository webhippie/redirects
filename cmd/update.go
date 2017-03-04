package cmd

import (
	"fmt"
	"github.com/tboerger/redirects/store"
	"gopkg.in/urfave/cli.v2"
	"os"
)

// Update provides the sub-command to update redirect patterns.
func Update() *cli.Command {
	return &cli.Command{
		Name:      "update",
		Usage:     "Update a redirect pattern",
		ArgsUsage: "<id>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "source",
				Value: "",
				Usage: "Source for the redirect",
			},
			&cli.StringFlag{
				Name:  "destination",
				Value: "",
				Usage: "Destination for the redirect",
			},
		},
		Action: func(c *cli.Context) error {
			return Handle(c, handleUpdate)
		},
	}
}

func handleUpdate(c *cli.Context, s store.Store) error {
	if c.NArg() != 1 {
		return cli.ShowSubcommandHelp(c)
	}

	record, err := s.GetRedirect(
		c.Args().First(),
	)

	if err != nil {
		return err
	}

	changed := false

	if val := c.String("source"); c.IsSet("source") && val != record.Source {
		record.Source = val
		changed = true
	}

	if val := c.String("destination"); c.IsSet("destination") && val != record.Destination {
		record.Destination = val
		changed = true
	}

	if changed {
		return s.UpdateRedirect(record)
	}

	fmt.Fprintf(os.Stderr, "Nothing to update...\n")
	return nil
}
