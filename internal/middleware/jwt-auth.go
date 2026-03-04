package middleware

import (
	"context"
	"strings"

	"isOdin/RestApi/configs"
	"isOdin/RestApi/internal/errors"
	"isOdin/RestApi/internal/middleware/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type AuthMiddleware struct {
	cfg *configs.InternalConfig
}

func NewAuthMiddleware(cfg *configs.InternalConfig) *AuthMiddleware {
	return &AuthMiddleware{cfg: cfg}
}

func (md *AuthMiddleware) JWTAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return errors.ResponseError(c, errors.ErrUnauthorized)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return errors.ResponseError(c, errors.ErrUnauthorized)
			}

			userId, err := md.parseJWTtoken(tokenString)
			if err != nil {
				return errors.ResponseError(c, errors.ErrUnauthorized)
			}

			// Созхранение userId в контекст
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "userId", userId)))

			return next(c)
		}
	}
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
