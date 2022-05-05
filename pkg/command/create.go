package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webhippie/redirects/pkg/model"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create redirect pattern",
		Run:   createAction,
		Args:  cobra.NoArgs,
	}
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().String("source", "", "Source for the redirect")
	viper.BindPFlag("create.source", createCmd.Flags().Lookup("source"))

	createCmd.Flags().String("destination", "", "Destination for the redirect")
	viper.BindPFlag("create.destination", createCmd.Flags().Lookup("destination"))

	createCmd.Flags().Int("priority", 0, "Priority for the redirect")
	viper.BindPFlag("create.priority", createCmd.Flags().Lookup("priority"))
}

func createAction(ccmd *cobra.Command, args []string) {
	record := &model.Redirect{}

	if val := viper.GetString("create.source"); viper.IsSet("create.source") && val != "" {
		record.Source = val
	} else {
		cobra.CheckErr(fmt.Errorf("you must provide a source"))
	}

	if val := viper.GetString("create.destination"); viper.IsSet("create.destination") && val != "" {
		record.Destination = val
	} else {
		cobra.CheckErr(fmt.Errorf("you must provide a destination"))
	}

	if val := viper.GetInt("create.priority"); viper.IsSet("create.priority") {
		record.Priority = val
	}

	if err := storage.CreateRedirect(record); err != nil {
		cobra.CheckErr(fmt.Errorf("failed to create pattern: %w", err))
	}
}
