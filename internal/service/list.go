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
	GetList(ctx context.Context, list *entities.List) (*entities.List, *errors.AppError)
	UpdateList(ctx context.Context, list *entities.UpdateList, updateInfo map[string]interface{}) (*entities.List, *errors.AppError)
	DeleteList(ctx context.Context, list *entities.List) *errors.AppError
}

type TodoListService struct {
	repo ListRepoInterface
	txMn ITransactionManager
}

func NewTodoListService(repo ListRepoInterface, txMn ITransactionManager) *TodoListService {
	return &TodoListService{repo: repo, txMn: txMn}
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
	return s.repo.GetList(ctx, list)
}

func (s *TodoListService) DeleteList(ctx context.Context, list *entities.List) *errors.AppError {
	return s.repo.DeleteList(ctx, list)
}

func (s *TodoListService) UpdateList(ctx context.Context, list *entities.UpdateList) (*entities.List, *errors.AppError) {
	updateInfo := make(map[string]interface{})
	k := reflect.TypeOf(*list)
	v := reflect.ValueOf(*list)

	for i := 0; i < v.NumField(); i++ {
		fieldName := k.Field(i).Name
		fieldValue := v.Field(i)

		if fieldName == "UserId" || fieldName == "ListId" {
			continue
		}

		if !fieldValue.IsNil() {
			dbColumnName := strings.ToLower(fieldName)
			updateInfo[dbColumnName] = fieldValue.Interface()
		}
	}

	if len(updateInfo) == 0 {
		return s.repo.GetList(ctx, &entities.List{
			UserId:      list.UserId,
			ListId:      list.ListId,
			Title:       *list.OptValues.Title,
			Description: *list.OptValues.Description,
		})
	}

	return s.repo.UpdateList(ctx, list, updateInfo)
}
