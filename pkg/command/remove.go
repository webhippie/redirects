package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:   "remove <pattern>",
		Short: "Remove redirect pattern",
		Run:   removeAction,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("missing pattern argument")
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

func removeAction(ccmd *cobra.Command, args []string) {
	record, err := storage.GetRedirect(
		args[0],
	)

	if err != nil {
		cobra.CheckErr(fmt.Errorf("failed to find pattern: %w", err))
	}

	if err := storage.DeleteRedirect(record.ID); err != nil {
		cobra.CheckErr(fmt.Errorf("failed to remove pattern: %w", err))
	}
}
