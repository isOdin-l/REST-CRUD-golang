package repository

import (
	"context"

	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/errors"
	"isOdin/RestApi/internal/repository/models"

	"github.com/google/uuid"
)

type IItemSql interface {
	InsertItem(values ...any) (string, []interface{}, error)
	SelectItem(itemId uuid.UUID) (string, []interface{}, error)
	UpdateItem(itemId uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error)
	DeleteItem(itemId uuid.UUID) (string, []interface{}, error)
}

type ItemRepository struct {
	db         IDatabase
	sqlBuilder IItemSql
}

func NewItemRepository(db IDatabase, sqlBuilder IItemSql) *ItemRepository {
	return &ItemRepository{db: db, sqlBuilder: sqlBuilder}
}

func (r *ItemRepository) CreateItem(ctx context.Context, item *entities.Item) *errors.AppError {
	itemDb := models.FromItemEntityToRepo(item)

	query, values, err := r.sqlBuilder.InsertItem(itemDb.Id, itemDb.List_id, itemDb.Title, itemDb.Description)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if err := r.db.Exec(ctx, query, values...); err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}

func (r *ItemRepository) GetItem(ctx context.Context, itemId uuid.UUID) (*entities.Item, *errors.AppError) {
	query, value, errBuilder := r.sqlBuilder.SelectItem(itemId)
	if errBuilder != nil {
		return nil, errors.NewInternalError(errBuilder)
	}

	item := &models.Item{}
	if errQuery := r.db.QueryRow(ctx, item, query, value...); errQuery != nil {
		if errQuery.Error() == "no rows in result set" {
			return nil, errors.ErrNotFound
		}
		return nil, errors.NewInternalError(errQuery)
	}

	return item.ToEntity(), nil
}
func (r *ItemRepository) DeleteItem(ctx context.Context, itemId uuid.UUID) *errors.AppError {
	query, value, err := r.sqlBuilder.DeleteItem(itemId)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if err := r.db.Exec(ctx, query, value...); err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}

func (r *ItemRepository) UpdateItem(ctx context.Context, itemId uuid.UUID, updateInfo map[string]interface{}) (*entities.Item, *errors.AppError) {
	query, values, err := r.sqlBuilder.UpdateItem(itemId, updateInfo)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	item := &models.Item{}
	if errDbQuery := r.db.QueryRow(ctx, item, query, values...); errDbQuery != nil {
		if errDbQuery.Error() == "no rows in result set" {
			return nil, errors.ErrNotFound
		}
		return nil, errors.NewInternalError(errDbQuery)
	}

	return item.ToEntity(), nil
}
