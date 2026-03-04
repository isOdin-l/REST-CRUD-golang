package postgresql

import (
	"context"

	"isOdin/RestApi/configs"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	conn *pgxpool.Pool
}

func NewPostgresDB(cfg *configs.Config) (*PostgresDB, error) {
	conn, err := pgxpool.New(context.Background(), cfg.DSN())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &PostgresDB{conn: conn}, nil
}

func (ps *PostgresDB) Exec(ctx context.Context, sql string, values ...interface{}) error {
	_, err := ps.conn.Exec(ctx, sql, values...)
	return err
}

func (ps *PostgresDB) QueryRow(ctx context.Context, recieveObject interface{}, sql string, values ...interface{}) error {
	return ps.conn.QueryRow(ctx, sql, values...).Scan(&recieveObject)
}

func (ps *PostgresDB) Close() {
	ps.conn.Close()
}
