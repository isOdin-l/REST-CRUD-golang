package handler

import (
	"context"
	"net/http"

	_ "isOdin/RestApi/api/apidto"
	"isOdin/RestApi/internal/handler/requestDTO"
	"isOdin/RestApi/internal/handler/responseDTO"
	reqSerDTO "isOdin/RestApi/internal/service/requestDTO"
	resSerDTO "isOdin/RestApi/internal/service/responseDTO"
	"isOdin/RestApi/tools/bindchi"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ListServiceInterface interface {
	CreateList(ctx context.Context, listInfo *reqSerDTO.CreateList) (uuid.UUID, error)
	GetAllLists(ctx context.Context, userId uuid.UUID) (*[]resSerDTO.GetList, error)
	GetListById(ctx context.Context, listInfo *reqSerDTO.GetListById) (*resSerDTO.GetListById, error)
	DeleteList(ctx context.Context, listInfo *reqSerDTO.DeleteList) error
	UpdateList(ctx context.Context, listInfo *reqSerDTO.UpdateList) error
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
func (h *List) CreateList(w http.ResponseWriter, r *http.Request) {
	var reqList requestDTO.CreateList
	if err := bindchi.BindValidate(r, &reqList, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listId, err := h.service.CreateList(r.Context(), reqList.ToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"listId": listId,
	})
}

// @Summary Get all todo-lists
// @Security ApiKeyAuth
// @Tags lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure default {string} string
// @Router /api/lists [get]
func (h *List) GetAllLists(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	listsResponsed, err := h.service.GetAllLists(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lists := make([]responseDTO.GetList, len(*listsResponsed))
	for i := range len(*listsResponsed) {
		// ------- Указатель на массив -> массив -> элемент массива -> перевод элемента в указатель на другой тип -> элемент другого типа -------
		lists[i] = *((*listsResponsed)[i].ToHandlerModel())
	}

	render.JSON(w, r, map[string]interface{}{
		"lists": lists,
	})
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
func (h *List) GetListById(w http.ResponseWriter, r *http.Request) {
	var listInfo requestDTO.GetListById
	if err := bindchi.BindValidate(r, &listInfo, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list, err := h.service.GetListById(r.Context(), listInfo.ToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"list": list,
	})
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
func (h *List) UpdateList(w http.ResponseWriter, r *http.Request) {
	var reqUpdList requestDTO.UpdateList
	if err := bindchi.BindValidate(r, &reqUpdList, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateList(r.Context(), reqUpdList.ToServiceModel()); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{})
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
func (h *List) DeleteList(w http.ResponseWriter, r *http.Request) {
	var listInfo requestDTO.DeleteList
	if err := bindchi.BindValidate(r, &listInfo, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.DeleteList(r.Context(), listInfo.ToServiceModel()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{})
}
