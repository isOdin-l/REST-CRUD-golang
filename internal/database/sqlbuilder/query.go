package sqlbuilder

import (
	"isOdin/RestApi/internal/database"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type SqlBuilder struct {
	sql sq.StatementBuilderType
}

func NewSqlBuilder() *SqlBuilder {
	return &SqlBuilder{
		sql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// ======== AUTH ========
func (builder *SqlBuilder) InsertUser(values ...any) (string, []interface{}, error) {
	return builder.sql.
		Insert(database.TableUsers).
		Columns("id", "name", "username", "password_hash").
		Values(values...).ToSql()
}

func (builder *SqlBuilder) SelectUser(userId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Select("name", "username", "password_hash").
		From(database.TableUsers).
		Where(sq.Eq{"id": userId}).ToSql()
}

// ======== ITEM ========
func (builder *SqlBuilder) InsertItem(values ...any) (string, []interface{}, error) {
	return builder.sql.
		Insert(database.TableItems).
		Columns("id", "list_id", "title", "description", "done").
		Values(values...).ToSql()
}

func (builder *SqlBuilder) SelectItem(itemId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Select("list_id", "title", "description", "done").
		From(database.TableItems).
		Where(sq.Eq{"id": itemId}).ToSql()
}

func (builder *SqlBuilder) UpdateItem(itemId uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error) {
	return builder.sql.
		Update(database.TableItems).
		SetMap(updateData).
		Where(sq.Eq{"id": itemId}).
		Suffix("RETURNING author_id, title, description").ToSql()
}

func (builder *SqlBuilder) DeleteItem(itemId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Delete(database.TableItems).
		Where(sq.Eq{"id": itemId}).ToSql()
}

// ======== LIST ========
func (builder *SqlBuilder) InsertList(values ...any) (string, []interface{}, error) {
	return builder.sql.
		Insert(database.TableLists).
		Columns("id", "author_id", "title", "description").
		Values(values...).ToSql()
}

func (builder *SqlBuilder) SelectList(listId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Select("author_id", "title", "description").
		From(database.TableLists).
		Where(sq.Eq{"id": listId}).ToSql()
}

func (builder *SqlBuilder) UpdateList(author_id uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error) {
	return builder.sql.
		Update(database.TableLists).
		SetMap(updateData).
		Where(sq.Eq{"author_id": author_id}).
		Suffix("RETURNING list_id, title, description, done").ToSql()
}

func (builder *SqlBuilder) DeleteList(listId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Delete(database.TableLists).
		Where(sq.Eq{"list_id": listId}).ToSql()
}
