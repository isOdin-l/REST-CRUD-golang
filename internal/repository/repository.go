package repository

import (
	"context"
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
	QueryRow(ctx context.Context, recieveObject interface{}, sql string, args ...interface{}) error
	Close()
}

func NewRepository(db IDatabase, sqlBuilder ISql) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db, sqlBuilder),
		ListRepository: NewListRepository(db, sqlBuilder),
		ItemRepository: NewItemRepository(db, sqlBuilder),
	}
}
