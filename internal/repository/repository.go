package repository

import (
	"context"
)

type Repository struct {
	*AuthRepository
	*ListRepository
	*ItemRepository
}

type IDatabase interface {
	Exec(ctx context.Context, sql string, args ...interface{}) error
	QueryRow(ctx context.Context, recieveObject interface{}, sql string, args ...interface{}) error
	Close()
}

func NewRepository(db IDatabase) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
		ListRepository: NewListRepository(db),
		ItemRepository: NewItemRepository(db),
	}
}
