package database

import (
	"exchange/internal/config"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(log *slog.Logger, cfg *config.Config) error {
	driver, err := mysql.WithInstance(DB, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		cfg.DbName.Value,
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Info("migrations applied successfully")
	return nil
}
