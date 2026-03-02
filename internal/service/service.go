package service

import (
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

func NewService(cfg *configs.InternalConfig, repo RepositoryInterface) *Service {
	return &Service{
		AuthService:     NewAuthService(cfg, repo),
		TodoListService: NewTodoListService(repo),
		TodoItemService: NewTodoItemService(repo),
	}
}
