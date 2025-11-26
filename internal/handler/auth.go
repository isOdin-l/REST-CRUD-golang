package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	_ "github.com/isOdin/RestApi/api/apidto"
	"github.com/isOdin/RestApi/internal/handler/requestDTO"
	servReqDTO "github.com/isOdin/RestApi/internal/service/requestDTO"
	"github.com/isOdin/RestApi/tools/bindchi"
	"github.com/sirupsen/logrus"
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
func (h *Auth) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var reqUser requestDTO.SignUpUser
	if err := bindchi.BindValidate(r, &reqUser, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := h.service.CreateUser(r.Context(), reqUser.ConvertToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"id": userId,
	})
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
func (h *Auth) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var reqUser requestDTO.SignInUser
	if err := bindchi.BindValidate(r, &reqUser, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	generatedToken, err := h.service.GenerateToken(r.Context(), reqUser.ConvertToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"token": generatedToken,
	})
}
