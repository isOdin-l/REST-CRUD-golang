package database

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

// ======== AUTH ========
func InsertUser(builder *sq.StatementBuilderType, values ...any) (string, []interface{}, error) {
	return builder.
		Insert(TableUsers).
		Columns("id", "name", "username", "password_hash").
		Values(values...).ToSql()
}

func SelectUser(builder *sq.StatementBuilderType, userId uuid.UUID) (string, []interface{}, error) {
	return builder.
		Select("name", "username", "password_hash").
		From(TableUsers).
		Where(sq.Eq{"id": userId}).ToSql()
}

// ======== ITEM ========
func InsertItem(builder *sq.StatementBuilderType, values ...any) (string, []interface{}, error) {
	return builder.
		Insert(TableItems).
		Columns("id", "list_id", "title", "description", "done").
		Values(values...).ToSql()
}

func SelectItem(builder *sq.StatementBuilderType, itemId uuid.UUID) (string, []interface{}, error) {
	return builder.
		Select("list_id", "title", "description", "done").
		From(TableItems).
		Where(sq.Eq{"id": itemId}).ToSql()
}

func UpdateItem(builder *sq.StatementBuilderType, list_id uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error) {
	return builder.
		Update(TableItems).
		SetMap(updateData).
		Where(sq.Eq{"list_id": list_id}).
		Suffix("RETURNING author_id, title, description").ToSql()
}

func DeleteItem(builder *sq.StatementBuilderType, item_id uuid.UUID) (string, []interface{}, error) {
	return builder.
		Delete(TableItems).
		Where(sq.Eq{"item_id": item_id}).ToSql()
}

// ======== LIST ========
func InsertList(builder *sq.StatementBuilderType, values ...any) (string, []interface{}, error) {
	return builder.
		Insert(TableLists).
		Columns("id", "author_id", "title", "description").
		Values(values...).ToSql()
}

func SelectList(builder *sq.StatementBuilderType, listId uuid.UUID) (string, []interface{}, error) {
	return builder.
		Select("author_id", "title", "description").
		From(TableLists).
		Where(sq.Eq{"id": listId}).ToSql()
}

func UpdateList(builder *sq.StatementBuilderType, author_id uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error) {
	return builder.
		Update(TableLists).
		SetMap(updateData).
		Where(sq.Eq{"author_id": author_id}).
		Suffix("RETURNING list_id, title, description, done").ToSql()
}

func DeleteList(builder *sq.StatementBuilderType, list_id uuid.UUID) (string, []interface{}, error) {
	return builder.
		Delete(TableLists).
		Where(sq.Eq{"list_id": list_id}).ToSql()
}
