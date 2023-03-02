package command

import (
	"context"
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

func removeAction(_ *cobra.Command, args []string) {
	ctx := context.Background()

	record, err := storage.GetRedirect(
		ctx,
		args[0],
	)

	if err != nil {
		cobra.CheckErr(fmt.Errorf("failed to find pattern: %w", err))
	}

	if err := storage.DeleteRedirect(ctx, record.ID); err != nil {
		cobra.CheckErr(fmt.Errorf("failed to remove pattern: %w", err))
	}
}
