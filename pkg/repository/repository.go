package repository

import (
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUserByUsername(username string) (todo.User, error)
}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
