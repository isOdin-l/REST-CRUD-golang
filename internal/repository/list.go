package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/database"
	"github.com/isOdin/RestApi/internal/repository/requestDTO"
	"github.com/isOdin/RestApi/internal/repository/responseDTO"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoListRepository struct {
	db *pgxpool.Pool
}

func NewTodoListRepository(db *pgxpool.Pool) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) CreateList(listInfo *requestDTO.CreateList) (uuid.UUID, error) { // UUID, error
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return uuid.Nil, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", database.TableTodoLists)
	rowCreateList := tx.QueryRow(context.Background(), createListQuery, listInfo.Title, listInfo.Description)

	var listId uuid.UUID
	if errScan := rowCreateList.Scan(&listId); errScan != nil {
		tx.Rollback(context.Background())
		return uuid.Nil, errScan
	}

	createUserListRelationQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", database.TableUsersLists)
	_, errExec := tx.Exec(context.Background(), createUserListRelationQuery, listInfo.UserId, listId)
	if errExec != nil {
		tx.Rollback(context.Background())
		return uuid.Nil, errExec
	}

	return listId, tx.Commit(context.Background())
}

func (r *TodoListRepository) GetAllLists(userId uuid.UUID) (*[]responseDTO.GetList, error) {
	lists := make([]responseDTO.GetList, 0)

	getAllListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", database.TableTodoLists, database.TableUsersLists)
	rowsGetAllLists, err := r.db.Query(context.Background(), getAllListsQuery, userId)
	if err != nil {
		return nil, err
	}

	lists, err = pgx.CollectRows(rowsGetAllLists, pgx.RowToStructByName[responseDTO.GetList])

	return &lists, err
}

func (r *TodoListRepository) GetListByIdAndUserId(ctx context.Context, listId uuid.UUID, userId uuid.UUID) (*responseDTO.GetListById, error) {
	var list responseDTO.GetListById

	getListByIdQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", database.TableTodoLists, database.TableUsersLists)
	err := r.db.QueryRow(ctx, getListByIdQuery, userId, listId).Scan(&list.Id, &list.Title, &list.Description)

	return &list, err
}

func (r *TodoListRepository) DeleteList(listInfo *requestDTO.DeleteList) error {
	queryDeleteList := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", database.TableTodoLists, database.TableUsersLists)
	_, err := r.db.Exec(context.Background(), queryDeleteList, listInfo.UserId, listInfo.ListId)

	return err
}

func (r *TodoListRepository) UpdateList(listInfo *requestDTO.UpdateList) error {
	queryUpdateList := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d", database.TableTodoLists, listInfo.SetValuesQuery, database.TableUsersLists, listInfo.ArgId, listInfo.ArgId+1)

	_, err := r.db.Exec(context.Background(), queryUpdateList, *listInfo.SetArgs...)

	return err
}
