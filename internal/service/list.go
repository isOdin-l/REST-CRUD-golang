package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	repoReqDTO "github.com/isOdin/RestApi/internal/repository/requestDTO"
	repoResDTO "github.com/isOdin/RestApi/internal/repository/responseDTO"
	"github.com/isOdin/RestApi/internal/service/requestDTO"
	"github.com/isOdin/RestApi/internal/service/responseDTO"
)

type ListRepoInterface interface {
	CreateList(ctx context.Context, listInfo *repoReqDTO.CreateList) (uuid.UUID, error)
	GetAllLists(ctx context.Context, userId uuid.UUID) (*[]repoResDTO.GetList, error)
	GetListById(ctx context.Context, listInfo *repoReqDTO.GetListById) (*repoResDTO.GetListById, error)
	DeleteList(ctx context.Context, listInfo *repoReqDTO.DeleteList) error
	UpdateList(ctx context.Context, listInfo *repoReqDTO.UpdateList) error
}

type TodoListService struct {
	repo ListRepoInterface
}

func NewTodoListService(repo ListRepoInterface) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateList(ctx context.Context, listInfo *requestDTO.CreateList) (uuid.UUID, error) {
	return s.repo.CreateList(ctx, listInfo.ConvertToRepoModel())
}

func (s *TodoListService) GetAllLists(ctx context.Context, userId uuid.UUID) (*[]responseDTO.GetList, error) {
	listsResponsed, err := s.repo.GetAllLists(ctx, userId)
	if err != nil {
		return nil, err
	}

	lists := make([]responseDTO.GetList, len(*listsResponsed))
	for i := range len(*listsResponsed) {
		// ------- Указатель на массив -> массив -> элемент массива -> перевод элемента в указатель на другой тип -> элемент другого типа -------
		lists[i] = *((*listsResponsed)[i].ToServiceModel())
	}

	return &lists, nil
}

func (s *TodoListService) GetListById(ctx context.Context, listInfo *requestDTO.GetListById) (*responseDTO.GetListById, error) {
	listResponsed, err := s.repo.GetListById(ctx, listInfo.ConvertToRepoModel())
	if err != nil {
		return nil, err
	}

	return listResponsed.ToServiceModel(), nil
}

func (s *TodoListService) DeleteList(ctx context.Context, listInfo *requestDTO.DeleteList) error {
	return s.repo.DeleteList(ctx, listInfo.ConvertToRepoModel())
}

func (s *TodoListService) UpdateList(ctx context.Context, listInfo *requestDTO.UpdateList) error {
	setValues := make([]string, 0)
	setArgs := make([]interface{}, 0)
	argId := 1

	if listInfo.Title == "" && listInfo.Description == "" {
		return errors.New("Update structure has no values")
	}

	if listInfo.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		setArgs = append(setArgs, listInfo.Title)
		argId++
	}

	if listInfo.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		setArgs = append(setArgs, listInfo.Description)
		argId++
	}

	setValuesQuery := strings.Join(setValues, ", ")
	setArgs = append(setArgs, listInfo.ListId, listInfo.UserId)

	return s.repo.UpdateList(ctx, listInfo.ConvertToRepoModel(&setArgs, argId, setValuesQuery))
}
