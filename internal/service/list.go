package service

import (
	"context"
	"reflect"
	"strings"

	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/errors"

	"github.com/google/uuid"
)

type ListRepoInterface interface {
	CreateList(ctx context.Context, list *entities.List) *errors.AppError
	GetList(ctx context.Context, listId uuid.UUID) (*entities.List, *errors.AppError)
	UpdateList(ctx context.Context, listId uuid.UUID, updateInfo map[string]interface{}) (*entities.List, *errors.AppError)
	DeleteList(ctx context.Context, listId uuid.UUID) *errors.AppError
}

type TodoListService struct {
	repo ListRepoInterface
}

func NewTodoListService(repo ListRepoInterface) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateList(ctx context.Context, list *entities.List) (*entities.List, *errors.AppError) {
	var err error
	list.ListId, err = uuid.NewV7()
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	errRepo := s.repo.CreateList(ctx, list)
	return list, errRepo
}

func (s *TodoListService) GetListById(ctx context.Context, list *entities.List) (*entities.List, *errors.AppError) {
	return s.repo.GetList(ctx, list.ListId)
}

func (s *TodoListService) DeleteList(ctx context.Context, list *entities.List) *errors.AppError {
	return s.repo.DeleteList(ctx, list.ListId)
}

func (s *TodoListService) UpdateList(ctx context.Context, list *entities.List) (*entities.List, *errors.AppError) {
	updateInfo := make(map[string]interface{})
	v := reflect.ValueOf(*list)
	t := reflect.TypeOf(*list)

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i)

		if fieldName == "UserId" || fieldName == "ListId" {
			continue
		}

		if !fieldValue.IsZero() {
			dbColumnName := strings.ToLower(fieldName)
			updateInfo[dbColumnName] = fieldValue.Interface()
		}
	}

	if len(updateInfo) == 0 {
		return s.repo.GetList(ctx, list.ListId)
	}

	return s.repo.UpdateList(ctx, list.ListId, updateInfo)
}
