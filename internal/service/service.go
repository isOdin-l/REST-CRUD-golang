package service

import (
	"context"
	"isOdin/RestApi/configs"
)

type Service struct {
	*AuthService
	*TodoListService
	*TodoItemService
}

type RepositoryInterface interface {
	AuthRepoInterface
	ListRepoInterface
	ItemRepoInterface
}

type ITransactionManager interface {
	WithinTx(ctx context.Context, fn func(context.Context) (*any, error)) (*any, error)
}

func NewService(cfg *configs.InternalConfig, repo RepositoryInterface, txMn ITransactionManager) *Service {
	return &Service{
		AuthService:     NewAuthService(cfg, repo, txMn),
		TodoListService: NewTodoListService(repo, txMn),
		TodoItemService: NewTodoItemService(repo, txMn),
	}
}
