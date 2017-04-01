package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/tboerger/redirects/store"
	"gopkg.in/urfave/cli.v2"
	"os"
	"text/template"
)

// tmplShow represents a specific redirect detail view.
var tmplShow = "ID: \x1b[33m{{ .ID }}\x1b[0m" + `
Source: {{ .Source }}
Destination: {{ .Destination }}
Priority: {{ .Priority }}
`

// Show provides the sub-command to show redirect patterns.
func Show() *cli.Command {
	return &cli.Command{
		Name:      "show",
		Usage:     "Show a redirect pattern",
		ArgsUsage: "<id>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Value: tmplShow,
				Usage: "Custom output format",
			},
			&cli.BoolFlag{
				Name:  "json",
				Usage: "Print in JSON format",
			},
		},
		Action: func(c *cli.Context) error {
			return Handle(c, handleShow)
		},
	}
}

func handleShow(c *cli.Context, s store.Store) error {
	if c.NArg() != 1 {
		return cli.ShowSubcommandHelp(c)
	}

	record, err := s.GetRedirect(
		c.Args().First(),
	)

	if err != nil {
		return err
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(record, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}
