package repository

import (
	"context"

	"isOdin/RestApi/internal/database"
	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/repository/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type AuthRepository struct {
	db   IDatabase
	psql sq.StatementBuilderType
}

func NewAuthRepository(db IDatabase) *AuthRepository {
	return &AuthRepository{db: db, psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *entities.User) error {
	userDb := models.FromUserEntityToRepo(user)
	query, value, err := database.InsertUser(&r.psql, userDb.Id, userDb.Name, userDb.Username, userDb.Password_hash)
	if err != nil {
		return err
	}

	return r.db.Exec(ctx, query, value...)
}

func (r *AuthRepository) GetUser(ctx context.Context, userId uuid.UUID) (*entities.User, error) {
	query, value, err := database.SelectUser(&r.psql, userId)
	if err != nil {
		return nil, err
	}

	userDb := &models.User{}
	err = r.db.QueryRow(ctx, userDb, query, value...)

	return userDb.ToEntity(), err
}
