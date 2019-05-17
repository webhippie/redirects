package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/dsielert/redirector/store"
	"gopkg.in/urfave/cli.v2"
)

// tmplList represents a row within redirect listing.
var tmplList = "ID: \x1b[33m{{ .ID }}\x1b[0m" + `
Source: {{ .Source }}
Destination: {{ .Destination }}
`

// List provides the sub-command to list redirect patterns.
func List() *cli.Command {
	return &cli.Command{
		Name:      "list",
		Usage:     "List available redirect patterns",
		ArgsUsage: " ",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Value: tmplList,
				Usage: "Custom output format",
			},
			&cli.StringFlag{
				Name:  "filter",
				Value: "",
				Usage: "Filter output by needle",
			},
			&cli.BoolFlag{
				Name:  "json",
				Usage: "Print in JSON format",
			},
		},
		Action: func(c *cli.Context) error {
			return Handle(c, handleList)
		},
	}
}

func handleList(c *cli.Context, s store.Store) error {
	records, err := s.GetRedirects()

	if err != nil {
		return err
	}

	if c.Bool("json") {
		if c.String("filter") != "" {
			os.Stderr.WriteString("Filters are ignored while printing JSON!\n")
		}

		res, err := json.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
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

	for _, record := range records {
		if c.String("filter") != "" && !record.Contains(c.String("filter")) {
			continue
		}

		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}
