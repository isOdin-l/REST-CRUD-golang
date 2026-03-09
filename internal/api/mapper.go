package api

import (
	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/pkg/api"

	"github.com/google/uuid"
)

// ========================
// ==== Request models ====
// ========================

// --------- USER ---------
func FromSignUpApiToEntity(req *api.SignUp) *entities.User {
	return &entities.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
}

func FromSignInApiToEntity(req *api.SignIn) *entities.User {
	return &entities.User{
		Username: req.Username,
		Password: req.Password,
	}
}

// --------- ITEM ---------
func FromCreateItemToEntity(req *api.CreateItem) *entities.Item {
	return &entities.Item{
		ListId:      req.ListId,
		Title:       req.Title,
		Description: req.Description,
		Done:        false,
	}
}

func FromGetItemToEntity(req *api.GetItem) *entities.Item {
	return &entities.Item{
		ItemId: req.ItemId,
	}
}

func FromUpdateItemToEntity(req *api.UpdateItem) *entities.UpdateItem {
	return &entities.UpdateItem{
		ItemId: req.ItemId,
		OptValues: struct {
			Title       *string
			Description *string
			Done        *bool
		}{
			Title:       req.Title,
			Description: req.Description,
			Done:        req.Done,
		},
	}
}

func FromDeleteItemToEntity(req *api.DeleteItem) *entities.Item {
	return &entities.Item{
		ItemId: req.ItemId,
	}
}

// --------- LIST ---------
func FromCreateListToEntity(req *api.CreateList) *entities.List {
	return &entities.List{
		Title:       req.Title,
		Description: req.Description,
	}
}

func FromGetListToEntity(req *api.GetList) *entities.List {
	return &entities.List{
		ListId: req.ListId,
	}
}

func FromUpdateListToEntity(req *api.UpdateList) *entities.UpdateList {
	return &entities.UpdateList{
		ListId: req.ListId,
		OptValues: struct {
			Title       *string
			Description *string
		}{
			Title:       req.Title,
			Description: req.Description,
		},
	}
}

func FromDeleteListToEntity(req *api.DeleteList) *entities.List {
	return &entities.List{
		ListId: req.ListId,
	}
}

// =======================
// === Response Models ===
// =======================

// -------- USER ---------

// -------- LIST ---------
func FromEntityToListApi(req *entities.List) *api.ResponseList {
	return &api.ResponseList{
		List: struct {
			ListId      uuid.UUID "json:\"id\""
			Title       string    "json:\"title\""
			Description string    "json:\"description\""
		}{
			ListId:      req.ListId,
			Title:       req.Title,
			Description: req.Description,
		},
	}
}

// -------- ITEM ---------
func FromEntityToItemApi(req *entities.Item) *api.ResponseItem {
	return &api.ResponseItem{
		Item: struct {
			ItemId      uuid.UUID "json:\"item_id\""
			Title       string    "json:\"title\""
			Description string    "json:\"description\""
			Done        bool      "json:\"done\""
		}{
			ItemId:      req.ItemId,
			Title:       req.Title,
			Description: req.Description,
			Done:        req.Done,
		},
	}
}
