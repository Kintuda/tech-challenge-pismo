package cmd

import (
	"strings"

	"github.com/kintuda/tech-challenge-pismo/pkg/config"
	"github.com/kintuda/tech-challenge-pismo/pkg/postgres"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var migrationDir string

func NewMigrationCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "migration",
		Short: "Migration CLI",
	}

	migrateUp := &cobra.Command{
		Use:  "up",
		RunE: migrationUp,
	}

	migrateDown := &cobra.Command{
		Use:  "down",
		RunE: migrationDown,
	}

	createMigrationFile := &cobra.Command{
		RunE: createMigration,
		Use:  "create",
		Args: cobra.MinimumNArgs(1),
	}

	rootCmd.PersistentFlags().StringVarP(&migrationDir, "migration-dir", "m", "", "directory that holds the migration files")
	rootCmd.AddCommand(createMigrationFile)
	rootCmd.AddCommand(migrateDown)
	rootCmd.AddCommand(migrateUp)

	return rootCmd
}

func migrationDown(cmd *cobra.Command, args []string) error {
	cfg := new(config.MigrationConfig)
	err := config.LoadConfigFromEnv(cfg)

	if err != nil {
		return err
	}

	m, err := postgres.NewMigrator(cfg)

	if err != nil {
		return err
	}

	n, err := m.Down()

	if err != nil {
		return err
	}

	log.Info().Msgf("Successful migrate, %v resources", n)
	return nil
}

func migrationUp(cmd *cobra.Command, args []string) error {
	cfg := new(config.MigrationConfig)
	err := config.LoadConfigFromEnv(cfg)

	if err != nil {
		return err
	}

	m, err := postgres.NewMigrator(cfg)

	if err != nil {
		return err
	}

	n, err := m.Up()

	if err != nil {
		return err
	}

	log.Info().Msgf("Successful migrate, %v resources", n)
	return nil
}

func createMigration(cmd *cobra.Command, args []string) error {
	var migrationName = strings.Join(args, "_")
	cfg := new(config.MigrationConfig)
	err := config.LoadConfigFromEnv(cfg)

	if err != nil {
		return err
	}

	m, err := postgres.NewMigrator(cfg)

	if err != nil {
		return err
	}

	return m.CreateFile(migrationName)
}
