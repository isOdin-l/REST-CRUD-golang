package repository

import (
	"context"

	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/errors"
	"isOdin/RestApi/internal/repository/models"

	"github.com/google/uuid"
)

type IItemSql interface {
	InsertItem(userId uuid.UUID, item *models.Item) (string, []interface{}, error)
	SelectItem(itemId, userId uuid.UUID) (string, []interface{}, error)
	UpdateItem(itemId, userId uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error)
	DeleteItem(itemId, userId uuid.UUID) (string, []interface{}, error)
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

	query, values, err := r.sqlBuilder.InsertItem(item.UserId, itemDb)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if err := r.db.Exec(ctx, query, values...); err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}

func (r *ItemRepository) GetItem(ctx context.Context, item *entities.Item) (*entities.Item, *errors.AppError) {
	itemDb := models.FromItemEntityToRepo(item)

	query, value, errBuilder := r.sqlBuilder.SelectItem(itemDb.Id, item.UserId)
	if errBuilder != nil {
		return nil, errors.NewInternalError(errBuilder)
	}

	if errQuery := r.db.QueryRow(ctx, itemDb, query, value...); errQuery != nil {
		if errQuery.Error() == "no rows in result set" {
			return nil, errors.ErrNotFound
		}
		return nil, errors.NewInternalError(errQuery)
	}

	return itemDb.ToEntity(), nil
}
func (r *ItemRepository) DeleteItem(ctx context.Context, item *entities.Item) *errors.AppError {
	itemDb := models.FromItemEntityToRepo(item)

	query, value, err := r.sqlBuilder.DeleteItem(itemDb.Id, item.UserId)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if err := r.db.Exec(ctx, query, value...); err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}

func (r *ItemRepository) UpdateItem(ctx context.Context, item *entities.UpdateItem, updateInfo map[string]interface{}) (*entities.Item, *errors.AppError) {
	itemDb := models.FromUpdateItemEntityToRepo(item)

	query, values, err := r.sqlBuilder.UpdateItem(itemDb.Id, item.UserId, updateInfo)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if errDbQuery := r.db.QueryRow(ctx, itemDb, query, values...); errDbQuery != nil {
		if errDbQuery.Error() == "no rows in result set" {
			return nil, errors.ErrNotFound
		}
		return nil, errors.NewInternalError(errDbQuery)
	}

	return itemDb.ToEntity(), nil
}
