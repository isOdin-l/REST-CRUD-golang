package middleware

import (
	"isOdin/RestApi/configs"
)

type Middleware struct {
	*AuthMiddleware
}

func NewMiddleware(cfg *configs.InternalConfig) *Middleware {
	return &Middleware{
		AuthMiddleware: NewAuthMiddleware(cfg),
	}
}
