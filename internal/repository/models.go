package repository

import "github.com/google/uuid"

type item struct {
	ItemId      uuid.UUID `db:"id"`
	ListId      uuid.UUID `db:"list_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Done        bool      `db:"done"`
}

func getInsertItemColumns() []string {
	return []string{"id", "list_id", "title", "description"}
}
