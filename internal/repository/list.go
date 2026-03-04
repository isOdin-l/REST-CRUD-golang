package repository

import (
	"context"

	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/repository/models"

	"github.com/google/uuid"
)

type IListSql interface {
	InsertList(values ...any) (string, []interface{}, error)
	SelectList(listId uuid.UUID) (string, []interface{}, error)
	UpdateList(author_id uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error)
	DeleteList(listId uuid.UUID) (string, []interface{}, error)
}

type ListRepository struct {
	db         IDatabase
	sqlBuilder IListSql
}

func NewListRepository(db IDatabase, sqlBuilder IListSql) *ListRepository {
	return &ListRepository{db: db, sqlBuilder: sqlBuilder}
}

func (r *ListRepository) CreateList(ctx context.Context, list *entities.List) error {
	listDb := models.FromListEntityToRepo(list)
	query, value, err := r.sqlBuilder.InsertList(listDb.Id, listDb.Author_id, listDb.Title, listDb.Description)
	if err != nil {
		return err
	}

	return r.db.Exec(ctx, query, value...)
}

func (r *ListRepository) GetList(ctx context.Context, listId uuid.UUID) (*entities.List, error) {
	query, value, err := r.sqlBuilder.SelectList(listId)
	if err != nil {
		return nil, err
	}

	list := &models.List{}
	err = r.db.QueryRow(ctx, list, query, value...)

	return list.ToEntity(), err
}

func (r *ListRepository) DeleteList(ctx context.Context, listId uuid.UUID) error {
	query, value, err := r.sqlBuilder.DeleteList(listId)
	if err != nil {
		return err
	}
	return r.db.Exec(ctx, query, value...)
}

func (r *ListRepository) UpdateList(ctx context.Context, listId uuid.UUID, updateInfo map[string]interface{}) (*entities.List, error) {
	query, value, err := r.sqlBuilder.UpdateList(listId, updateInfo)
	if err != nil {
		return nil, err
	}

	list := &models.List{}
	errDbQuery := r.db.QueryRow(ctx, list, query, value...)

	return list.ToEntity(), errDbQuery
}
