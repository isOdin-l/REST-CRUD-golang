package repository

import (
	"context"

	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/errors"
	"isOdin/RestApi/internal/repository/models"

	"github.com/google/uuid"
)

type IAuthSql interface {
	InsertUser(values ...any) (string, []interface{}, error)
	SelectUser(userId uuid.UUID) (string, []interface{}, error)
}

type AuthRepository struct {
	db         IDatabase
	sqlBuilder IAuthSql
}

func NewAuthRepository(db IDatabase, sqlBuilder IAuthSql) *AuthRepository {
	return &AuthRepository{db: db, sqlBuilder: sqlBuilder}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *entities.User) *errors.AppError {
	userDb := models.FromUserEntityToRepo(user)
	query, value, err := r.sqlBuilder.InsertUser(userDb.Id, userDb.Name, userDb.Username, userDb.Password_hash)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if err := r.db.Exec(ctx, query, value...); err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}

func (r *AuthRepository) GetUser(ctx context.Context, userId uuid.UUID) (*entities.User, *errors.AppError) {
	query, value, err := r.sqlBuilder.SelectUser(userId)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	userDb := &models.User{}
	if err := r.db.QueryRow(ctx, userDb, query, value...); err != nil {
		if err.Error() == "no rows in result set" {
			return nil, errors.ErrNotFound
		}
		return nil, errors.NewInternalError(err)
	}

	return userDb.ToEntity(), nil
}
