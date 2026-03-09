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
	GetUserByUsernameAndPassword(user *models.User) (string, []interface{}, error)
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
	query, value, err := r.sqlBuilder.InsertUser(userDb.Id.String(), userDb.Name, userDb.Username, userDb.Password_hash)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if err := r.db.Exec(ctx, query, value...); err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}

func (r *AuthRepository) GetUser(ctx context.Context, user *entities.User) (*entities.User, *errors.AppError) {
	userDb := models.FromUserEntityToRepo(user)

	query, value, err := r.sqlBuilder.GetUserByUsernameAndPassword(userDb)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	var tmpUserId string
	if err := r.db.QueryRow(ctx, query, value...).Scan(&tmpUserId, &userDb.Name, &userDb.Username, &userDb.Password_hash); err != nil {
		if err.Error() == "no rows in result set" {
			return nil, errors.ErrNotFound
		}
		return nil, errors.NewInternalError(err)
	}
	userDb.Id, _ = uuid.Parse(tmpUserId)

	return userDb.ToEntity(), nil
}
