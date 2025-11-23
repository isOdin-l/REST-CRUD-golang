package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/handler/requestDTO"
	"github.com/isOdin/RestApi/internal/models"
	serReqDTO "github.com/isOdin/RestApi/internal/service/requestDTO"
	serResDTO "github.com/isOdin/RestApi/internal/service/responseDTO"
	"github.com/isOdin/RestApi/tools/bindchi"
	"github.com/sirupsen/logrus"
)

type ItemServiceInterface interface {
	CreateItem(ctx context.Context, itemInfo models.CreateItemParams) (uuid.UUID, error)
	GetAllItems(userId uuid.UUID) (*[]serResDTO.GetItem, error)
	GetItemById(itemInfo *serReqDTO.GetItemById) (*serResDTO.GetItemById, error)
	DeleteItem(itemInfo *serReqDTO.DeleteItem) error
	UpdateItem(itemInfo *serReqDTO.UpdateItem) error
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
func (h *Item) CreateItem(w http.ResponseWriter, r *http.Request) {
	var reqItem requestDTO.CreateItem
	if err := bindchi.BindValidate(r, &reqItem, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	itemParams := models.CreateItemParams{
		UserId:      reqItem.UserId,
		ListId:      reqItem.ListId,
		Title:       reqItem.Title,
		Description: reqItem.Description,
	}
	itemId, err := h.service.CreateItem(r.Context(), itemParams) // TODO: to private, контекст и т.д
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Item -

	render.JSON(w, r, map[string]interface{}{
		"itemId": itemId,
	})
}

// @Summary Get all todo-items
// @Security ApiKeyAuth
// @Tags items
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists/items [get]
func (h *Item) GetAllItems(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	items, err := h.service.GetAllItems(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"items": *items,
	})
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
func (h *Item) GetItemById(w http.ResponseWriter, r *http.Request) {
	var itemInfo requestDTO.GetItemById
	if err := bindchi.BindValidate(r, &itemInfo, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item, err := h.service.GetItemById(itemInfo.ToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, *item)
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
func (h *Item) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updItem requestDTO.UpdateItem
	if err := bindchi.BindValidate(r, &updItem, h.validate); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.UpdateItem(updItem.ToServiceModel()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{})
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
func (h *Item) DeleteItem(w http.ResponseWriter, r *http.Request) {
	var itemInfo requestDTO.DeleteItem
	if err := bindchi.BindValidate(r, &itemInfo, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.DeleteItem(itemInfo.ToServiceModel()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{})
}
