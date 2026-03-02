package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"isOdin/RestApi/configs"
	"isOdin/RestApi/internal/entities"
	jwtToken "isOdin/RestApi/internal/middleware/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthRepoInterface interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, userId uuid.UUID) (*entities.User, error)
}

type AuthService struct {
	repo AuthRepoInterface
	cfg  *configs.InternalConfig
}

func NewAuthService(cfg *configs.InternalConfig, repo AuthRepoInterface) *AuthService {
	return &AuthService{cfg: cfg, repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, user *entities.User) (uuid.UUID, error) {
	var err error
	user.UserId, err = uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}
	user.Password = s.generatePasswordHash(user.Password)

	return user.UserId, s.repo.CreateUser(ctx, user)
}

func (s *AuthService) GenerateToken(ctx context.Context, user *entities.User) (string, error) {
	user.Password = s.generatePasswordHash(user.Password)

	userFromDB, err := s.repo.GetUser(ctx, user.UserId)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtToken.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.TOKEN_TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: userFromDB.UserId,
	})

	return token.SignedString([]byte(s.cfg.JWT_SIGNING_KEY))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.SALT)))
}
