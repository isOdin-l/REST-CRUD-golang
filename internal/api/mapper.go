package api

import (
	"isOdin/RestApi/internal/entities"
	"isOdin/RestApi/pkg/api"

	"github.com/google/uuid"
)

// === USER ===
func FromSignUpApiToEntity(req *api.SignUp) *entities.User {
	return &entities.User{
		UserId:   uuid.Nil,
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

func FromEntityToSignUpApi(req *entities.User) *api.ResponseSignUp {
	return &api.ResponseSignUp{
		User: struct {
			UserId uuid.UUID "json:\"id\""
		}{
			UserId: req.UserId,
		},
	}
}

// === ITEM ===
func FromCreateItemToEntity(req *api.CreateItem) *entities.Item {
	return &entities.Item{
		ListId:      req.ListId,
		ItemId:      uuid.Nil,
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

func FromUpdateItemToEntity(req *api.UpdateItem) *entities.Item {
	return &entities.Item{
		ItemId:      req.ItemId,
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
	}
}

func FromDeleteItemToEntity(req *api.DeleteItem) *entities.Item {
	return &entities.Item{
		ItemId: req.ItemId,
	}
}

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

// === LIST ===
func FromCreateListToEntity(req *api.CreateList) *entities.List {
	return &entities.List{
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
	}
}

func FromGetListToEntity(req *api.GetList) *entities.List {
	return &entities.List{
		ListId: req.ListId,
	}
}

func FromUpdateListToEntity(req *api.UpdateList) *entities.List {
	return &entities.List{
		UserId:      req.UserId,
		ListId:      req.ListId,
		Title:       req.Title,
		Description: req.Description,
	}
}

func FromDeleteListToEntity(req *api.DeleteList) *entities.List {
	return &entities.List{
		ListId: req.ListId,
	}
}

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
