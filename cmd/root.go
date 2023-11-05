package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	ctx := context.Background()

	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "tech-challenge-pismo cli",
	}

	rootCmd.AddCommand(NewServerCmd(ctx))
	rootCmd.AddCommand(NewMigrationCmd())
	rootCmd.AddCommand(NewMigrationCmd())

	return rootCmd
}
