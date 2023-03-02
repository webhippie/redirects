package command

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show redirect pattern",
		Run:   showAction,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("missing pattern argument")
			}

			return nil
		},
	}
)

// tmplShow represents a specific redirect detail view.
var tmplShow = "ID: \x1b[33m{{ .ID }}\x1b[0m" + `
Source: {{ .Source }}
Destination: {{ .Destination }}
Priority: {{ .Priority }}
`

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().String("format", tmplShow, "Custom output format")
	viper.BindPFlag("show.format", showCmd.Flags().Lookup("format"))

	showCmd.Flags().Bool("json", false, "Print in JSON format")
	viper.BindPFlag("show.json", showCmd.Flags().Lookup("json"))
}

func showAction(_ *cobra.Command, args []string) {
	ctx := context.Background()

	record, err := storage.GetRedirect(
		ctx,
		args[0],
	)

	if err != nil {
		cobra.CheckErr(fmt.Errorf("failed to find pattern: %w", err))
	}

	if viper.GetBool("show.json") {
		res, err := json.MarshalIndent(record, "", "  ")

		if err != nil {
			cobra.CheckErr(fmt.Errorf("failed to parse pattern: %w", err))
		}

		fmt.Println(string(res))
		return
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Parse(
		fmt.Sprintln(viper.GetString("show.format")),
	)

	if err != nil {
		cobra.CheckErr(fmt.Errorf("failed to parse format: %w", err))
	}

	if err := tmpl.Execute(os.Stdout, record); err != nil {
		cobra.CheckErr(fmt.Errorf("failed to render pattern: %w", err))
	}
}
