package command

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update redirect pattern",
		Run:   updateAction,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("missing pattern argument")
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().String("source", "", "Source for the redirect")
	viper.BindPFlag("update.source", updateCmd.Flags().Lookup("source"))

	updateCmd.Flags().String("destination", "", "Destination for the redirect")
	viper.BindPFlag("update.destination", updateCmd.Flags().Lookup("destination"))

	updateCmd.Flags().Int("priority", 0, "Priority for the redirect")
	viper.BindPFlag("update.priority", updateCmd.Flags().Lookup("priority"))
}

func updateAction(ccmd *cobra.Command, args []string) {
	ctx := context.Background()

	record, err := storage.GetRedirect(
		ctx,
		args[0],
	)

	if err != nil {
		cobra.CheckErr(fmt.Errorf("failed to find pattern: %w", err))
	}

	changed := false

	if val := viper.GetString("update.source"); viper.IsSet("update.source") && val != record.Source {
		record.Source = val
		changed = true
	}

	if val := viper.GetString("update.destination"); viper.IsSet("update.destination") && val != record.Destination {
		record.Destination = val
		changed = true
	}

	if val := viper.GetInt("update.priority"); viper.IsSet("update.priority") && val != record.Priority {
		record.Priority = val
		changed = true
	}

	if changed {
		if err := storage.UpdateRedirect(ctx, record); err != nil {
			cobra.CheckErr(fmt.Errorf("failed to update pattern: %w", err))
		}
	}
}
