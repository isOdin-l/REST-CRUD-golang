package repository

import (
	"context"

	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/errors"
	mapper "isOdin/RestApi/internal/repository/models"

	"github.com/google/uuid"
)

type IListSql interface {
	InsertList(values ...any) (string, []interface{}, error)
	SelectList(listId, userId uuid.UUID) (string, []interface{}, error)
	UpdateList(author_id, listId uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error)
	DeleteList(listId, userId uuid.UUID) (string, []interface{}, error)
}

type ListRepository struct {
	db         IDatabase
	sqlBuilder IListSql
}

func NewListRepository(db IDatabase, sqlBuilder IListSql) *ListRepository {
	return &ListRepository{db: db, sqlBuilder: sqlBuilder}
}

func (r *ListRepository) CreateList(ctx context.Context, list *entities.List) *errors.AppError {
	listDb := mapper.FromListEntityToRepo(list)

	query, value, err := r.sqlBuilder.InsertList(listDb.Id, listDb.Author_id, listDb.Title, listDb.Description)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if err := r.db.Exec(ctx, query, value...); err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}

func (r *ListRepository) GetList(ctx context.Context, list *entities.List) (*entities.List, *errors.AppError) {
	listDb := mapper.FromListEntityToRepo(list)

	query, value, err := r.sqlBuilder.SelectList(list.ListId, listDb.Author_id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if err := r.db.QueryRow(ctx, listDb, query, value...); err != nil {
		if err.Error() == "no rows in result set" {
			return nil, errors.ErrNotFound
		}
		return nil, errors.NewInternalError(err)
	}

	return listDb.ToEntity(), nil
}

func (r *ListRepository) DeleteList(ctx context.Context, list *entities.List) *errors.AppError {
	listDb := mapper.FromListEntityToRepo(list)

	query, value, err := r.sqlBuilder.DeleteList(listDb.Id, listDb.Author_id)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if err := r.db.Exec(ctx, query, value...); err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}

func (r *ListRepository) UpdateList(ctx context.Context, list *entities.UpdateList, updateInfo map[string]interface{}) (*entities.List, *errors.AppError) {
	listDb := mapper.FromUpdateListEntityToRepo(list)

	query, value, err := r.sqlBuilder.UpdateList(listDb.Author_id, listDb.Id, updateInfo)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if err := r.db.QueryRow(ctx, listDb, query, value...); err != nil {
		if err.Error() == "no rows in result set" {
			return nil, errors.ErrNotFound
		}
		return nil, errors.NewInternalError(err)
	}

	return listDb.ToEntity(), nil
}
