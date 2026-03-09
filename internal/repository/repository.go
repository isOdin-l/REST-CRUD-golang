package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	*AuthRepository
	*ListRepository
	*ItemRepository
}

type ISql interface {
	IAuthSql
	IListSql
	IItemSql
}

type IDatabase interface {
	Exec(ctx context.Context, sql string, args ...interface{}) error
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Scan(row pgx.Row, dest ...any) error
	Close()
}

func NewRepository(db IDatabase, sqlBuilder ISql) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db, sqlBuilder),
		ListRepository: NewListRepository(db, sqlBuilder),
		ItemRepository: NewItemRepository(db, sqlBuilder),
	}
}
