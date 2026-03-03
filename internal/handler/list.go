package handler

import (
	"context"
	"net/http"

	mapper "isOdin/RestApi/internal/api"
	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/pkg/api"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type ListServiceInterface interface {
	CreateList(ctx context.Context, list *entities.List) (uuid.UUID, error)
	GetListById(ctx context.Context, list *entities.List) (*entities.List, error)
	DeleteList(ctx context.Context, list *entities.List) error
	UpdateList(ctx context.Context, list *entities.List) (*entities.List, error)
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
	if err := c.Bind(listApi); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	listId, errService := h.service.CreateList(c.Request().Context(), mapper.FromCreateListToEntity(&listApi))
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, errService.Error)
	}

	return c.JSON(http.StatusOK, listId)
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
	if err := c.Bind(listApi); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	list, errService := h.service.GetListById(c.Request().Context(), mapper.FromGetListToEntity(&listApi))
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, errService.Error)
	}

	return c.JSON(http.StatusOK, list)
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
	if err := c.Bind(listApi); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	list, errService := h.service.UpdateList(c.Request().Context(), mapper.FromUpdateListToEntity(&listApi))
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, errService.Error)
	}

	return c.JSON(http.StatusOK, list)
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
	if err := c.Bind(listApi); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	errService := h.service.DeleteList(c.Request().Context(), mapper.FromDeleteListToEntity(&listApi))
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, errService.Error)
	}

	return c.JSON(http.StatusOK, "List deleted")
}
