package config

type MigrationConfig struct {
	PostgresDns string `env:"POSTGRES_DNS"`
}
