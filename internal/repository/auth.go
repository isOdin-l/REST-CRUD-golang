package repository

import (
	"context"

	"isOdin/RestApi/internal/entities"
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

func (r *AuthRepository) CreateUser(ctx context.Context, user *entities.User) error {
	userDb := models.FromUserEntityToRepo(user)
	query, value, err := r.sqlBuilder.InsertUser(userDb.Id, userDb.Name, userDb.Username, userDb.Password_hash)
	if err != nil {
		return err
	}

	return r.db.Exec(ctx, query, value...)
}

func (r *AuthRepository) GetUser(ctx context.Context, userId uuid.UUID) (*entities.User, error) {
	query, value, err := r.sqlBuilder.SelectUser(userId)
	if err != nil {
		return nil, err
	}

	userDb := &models.User{}
	err = r.db.QueryRow(ctx, userDb, query, value...)

	return userDb.ToEntity(), err
}
