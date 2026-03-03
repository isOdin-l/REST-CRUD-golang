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

type ItemServiceInterface interface {
	CreateItem(ctx context.Context, item *entities.Item) (uuid.UUID, error)
	GetItem(ctx context.Context, item *entities.Item) (*entities.Item, error)
	DeleteItem(ctx context.Context, item *entities.Item) error
	UpdateItem(ctx context.Context, item *entities.Item) (*entities.Item, error)
}

type Item struct {
	validate *validator.Validate
	service  ItemServiceInterface
}

func NewItemHandler(validate *validator.Validate, service ItemServiceInterface) *Item {
	return &Item{validate: validate, service: service}
}

// @Summary Create todo-item
// @Security ApiKeyAuth
// @Tags items
// @ID create-item
// @Accept  json
// @Produce  json
// @Param list_id path string true "List Id"
// @Param input body apidto.CreateItem true "Item info"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists/{list_id}/items [post]
func (h *Item) CreateItem(c *echo.Context) error {
	var itemApi api.CreateItem
	if err := c.Bind(itemApi); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	itemId, errService := h.service.CreateItem(c.Request().Context(), mapper.FromCreateItemToEntity(&itemApi))
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, errService.Error())
	}

	return c.JSON(http.StatusOK, itemId)
}

// @Summary Get todo-item by Id
// @Security ApiKeyAuth
// @Tags items
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param item_id path string true "Item Id"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists/items/{item_id} [get]
func (h *Item) GetItem(c *echo.Context) error {
	var itemApi api.GetItem
	if err := c.Bind(itemApi); err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	item, errService := h.service.GetItem(c.Request().Context(), mapper.FromGetItemToEntity(&itemApi))
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, errService.Error())
	}

	return c.JSON(http.StatusOK, item)
}

// @Summary Update todo-item
// @Security ApiKeyAuth
// @Tags items
// @ID update-item
// @Accept  json
// @Produce  json
// @Param item_id path string true "Item Id"
// @Param input body apidto.UpdateItem true "Item info"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists/items/{item_id} [put]
func (h *Item) UpdateItem(c *echo.Context) error {
	var itemApi api.UpdateItem
	if err := c.Bind(itemApi); err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	item, errService := h.service.UpdateItem(c.Request().Context(), mapper.FromUpdateItemToEntity(&itemApi))
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, errService.Error)
	}

	return c.JSON(http.StatusOK, item)
}

// @Summary Delete todo-item
// @Security ApiKeyAuth
// @Tags items
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param item_id path string true "Item Id"
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists/items/{item_id} [delete]
func (h *Item) DeleteItem(c *echo.Context) error {
	var itemApi api.DeleteItem
	if err := c.Bind(itemApi); err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	errService := h.service.DeleteItem(c.Request().Context(), mapper.FromDeleteItemToEntity(&itemApi))
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, errService.Error)
	}

	return c.JSON(http.StatusOK, "Item deleted")
}
