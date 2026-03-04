package repository

import (
	"context"

	"isOdin/RestApi/internal/database"
	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/repository/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type ItemRepository struct {
	db   IDatabase
	psql sq.StatementBuilderType
}

func NewItemRepository(db IDatabase) *ItemRepository {
	return &ItemRepository{db: db, psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (r *ItemRepository) CreateItem(ctx context.Context, item *entities.Item) error {
	itemDb := models.FromItemEntityToRepo(item)
	query, values, err := database.InsertItem(&r.psql, itemDb.Id, itemDb.List_id, itemDb.Title, itemDb.Description)
	if err != nil {
		return err
	}

	return r.db.Exec(ctx, query, values...)
}

func (r *ItemRepository) GetItem(ctx context.Context, itemId uuid.UUID) (*entities.Item, error) {
	query, value, errBuilder := database.SelectItem(&r.psql, itemId)
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
	query, value, err := database.DeleteItem(&r.psql, itemId)
	if err != nil {
		return err
	}

	return r.db.Exec(ctx, query, value...)
}

func (r *ItemRepository) UpdateItem(ctx context.Context, itemId uuid.UUID, updateInfo map[string]interface{}) (*entities.Item, error) {
	query, values, err := database.UpdateItem(&r.psql, itemId, updateInfo)
	if err != nil {
		return nil, err
	}
	item := &models.Item{}
	if errDbQuery := r.db.QueryRow(ctx, item, query, values...); errDbQuery != nil {
		return nil, errDbQuery
	}

	return item.ToEntity(), nil
}
