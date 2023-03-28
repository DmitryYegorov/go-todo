package repository

import (
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUserByUsername(username string) (todo.User, error)
}

type TodoList interface {
	CreateNew(list todo.TodoList, userId int) (int, error)
	GetAllByUserId(userId int) ([]todo.TodoList, error)
	GetListById(listId int, userId int) (todo.TodoList, error)
	DeleteById(listId int) error
	UpdateList(listId int, userId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	CreateNew(listId int, input todo.TodoItem) (int, error)
	GetAllByListIdAndUserId(listId int, userId int) ([]todo.TodoItem, error)
	UpdateListItem(itemId int, input todo.UpdateTodoItem) error
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
