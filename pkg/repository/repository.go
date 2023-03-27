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
}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewListPostgres(db),
	}
}
