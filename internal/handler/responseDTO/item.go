package responseDTO

import "github.com/google/uuid"

type ItemResponse struct {
	ItemId      uuid.UUID `json:"item_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}

// create converter item -> response
