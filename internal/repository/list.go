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

func (r *ListRepository) CreateList(ctx context.Context, list *entities.List) error {
	listDb := models.FromListEntityToRepo(list)
	query, value, err := database.InsertList(&r.psql, listDb.Id, listDb.Author_id, listDb.Title, listDb.Description)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, value...)
	return err
}

func (r *ListRepository) GetList(ctx context.Context, listId uuid.UUID) (*entities.List, error) {
	query, value, err := database.SelectList(&r.psql, listId)
	if err != nil {
		return nil, err
	}

	list := &models.List{}
	err = r.db.QueryRow(ctx, query, value...).Scan(list)

	return list.ToEntity(), err
}

func (r *ListRepository) DeleteList(ctx context.Context, listId uuid.UUID) error {
	query, value, err := database.DeleteList(&r.psql, listId)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, query, value...)

	return err
}

func (r *ListRepository) UpdateList(ctx context.Context, listId uuid.UUID, updateInfo map[string]interface{}) (*entities.List, error) {
	query, value, err := database.UpdateList(&r.psql, listId, updateInfo)
	if err != nil {
		return nil, err
	}

	list := &models.List{}
	errDbQuery := r.db.QueryRow(ctx, query, value...).Scan(list)

	return list.ToEntity(), errDbQuery
}
