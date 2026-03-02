package models

import (
	"isOdin/RestApi/internal/entities"
)

func (req *User) ToEntity() *entities.User {
	return &entities.User{
		UserId:   req.Id,
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password_hash,
	}
}

func (req *Item) ToEntity() *entities.Item {
	return &entities.Item{
		ItemId:      req.Id,
		ListId:      req.List_id,
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
	}
}

func (req *List) ToEntity() *entities.List {
	return &entities.List{
		ListId:      req.Id,
		UserId:      req.Author_id,
		Title:       req.Title,
		Description: req.Description,
	}
}

func FromUserEntityToRepo(req *entities.User) *User {
	return &User{
		Id:            req.UserId,
		Name:          req.Name,
		Username:      req.Username,
		Password_hash: req.Password,
	}
}

func FromItemEntityToRepo(req *entities.Item) *Item {
	return &Item{
		Id:          req.ItemId,
		List_id:     req.ListId,
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
	}
}

func FromListEntityToRepo(req *entities.List) *List {
	return &List{
		Id:          req.ListId,
		Author_id:   req.UserId,
		Title:       req.Title,
		Description: req.Description,
	}
}
