package postgresql

import (
	"context"

	"github.com/isOdin/RestApi/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewPostgresDB(cfg *configs.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), cfg.DSN())

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	logrus.Info("Database connected")
	return conn, nil
}
