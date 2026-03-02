package handler

import (
	"context"
	"net/http"

	_ "isOdin/RestApi/api/apidto"
	"isOdin/RestApi/internal/handler/requestDTO"
	servReqDTO "isOdin/RestApi/internal/service/requestDTO"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type AuthServiceInterface interface {
	CreateUser(ctx context.Context, user *servReqDTO.CreateUser) (uuid.UUID, error)
	GenerateToken(ctx context.Context, user *servReqDTO.GenerateToken) (string, error)
}

type Auth struct {
	validate *validator.Validate
	service  AuthServiceInterface
}

func NewAuthHandler(validate *validator.Validate, service AuthServiceInterface) *Auth {
	return &Auth{validate: validate, service: service}
}

// @Summary SignUp
// @Tags auth
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body apidto.SignUpAPI true "account info"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /auth/sign-up [post]
func (h *Auth) SignUpHandler(c *echo.Context) error {
	var reqUser requestDTO.SignUpUser

	userId, err := h.service.CreateUser(c.Request().Context(), reqUser.ConvertToServiceModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userId)
}

// @Summary SignIn
// @Tags auth
// @ID log-into-account
// @Accept  json
// @Produce  json
// @Param input body apidto.SignInAPI true "account info"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /auth/sign-in [post]
func (h *Auth) SignInHandler(c *echo.Context) error {
	var reqUser requestDTO.SignInUser

	generatedToken, err := h.service.GenerateToken(c.Request().Context(), reqUser.ConvertToServiceModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, generatedToken)
}
