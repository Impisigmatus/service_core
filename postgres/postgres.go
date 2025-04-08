package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog"
)

type Config struct {
	Logger zerolog.Logger

	Hostname string
	Port     uint64
	Database string
	User     string
	Password string
}

func NewPostgres(cfg Config) *sql.DB {
	const driver = "pgx"

	pattern := fmt.Sprintf(
		"host=%s port=%d database=%s user=%s password=%s sslmode=disable",
		cfg.Hostname, cfg.Port, cfg.Database, cfg.User, cfg.Password,
	)

	config, err := pgx.ParseConfig(pattern)
	if err != nil {
		cfg.Logger.Panic().Msgf("Invalid postgres config: %s", err)
	}

	connection := stdlib.RegisterConnConfig(config)
	db, err := sql.Open(driver, connection)
	if err != nil {
		cfg.Logger.Panic().Msgf("Invalid postgres connect: %s", err)
	}

	if err := db.Ping(); err != nil {
		cfg.Logger.Panic().Msgf("Invalid postgres ping: %s", err)
	}

	return db
}
