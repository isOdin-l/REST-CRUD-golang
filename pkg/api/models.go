package api

import "github.com/google/uuid"

// ========================
// ==== REQUEST MODELS ====
// ========================

// Authentication
type SignUp struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Item
type CreateItem struct {
	ListId      uuid.UUID `param:"list_id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
}

type GetItem struct {
	ItemId uuid.UUID `param:"item_id"`
}

type UpdateItem struct {
	ItemId      uuid.UUID `param:"item_id"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Done        *bool     `json:"done"`
}

type DeleteItem struct {
	ItemId uuid.UUID `param:"item_id"`
}

// List
type CreateList struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetList struct {
	ListId uuid.UUID `param:"list_id"`
}

type UpdateList struct {
	ListId      uuid.UUID `param:"list_id"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
}

type DeleteList struct {
	ListId uuid.UUID `param:"list_id"`
}

// =========================
// ==== RESPONSE MODELS ====
// =========================

// Authentication
type ResponseJwtToken struct {
	Token string `json:"token"`
}

// Item
type ResponseItem struct {
	Item struct {
		ItemId      uuid.UUID `json:"item_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Done        bool      `json:"done"`
	} `json:"item"`
}

// List
type ResponseList struct {
	List struct {
		ListId      uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		//Items       []ResponseItem `json:"items"`
	} `json:"list"`
}

// Error
type ResponseError struct {
	Error struct {
		HttpCode int    `json:"code"`
		Message  string `json:"message"`
	}
}
