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

type ListRepository struct {
	db   *pgxpool.Pool
	psql sq.StatementBuilderType
}

func NewListRepository(db *pgxpool.Pool) *ListRepository {
	return &ListRepository{db: db, psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (r *ListRepository) CreateList(ctx context.Context, list *entities.List) (uuid.UUID, error) {
	query, value, err := database.InsertList(&r.psql, list.Id, list.Author_id, list.Title, list.Description)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = r.db.Exec(ctx, query, value...)
	return list.Id, err
}

func (r *ListRepository) GetList(ctx context.Context, list *models.List) (*models.List, error) {
	query, value, err := database.SelectList(&r.psql, list.Id)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(ctx, query, value...).Scan(list)

	return list, err
}

func (r *ListRepository) DeleteList(ctx context.Context, list *models.List) error {
	query, value, err := database.DeleteList(&r.psql, list.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, query, value...)

	return err
}

func (r *ListRepository) UpdateList(ctx context.Context, list *models.List, updateInfo map[string]interface{}) (*models.List, error) {
	query, value, err := database.UpdateList(&r.psql, list.Id, updateInfo)
	if err != nil {
		return nil, err
	}

	errDbQuery := r.db.QueryRow(ctx, query, value...).Scan(list)

	return list, errDbQuery
}
