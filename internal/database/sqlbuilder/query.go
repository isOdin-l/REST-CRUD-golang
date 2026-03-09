package sqlbuilder

import (
	"fmt"
	"isOdin/RestApi/internal/database"
	"isOdin/RestApi/internal/repository/models"

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
func (builder *SqlBuilder) InsertItem(userId uuid.UUID, item *models.Item) (string, []interface{}, error) {
	query := fmt.Sprintf(`
	INSERT INTO %s (id, list_id, title, description, done)
	SELECT $1, $2, $3, $4, $5
	WHERE EXISTS (
		SELECT 1
		FROM %s
		WHERE id = $2 AND author_id = $6
	)`, database.TableItems, database.TableLists)

	args := []interface{}{
		item.Id,
		item.List_id,
		item.Title,
		item.Description,
		item.Done,
		userId,
	}

	return query, args, nil
}

func (builder *SqlBuilder) SelectItem(itemId, userId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Select("item.list_id", "item.title", "item.description", "item.done").
		From(fmt.Sprintf("%s item", database.TableItems)).
		Join(fmt.Sprintf("%s, list ON list.id = item.list_id", database.TableLists)).
		Where(sq.Eq{
			"item.id":        itemId,
			"list.author_id": userId,
		}).ToSql()
}

func (builder *SqlBuilder) UpdateItem(itemId, userId uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error) {
	return builder.sql.
		Update(fmt.Sprintf("%s item", database.TableItems)).
		SetMap(updateData).
		From(fmt.Sprintf("%s list", database.TableLists)).
		Where(sq.Eq{
			"item.id":        itemId,
			"list.author_id": userId,
			"item.list_id":   "list.id",
		}).
		Suffix("RETURNING item.author_id, item.title, item.description").ToSql()
}

func (builder *SqlBuilder) DeleteItem(itemId, userId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Delete(fmt.Sprintf("%s item", database.TableItems)).
		Where(sq.Eq{"id": itemId}).ToSql()
}

// ======== LIST ========
func (builder *SqlBuilder) InsertList(values ...any) (string, []interface{}, error) {
	return builder.sql.
		Insert(database.TableLists).
		Columns("id", "author_id", "title", "description").
		Values(values...).ToSql()
}

func (builder *SqlBuilder) SelectList(listId, userId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Select("author_id", "title", "description").
		From(database.TableLists).
		Where(sq.Eq{
			"id":        listId,
			"author_id": userId,
		}).ToSql()
}

func (builder *SqlBuilder) UpdateList(author_id, listId uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error) {
	return builder.sql.
		Update(database.TableLists).
		SetMap(updateData).
		Where(sq.Eq{
			"author_id": author_id,
			"id":        listId,
		}).
		Suffix("RETURNING id, title, description, done").ToSql()
}

func (builder *SqlBuilder) DeleteList(listId, userId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Delete(database.TableLists).
		Where(sq.Eq{
			"id":        listId,
			"author_id": userId,
		}).ToSql()
}
