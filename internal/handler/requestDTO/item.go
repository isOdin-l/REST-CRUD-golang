package requestDTO

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/service/requestDTO"
)

type CreateItem struct {
	UserId      uuid.UUID `form:"userId" validate:"required"`
	ListId      uuid.UUID `form:"list_id" validate:"required"`
	Title       string    `json:"title" form:"title" validate:"required"`
	Description string    `json:"description" form:"description"`
}

type GetItemById struct {
	UserId uuid.UUID `json:"user_id" form:"userId" validate:"required"`
	ItemId uuid.UUID `json:"item_id" form:"item_id"`
}

type DeleteItem struct {
	UserId uuid.UUID `json:"user_id" form:"userId" validate:"required"`
	ItemId uuid.UUID `json:"item_id" form:"item_id"`
}

type UpdateItem struct {
	UserId      uuid.UUID `json:"user_id" form:"userId" validate:"required"`
	ItemId      uuid.UUID `json:"item_id" form:"item_id"`
	Title       string    `json:"title" form:"titel"`
	Description string    `json:"description" form:"description"`
	Done        *bool     `json:"done" form:"done"`
}

func (m *CreateItem) ToServiceModel() *requestDTO.CreateItem {
	return &requestDTO.CreateItem{
		UserId:      m.UserId,
		ListId:      m.ListId,
		Title:       m.Title,
		Description: m.Description,
	}
}

func (m *GetItemById) ToServiceModel() *requestDTO.GetItemById {
	return &requestDTO.GetItemById{
		UserId: m.UserId,
		ItemId: m.ItemId,
	}
}

func (m *DeleteItem) ToServiceModel() *requestDTO.DeleteItem {
	return &requestDTO.DeleteItem{
		UserId: m.UserId,
		ItemId: m.ItemId,
	}
}

func (m *UpdateItem) ToServiceModel() *requestDTO.UpdateItem {
	return &requestDTO.UpdateItem{
		UserId:      m.UserId,
		ItemId:      m.ItemId,
		Title:       m.Title,
		Description: m.Description,
		Done:        m.Done,
	}
}
