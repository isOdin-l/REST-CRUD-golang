package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"isOdin/RestApi/configs"
	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/errors"
	jwtToken "isOdin/RestApi/internal/middleware/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthRepoInterface interface {
	CreateUser(ctx context.Context, user *entities.User) *errors.AppError
	GetUser(ctx context.Context, user *entities.User) (*entities.User, *errors.AppError)
}

type AuthService struct {
	repo AuthRepoInterface
	cfg  *configs.InternalConfig
	txMn ITransactionManager
}

func NewAuthService(cfg *configs.InternalConfig, repo AuthRepoInterface, txMn ITransactionManager) *AuthService {
	return &AuthService{cfg: cfg, repo: repo, txMn: txMn}
}

func (s *AuthService) CreateUser(ctx context.Context, user *entities.User) (string, *errors.AppError) {
	var err error
	user.UserId, err = uuid.NewV7()
	if err != nil {
		return "", errors.NewInternalError(err)
	}
	user.Password = s.generatePasswordHash(user.Password)

	errCreate := s.repo.CreateUser(ctx, user)
	if errCreate != nil{
		return "", errCreate
	}
	
	return s.signJwtToken(user.UserId)
}


func (s *AuthService) LogInUser(ctx context.Context, user *entities.User) (string, *errors.AppError) {
	user.Password = s.generatePasswordHash(user.Password)

	userFromDB, errRepo := s.repo.GetUser(ctx, user)
	if errRepo != nil {
		return "", errRepo
	}

	return s.signJwtToken(userFromDB.UserId)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.SALT)))
}

func (s *AuthService) signJwtToken(userId uuid.UUID) (string, *errors.AppError){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtToken.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.TOKEN_TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: userId,
	})

	tokenString, errJwt := token.SignedString([]byte(s.cfg.JWT_SIGNING_KEY))
	if errJwt != nil {
		return "", errors.NewInternalError(errJwt)
	}

	return tokenString, nil
}