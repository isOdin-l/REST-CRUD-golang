package entities

import "github.com/google/uuid"

type User struct {
	UserId   uuid.UUID
	Name     string
	Username string
	Password string
}

type List struct {
	UserId      uuid.UUID
	ListId      uuid.UUID
	Title       string
	Description string
}

type UpdateList struct {
	UserId    uuid.UUID
	ListId    uuid.UUID
	OptValues struct {
		Title       *string
		Description *string
	}
}

type Item struct {
	UserId      uuid.UUID
	ListId      uuid.UUID
	ItemId      uuid.UUID
	Title       string
	Description string
	Done        bool
}

type UpdateItem struct {
	UserId    uuid.UUID
	ListId    uuid.UUID
	ItemId    uuid.UUID
	OptValues struct {
		Title       *string
		Description *string
		Done        *bool
	}
}
