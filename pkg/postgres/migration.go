package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	packr "github.com/gobuffalo/packr/v2"
	"github.com/kintuda/tech-challenge-pismo/pkg/config"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	migrationDriver = "postgres"
)

type Migrator struct {
	Db         *sql.DB
	Cfg        *config.MigrationConfig
	Migrations *migrate.PackrMigrationSource
}

func NewMigrator(cfg *config.MigrationConfig) (*Migrator, error) {
	db, err := sql.Open(migrationDriver, cfg.PostgresDns)

	if err != nil {
		return nil, err
	}

	migrations := &migrate.PackrMigrationSource{
		Box: packr.New(migrationDriver, "./migrations"),
	}

	return &Migrator{
		Cfg:        cfg,
		Migrations: migrations,
		Db:         db,
	}, nil
}

func (m *Migrator) Up() (int, error) {
	n, err := migrate.Exec(m.Db, migrationDriver, m.Migrations, migrate.Up)

	if err != nil {
		return 0, err
	}

	return n, nil
}

func (m *Migrator) Down() (int, error) {
	n, err := migrate.Exec(m.Db, migrationDriver, m.Migrations, migrate.Down)

	if err != nil {
		return 0, err
	}

	return n, nil
}

func (m *Migrator) CreateFile(name string) error {
	var box = packr.New("migrations", "./migrations")

	if _, err := os.Stat(box.Path); os.IsNotExist(err) {
		return err
	}

	var templateContent = `
-- +migrate Up
-- +migrate Down
		`
	var tpl *template.Template = template.Must(template.New("new_migration").Parse(templateContent))
	fileName := fmt.Sprintf("%s-%s.sql", time.Now().Format("20060102150405"), strings.TrimSpace(name))
	pathName := path.Join(box.Path, fileName)
	f, err := os.Create(pathName)

	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	if err := tpl.Execute(f, nil); err != nil {
		return err
	}

	return nil
}
