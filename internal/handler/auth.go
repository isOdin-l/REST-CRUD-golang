package handler

import (
	"context"
	mapper "isOdin/RestApi/internal/api"
	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/pkg/api"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type AuthServiceInterface interface {
	CreateUser(ctx context.Context, user *entities.User) (uuid.UUID, error)
	GenerateToken(ctx context.Context, user *entities.User) (string, error)
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
	userFromApi := new(api.SignUp)
	if err := c.Bind(userFromApi); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userId, err := h.service.CreateUser(c.Request().Context(), mapper.FromSignUpApiToEntity(userFromApi))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userResp := &entities.User{
		UserId: userId,
	}

	return c.JSON(http.StatusOK, mapper.FromEntityToSignUpApi(userResp))
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
	var userApi api.SignIn
	if err := c.Bind(userApi); err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	generatedToken, err := h.service.GenerateToken(c.Request().Context(), mapper.FromSignInApiToEntity(&userApi))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &api.ResponseSignIn{Token: generatedToken})
}
