package repository

import (
	"context"

	"isOdin/RestApi/internal/database"
	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/repository/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemRepository struct {
	db   *pgxpool.Pool
	psql sq.StatementBuilderType
}

func NewItemRepository(db *pgxpool.Pool) *ItemRepository {
	return &ItemRepository{db: db, psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (r *ItemRepository) CreateItem(ctx context.Context, item *models.Item) (uuid.UUID, error) {
	query, values, err := database.InsertItem(&r.psql, item.Id, item.List_id, item.Title, item.Description)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = r.db.Exec(ctx, query, values...)

	return item.Id, err
}

func (r *ItemRepository) GetItem(ctx context.Context, itemId uuid.UUID) (*entities.Item, error) {
	query, value, errBuilder := database.SelectItem(&r.psql, itemId)
	if errBuilder != nil {
		return nil, errBuilder
	}

	item := &entities.Item{}
	errQuery := r.db.QueryRow(ctx, query, value...).Scan(item)

	return item, errQuery
}
func (r *ItemRepository) DeleteItem(ctx context.Context, itemId uuid.UUID) error {
	query, value, err := database.DeleteItem(&r.psql, itemId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, value...)

	return err
}

func (r *ItemRepository) UpdateItem(ctx context.Context, itemId uuid.UUID, updateInfo map[string]interface{}) (*entities.Item, error) {
	query, values, err := database.UpdateItem(&r.psql, itemId, updateInfo)
	if err != nil {
		return nil, err
	}
	item := &entities.Item{}
	errDbQuery := r.db.QueryRow(ctx, query, values...).Scan(item)

	return item, errDbQuery
}
