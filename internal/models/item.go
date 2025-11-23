package models

import "github.com/google/uuid"

type Item struct {
	ItemId      uuid.UUID
	Title       string
	Description string
	Done        bool
}

type CreateItemParams struct {
	ItemId      uuid.UUID
	UserId      uuid.UUID
	ListId      uuid.UUID
	Title       string
	Description string
}
