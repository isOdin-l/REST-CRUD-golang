package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/database"
	"github.com/isOdin/RestApi/internal/repository/requestDTO"
	"github.com/isOdin/RestApi/internal/repository/responseDTO"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *requestDTO.CreateUser) (uuid.UUID, error) {
	queryString := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", database.TableUsers)
	row := r.db.QueryRow(ctx, queryString, user.Name, user.Username, user.PasswordHash)

	var userId uuid.UUID
	err := row.Scan(&userId)

	return userId, err
}

func (r *AuthRepository) GetUser(ctx context.Context, user *requestDTO.GetUser) (*responseDTO.GetedUser, error) {
	var userResp responseDTO.GetedUser

	queryString := fmt.Sprintf("SELECT id, name, username, password_hash FROM %s WHERE username = $1 AND password_hash = $2 LIMIT 1", database.TableUsers)
	err := r.db.QueryRow(ctx, queryString, user.Username, user.PasswordHash).Scan(&userResp.Id, &userResp.Name, &userResp.Username, &userResp.PasswordHash)

	return &userResp, err
}
