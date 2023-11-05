package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kintuda/tech-challenge-pismo/pkg/config"
)

type Pool struct {
	Conn *pgxpool.Pool
}

func NewConnectionPool(app *config.ServerConfig) (*Pool, error) {
	config, err := pgxpool.ParseConfig(app.PostgresDns)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config: %v\n", err)
		os.Exit(1)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Pool{Conn: conn}, nil
}

func (d *Pool) CloseConnection() {
	d.Conn.Close()
}
