package service

import (
	"context"
	"reflect"

	"isOdin/RestApi/internal/entities"

	"github.com/google/uuid"
)

type ListRepoInterface interface {
	CreateList(ctx context.Context, list *entities.List) error
	GetList(ctx context.Context, listId uuid.UUID) (*entities.List, error)
	UpdateList(ctx context.Context, listId uuid.UUID, updateInfo map[string]interface{}) (*entities.List, error)
	DeleteList(ctx context.Context, listId uuid.UUID) error
}

type TodoListService struct {
	repo ListRepoInterface
}

func NewTodoListService(repo ListRepoInterface) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateList(ctx context.Context, list *entities.List) (*entities.List, error) {
	var err error
	list.ListId, err = uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return list, s.repo.CreateList(ctx, list)
}

func (s *TodoListService) GetListById(ctx context.Context, list *entities.List) (*entities.List, error) {
	return s.repo.GetList(ctx, list.ListId)
}

func (s *TodoListService) DeleteList(ctx context.Context, list *entities.List) error {
	return s.repo.DeleteList(ctx, list.ListId)
}

func (s *TodoListService) UpdateList(ctx context.Context, list *entities.List) (*entities.List, error) {
	k := reflect.TypeOf(*list)
	v := reflect.ValueOf(*list)
	updateInfo := make(map[string]interface{})

	for i := 0; i < k.NumField(); i++ {
		if !v.IsNil() {
			updateInfo[k.Field(i).Name] = v.Field(i)
		}
	}
	return s.repo.UpdateList(ctx, list.ListId, updateInfo)
}
