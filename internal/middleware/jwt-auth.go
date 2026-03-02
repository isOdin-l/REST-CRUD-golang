package middleware

import (
	"context"
	"net/http"
	"strings"

	"isOdin/RestApi/configs"
	"isOdin/RestApi/internal/middleware/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthMiddleware struct {
	cfg *configs.InternalConfig
}

func NewAuthMiddleware(cfg *configs.InternalConfig) *AuthMiddleware {
	return &AuthMiddleware{cfg: cfg}
}

func (md *AuthMiddleware) JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Получаем токен из заголовка Authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			// Проверяем формат: "Bearer <token>"
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			// Парсим токен
			userId, err := md.parseJWTtoken(tokenString)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			// Записываем в контекст, чтобы другие хэндлеры/мидлвейры могли работать с данными
			r = r.WithContext(context.WithValue(r.Context(), "userId", userId))
			next.ServeHTTP(w, r)
		},
	)
}

func (md *AuthMiddleware) parseJWTtoken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &dto.TokenClaims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrInvalidKeyType
		}

		return []byte(md.cfg.JWT_SIGNING_KEY), nil
	})

	if err != nil {
		return uuid.Nil, err
	}
	claims, ok := token.Claims.(*dto.TokenClaims)
	if !ok {
		return uuid.Nil, jwt.ErrTokenInvalidClaims
	}

	return claims.UserId, nil
}
