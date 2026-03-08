package service

import (
	"context"
	"reflect"
	"strings"

	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/errors"

	"github.com/google/uuid"
)

type ItemRepoInterface interface {
	CreateItem(ctx context.Context, item *entities.Item) *errors.AppError
	GetItem(ctx context.Context, itemId uuid.UUID) (*entities.Item, *errors.AppError)
	UpdateItem(ctx context.Context, itemId uuid.UUID, updateInfo map[string]interface{}) (*entities.Item, *errors.AppError)
	DeleteItem(ctx context.Context, itemId uuid.UUID) *errors.AppError
}

type TodoItemService struct {
	repo ItemRepoInterface
	txMn ITransactionManager
}

func NewTodoItemService(repo ItemRepoInterface, txMn ITransactionManager) *TodoItemService {
	return &TodoItemService{repo: repo, txMn: txMn}
}

func (s *TodoItemService) CreateItem(ctx context.Context, item *entities.Item) (*entities.Item, *errors.AppError) {
	var err error
	item.ItemId, err = uuid.NewV7()
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	errRepo := s.repo.CreateItem(ctx, item)
	return item, errRepo
}

func (s *TodoItemService) GetItem(ctx context.Context, item *entities.Item) (*entities.Item, *errors.AppError) {
	return s.repo.GetItem(ctx, item.ItemId)
}

func (s *TodoItemService) DeleteItem(ctx context.Context, item *entities.Item) *errors.AppError {
	return s.repo.DeleteItem(ctx, item.ItemId)
}

func (s *TodoItemService) UpdateItem(ctx context.Context, item *entities.UpdateItem) (*entities.Item, *errors.AppError) {
	updateInfo := make(map[string]interface{})
	k := reflect.TypeOf(item.OptValues)
	v := reflect.ValueOf(item.OptValues)

	for i := 0; i < v.NumField(); i++ {
		fieldName := k.Field(i).Name
		fieldValue := v.Field(i)

		if !fieldValue.IsNil() {
			dbColumnName := strings.ToLower(fieldName)
			updateInfo[dbColumnName] = fieldValue.Interface()
		}
	}

	if len(updateInfo) == 0 {
		return s.repo.GetItem(ctx, item.ItemId)
	}

	return s.repo.UpdateItem(ctx, item.ItemId, updateInfo)
}
