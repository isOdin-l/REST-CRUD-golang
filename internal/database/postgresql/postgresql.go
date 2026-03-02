package postgresql

import (
	"context"

	"isOdin/RestApi/configs"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(cfg *configs.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), cfg.DSN())

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}
