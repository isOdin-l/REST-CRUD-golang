package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/configs"
	jwtToken "github.com/isOdin/RestApi/internal/middleware/dto"
	repoReqDTO "github.com/isOdin/RestApi/internal/repository/requestDTO"
	repoResDTO "github.com/isOdin/RestApi/internal/repository/responseDTO"
	"github.com/isOdin/RestApi/internal/service/requestDTO"
)

type AuthRepoInterface interface {
	CreateUser(ctx context.Context, user *repoReqDTO.CreateUser) (uuid.UUID, error)
	GetUser(ctx context.Context, user *repoReqDTO.GetUser) (*repoResDTO.GetedUser, error)
}

type AuthService struct {
	repo AuthRepoInterface
	cfg  *configs.InternalConfig
}

func NewAuthService(cfg *configs.InternalConfig, repo AuthRepoInterface) *AuthService {
	return &AuthService{cfg: cfg, repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, user *requestDTO.CreateUser) (uuid.UUID, error) {
	return s.repo.CreateUser(ctx, user.ConvertToRepoModel(s.generatePasswordHash(user.Password)))
}

func (s *AuthService) GenerateToken(ctx context.Context, user *requestDTO.GenerateToken) (string, error) {
	userFromDB, err := s.repo.GetUser(ctx, user.ConvertToRepoModel(s.generatePasswordHash(user.Password)))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtToken.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.TOKEN_TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: userFromDB.Id,
	})

	return token.SignedString([]byte(s.cfg.JWT_SIGNING_KEY))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.SALT)))
}
