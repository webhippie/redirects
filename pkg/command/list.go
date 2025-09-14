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
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List redirect patterns",
		Run:   listAction,
		Args:  cobra.NoArgs,
	}
)

// tmplList represents a row within redirect listing.
var tmplList = "ID: \x1b[33m{{ .ID }}\x1b[0m" + `
Source: {{ .Source }}
Destination: {{ .Destination }}
`

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().String("filter", "", "Filter output by needle")
	_ = viper.BindPFlag("list.filter", listCmd.Flags().Lookup("filter"))

	listCmd.Flags().String("format", tmplList, "Custom output format")
	_ = viper.BindPFlag("list.format", listCmd.Flags().Lookup("format"))

	listCmd.Flags().Bool("json", false, "Print in JSON format")
	_ = viper.BindPFlag("list.json", listCmd.Flags().Lookup("json"))
}

func listAction(_ *cobra.Command, _ []string) {
	ctx := context.Background()
	records, err := storage.GetRedirects(ctx)

	if err != nil {
		cobra.CheckErr(fmt.Errorf("failed to find patterns: %w", err))
	}

	if viper.GetBool("list.json") {
		if viper.GetString("list.filter") != "" {
			_, _ = os.Stderr.WriteString("Filters are ignored while printing JSON!\n")
		}

		res, err := json.MarshalIndent(records, "", "  ")

		if err != nil {
			cobra.CheckErr(fmt.Errorf("failed to parse patterns: %w", err))
		}

		fmt.Println(string(res))
		return
	}

	if len(records) == 0 {
		_, _ = os.Stderr.WriteString("No patterns found\n")
		return
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Parse(
		fmt.Sprintln(viper.GetString("list.format")),
	)

	if err != nil {
		cobra.CheckErr(fmt.Errorf("failed to parse format: %w", err))
	}

	for _, record := range records {
		if viper.GetString("list.filter") != "" && !record.Contains(viper.GetString("list.filter")) {
			continue
		}

		if err := tmpl.Execute(os.Stdout, record); err != nil {
			cobra.CheckErr(fmt.Errorf("failed to render pattern: %w", err))
		}
	}
}
