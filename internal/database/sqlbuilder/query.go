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

func (builder *SqlBuilder) GetUserByUsernameAndPassword(user *models.User) (string, []interface{}, error){
	return builder.sql.
		Select("id", "name", "username", "password_hash").
		From(database.TableUsers).
		Where(sq.Eq{
			"username": user.Username,
			"password_hash": user.Password_hash,
		}).ToSql()
}

func (builder *SqlBuilder) SelectUser(userId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Select("name", "username", "password_hash").
		From(database.TableUsers).
		Where(sq.Eq{"id": userId.String()}).ToSql()
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
		item.Id.String(),
		item.List_id.String(),
		item.Title,
		item.Description,
		item.Done,
		userId.String(),
	}

	return query, args, nil
}

func (builder *SqlBuilder) SelectItem(itemId, userId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Select("item.list_id", "item.title", "item.description", "item.done").
		From(fmt.Sprintf("%s item", database.TableItems)).
		Join(fmt.Sprintf("%s list ON list.id = item.list_id", database.TableLists)).
		Where(sq.Eq{
			"item.id":        itemId.String(),
			"list.author_id": userId.String(),
		}).ToSql()
}

func (builder *SqlBuilder) UpdateItem(itemId, userId, listId uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error) {
	return builder.sql.
		Update(database.TableItems).
		SetMap(updateData).
		From(database.TableLists).
		Where(sq.Expr(fmt.Sprintf("%s.id = %s.list_id", database.TableLists, database.TableItems))).
		Where(sq.Eq{
			fmt.Sprintf("%s.author_id", database.TableLists): userId.String(),
			fmt.Sprintf("%s.id", database.TableItems): itemId.String(),
		}).
		Suffix("RETURNING items.title, items.description, items.Done").ToSql()
}

func (builder *SqlBuilder) DeleteItem(itemId, userId uuid.UUID) (string, []interface{}, error) {

	query := `
		DELETE FROM items
		USING lists
		WHERE items.list_id = lists.id
		AND items.id = $1
		AND lists.author_id = $2
	`

	return query, []any{itemId, userId}, nil
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
			"id":        listId.String(),
			"author_id": userId.String(),
		}).ToSql()
}

func (builder *SqlBuilder) UpdateList(author_id, listId uuid.UUID, updateData map[string]interface{}) (string, []interface{}, error) {
	return builder.sql.
		Update(database.TableLists).
		SetMap(updateData).
		Where(sq.Eq{
			"author_id": author_id.String(),
			"id":        listId.String(),
		}).
		Suffix("RETURNING title, description").ToSql()
}

func (builder *SqlBuilder) DeleteList(listId, userId uuid.UUID) (string, []interface{}, error) {
	return builder.sql.
		Delete(database.TableLists).
		Where(sq.Eq{
			"id":        listId.String(),
			"author_id": userId.String(),
		}).ToSql()
}
