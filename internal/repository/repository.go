package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// type Authorization interface {
// 	CreateUser(user *requestDTO.CreateUser) (uuid.UUID, error)
// 	GetUser(user *requestDTO.GetUser) (*responseDTO.GetedUser, error)
// }

// type TodoList interface {
// 	CreateList(listInfo *requestDTO.CreateList) (uuid.UUID, error)
// 	GetAllLists(userId uuid.UUID) (*[]responseDTO.GetList, error)
// 	GetListById(listInfo *requestDTO.GetListById) (*responseDTO.GetListById, error)
// 	DeleteList(listInfo *requestDTO.DeleteList) error
// 	UpdateList(listInfo *requestDTO.UpdateList) error
// }

// type TodoItem interface {
// 	CreateItem(itemInfo *requestDTO.CreateItem) (uuid.UUID, error)
// 	GetAllItems(userId uuid.UUID) (*[]responseDTO.GetItem, error)
// 	GetItemById(itemInfo *requestDTO.GetItemById) (*responseDTO.GetItemById, error)
// 	DeleteItem(itemInfo *requestDTO.DeleteItem) error
// 	UpdateItem(itemInfo *requestDTO.UpdateItem) error
// }

type Repository struct {
	*AuthRepository
	*ListRepository
	*ItemRepository
}

// For future
type DatabaseInterface interface {
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
		ListRepository: NewListRepository(db),
		ItemRepository: NewItemRepository(db),
	}
}
