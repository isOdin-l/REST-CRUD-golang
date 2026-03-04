package repository

import (
	"context"

	"isOdin/RestApi/internal/entities"
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

func (r *ItemRepository) CreateItem(ctx context.Context, item *entities.Item) error {
	itemDb := models.FromItemEntityToRepo(item)
	query, values, err := r.sqlBuilder.InsertItem(itemDb.Id, itemDb.List_id, itemDb.Title, itemDb.Description)
	if err != nil {
		return err
	}

	return r.db.Exec(ctx, query, values...)
}

func (r *ItemRepository) GetItem(ctx context.Context, itemId uuid.UUID) (*entities.Item, error) {
	query, value, errBuilder := r.sqlBuilder.SelectItem(itemId)
	if errBuilder != nil {
		return nil, errBuilder
	}

	item := &models.Item{}
	if errQuery := r.db.QueryRow(ctx, item, query, value...); errQuery != nil {
		return nil, errQuery
	}

	return item.ToEntity(), nil
}
func (r *ItemRepository) DeleteItem(ctx context.Context, itemId uuid.UUID) error {
	query, value, err := r.sqlBuilder.DeleteItem(itemId)
	if err != nil {
		return err
	}

	return r.db.Exec(ctx, query, value...)
}

func (r *ItemRepository) UpdateItem(ctx context.Context, itemId uuid.UUID, updateInfo map[string]interface{}) (*entities.Item, error) {
	query, values, err := r.sqlBuilder.UpdateItem(itemId, updateInfo)
	if err != nil {
		return nil, err
	}
	item := &models.Item{}
	if errDbQuery := r.db.QueryRow(ctx, item, query, values...); errDbQuery != nil {
		return nil, errDbQuery
	}

	return item.ToEntity(), nil
}
