package database

import (
	"database/sql"
	"exchange/internal/config"
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(log *slog.Logger, cfg *config.Config) error {
	connStr := getConnectionString(cfg)
	log.Info("connecting with connection string:")
	log.Debug("connection string:", slog.String("connStr", connStr))

	var err error
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed ping check: %w", err)
	}

	log.Info("successfully connected to database")
	return nil
}

func getConnectionString(cfg *config.Config) string {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DbUser.Value,
		cfg.DbPass.Value,
		cfg.DbHost.Value,
		cfg.DbPort.Value,
		cfg.DbName.Value,
	)

	return connStr
}
