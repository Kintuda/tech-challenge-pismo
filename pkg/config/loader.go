package config

import (
	"os"

	"github.com/caarlos0/env"
	gotdotenv "github.com/joho/godotenv"
)

type Configs interface {
	ServerConfig |
		MigrationConfig
}

var (
	PROD_ENV = "production"
)

func LoadAppConfig[k Configs]() (*k, error) {
	var config k

	if os.Getenv("ENV") != PROD_ENV {
		if err := gotdotenv.Load(); err != nil {
			return &config, err
		}
	}

	if err := env.Parse(&config); err != nil {
		return &config, err
	}

	return &config, nil
}

func LoadConfigFromEnv[K Configs](c *K) error {
	if os.Getenv("ENV") != PROD_ENV {
		if err := gotdotenv.Load(); err != nil {
			return err
		}
	}

	if err := env.Parse(c); err != nil {
		return err
	}

	return nil
}
