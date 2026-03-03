package service

import (
	"context"
	"reflect"

	"isOdin/RestApi/internal/entities"

	"github.com/google/uuid"
)

type ItemRepoInterface interface {
	CreateItem(ctx context.Context, item *entities.Item) error
	GetItem(ctx context.Context, itemId uuid.UUID) (*entities.Item, error)
	UpdateItem(ctx context.Context, itemId uuid.UUID, updateInfo map[string]interface{}) (*entities.Item, error)
	DeleteItem(ctx context.Context, itemId uuid.UUID) error
}

type TodoItemService struct {
	repo ItemRepoInterface
}

func NewTodoItemService(repo ItemRepoInterface) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) CreateItem(ctx context.Context, item *entities.Item) (uuid.UUID, error) {
	var err error
	item.ItemId, err = uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}
	return item.ItemId, s.repo.CreateItem(ctx, item)
}

func (s *TodoItemService) GetItem(ctx context.Context, item *entities.Item) (*entities.Item, error) {
	return s.repo.GetItem(ctx, item.ItemId)
}

func (s *TodoItemService) DeleteItem(ctx context.Context, item *entities.Item) error {
	return s.repo.DeleteItem(ctx, item.ItemId)
}

func (s *TodoItemService) UpdateItem(ctx context.Context, item *entities.Item) (*entities.Item, error) {
	k := reflect.TypeOf(*item)
	v := reflect.ValueOf(*item)
	updateInfo := make(map[string]interface{})

	for i := 0; i < k.NumField(); i++ {
		if !v.IsNil() {
			updateInfo[k.Field(i).Name] = v.Field(i)
		}
	}

	return s.repo.UpdateItem(ctx, item.ItemId, updateInfo)
}
