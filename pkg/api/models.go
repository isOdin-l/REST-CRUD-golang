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
	UserId   uuid.UUID
	Username string `json:"username"`
	Password string `json:"password"`
}

// Item
type CreateItem struct {
	ListId      uuid.UUID `json:"list_id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
}

type GetItem struct {
	ItemId uuid.UUID `json:"item_id"`
}

type UpdateItem struct {
	ItemId      uuid.UUID `json:"item_id"`
	Title       string    `json:"title" validate:"optional"`
	Description string    `json:"description" validate:"optional"`
	Done        bool      `json:"done" validate:"optional"`
}

type DeleteItem struct {
	ItemId uuid.UUID `json:"item_id"`
}

// List
type CreateList struct {
	UserId      uuid.UUID
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetList struct {
	ListId uuid.UUID `json:"list_id"`
}

type UpdateList struct {
	UserId      uuid.UUID
	ListId      uuid.UUID `json:"list_id"`
	Title       string    `json:"title" validate:"optional"`
	Description string    `json:"description" validate:"optional"`
	Done        bool      `json:"done" validate:"optional"`
}

type DeleteList struct {
	ListId uuid.UUID `json:"list_id"`
}

// =========================
// ==== RESPONSE MODELS ====
// =========================

// Authentication
type ResponseSignUp struct {
	User struct {
		UserId   uuid.UUID `json:"id"`
		Username string    `json:"name"`
	} `json:"user"`
}

type ResponseSignIn struct {
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
		ListId      uuid.UUID      `json:"id"`
		Title       string         `json:"title"`
		Description string         `json:"description"`
		Items       []ResponseItem `json:"items"`
	}
}
