package handler

import (
	"context"
	"fmt"
	"net/http"

	mapper "isOdin/RestApi/internal/api"
	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/internal/errors"
	"isOdin/RestApi/pkg/api"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type ListServiceInterface interface {
	CreateList(ctx context.Context, list *entities.List) (*entities.List, *errors.AppError)
	GetListById(ctx context.Context, list *entities.List) (*entities.List, *errors.AppError)
	DeleteList(ctx context.Context, list *entities.List) *errors.AppError
	UpdateList(ctx context.Context, list *entities.UpdateList) (*entities.List, *errors.AppError)
}

type List struct {
	validate *validator.Validate
	service  ListServiceInterface
}

func NewListHandler(validate *validator.Validate, service ListServiceInterface) *List {
	return &List{validate: validate, service: service}
}

// @Summary Create todo-list
// @Security ApiKeyAuth
// @Tags lists
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body apidto.CreateList true "list info"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists [post]
func (h *List) CreateList(c *echo.Context) error {
	var listApi api.CreateList
	if err := c.Bind(&listApi); err != nil {
		return errors.ResponseError(c, errors.ErrBadRequest)
	}

	if err := h.validate.Struct(listApi); err != nil {
		return errors.ResponseError(c, errors.ErrBadRequest)
	}

	listEntity := mapper.FromCreateListToEntity(&listApi)
	listEntity.UserId = c.Get("userId").(uuid.UUID)

	list, errService := h.service.CreateList(c.Request().Context(), listEntity)
	if errService != nil {
		return errors.ResponseError(c, errService)
	}

	return c.JSON(http.StatusOK, mapper.FromEntityToListApi(list))
}

// @Summary Get todo-lists by Id
// @Security ApiKeyAuth
// @Tags lists
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Param list_id path string true "List Id"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists/{list_id} [get]
func (h *List) GetList(c *echo.Context) error {
	var listApi api.GetList
	if err := c.Bind(&listApi); err != nil {
		return errors.ResponseError(c, errors.ErrBadRequest)
	}

	list, errService := h.service.GetListById(c.Request().Context(), mapper.FromGetListToEntity(&listApi))
	if errService != nil {
		return errors.ResponseError(c, errService)
	}

	return c.JSON(http.StatusOK, mapper.FromEntityToListApi(list))
}

// @Summary Update todo-list
// @Security ApiKeyAuth
// @Tags lists
// @ID update-list
// @Accept  json
// @Produce  json
// @Param list_id path string true "List Id"
// @Param input body apidto.UpdateList true "list info"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists/{list_id} [put]
func (h *List) UpdateList(c *echo.Context) error {
	var listApi api.UpdateList
	if err := c.Bind(&listApi); err != nil {
		return errors.ResponseError(c, errors.ErrBadRequest)
	}

	if err := h.validate.Struct(listApi); err != nil {
		return errors.ResponseError(c, errors.ErrValidation)
	}

	listEntity := mapper.FromUpdateListToEntity(&listApi)
	listEntity.UserId = c.Get("userId").(uuid.UUID)

	list, errService := h.service.UpdateList(c.Request().Context(), listEntity)
	if errService != nil {
		return errors.ResponseError(c, errService)
	}

	return c.JSON(http.StatusOK, mapper.FromEntityToListApi(list))
}

// @Summary Delete todo-list
// @Security ApiKeyAuth
// @Tags lists
// @ID delete-list
// @Accept  json
// @Produce  json
// @Param list_id path string true "List Id"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists/{list_id} [delete]
func (h *List) DeleteList(c *echo.Context) error {
	var listApi api.DeleteList
	if err := c.Bind(&listApi); err != nil {
		return errors.ResponseError(c, errors.ErrBadRequest)
	}

	errService := h.service.DeleteList(c.Request().Context(), mapper.FromDeleteListToEntity(&listApi))
	if errService != nil {
		return errors.ResponseError(c, errService)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("List %s deleted", listApi.ListId),
	})
}
