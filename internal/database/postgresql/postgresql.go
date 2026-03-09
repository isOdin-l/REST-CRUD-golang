package postgresql

import (
	"context"

	"isOdin/RestApi/configs"
	"isOdin/RestApi/internal/helpers"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IPostgresExecutor interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

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

	return &PostgresDB{conn}, nil
}

func (ps *PostgresDB) Exec(ctx context.Context, sql string, values ...interface{}) error {
	_, err := ps.getExecutor(ctx).Exec(ctx, sql, values...)
	return err
}

func (ps *PostgresDB) QueryRow(ctx context.Context, sql string, values ...interface{}) pgx.Row {
	return ps.getExecutor(ctx).QueryRow(ctx, sql, values...)
}

func (ps *PostgresDB) Scan(row pgx.Row, dest ...any) error{
	return row.Scan(dest...)
}

func (ps *PostgresDB) Close() {
	ps.conn.Close()
}

func (ps *PostgresDB) getExecutor(ctx context.Context) IPostgresExecutor {
	tx, ok := ctx.Value(helpers.TXKEY{}).(pgx.Tx)
	if !ok {
		return ps.conn
	}
	return tx
}

func (ps *PostgresDB) WithinTx(ctx context.Context, fn func(ctx context.Context) (*any, error)) (*any, error) {
	if _, ok := ctx.Value(helpers.TXKEY{}).(pgx.Tx); ok {
		return fn(ctx)
	}

	tx, errTx := ps.conn.BeginTx(ctx, pgx.TxOptions{})
	if errTx != nil {
		return nil, errTx
	}

	defer tx.Rollback(ctx)
	res, errFn := fn(ctx)
	if errFn != nil {
		return nil, errFn
	}

	return res, tx.Commit(ctx)

}
