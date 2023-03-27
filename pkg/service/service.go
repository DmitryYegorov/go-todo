package service

import (
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/DmitryYegorov/go-todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateNewTodoList(todo todo.TodoList, userId int) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetTodoListById(listId int, userId int) (todo.TodoList, error)
	DeleteListById(listId int, userId int) error
	UpdateList(listId int, userId int, input todo.UpdateListInput) error
}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(*repo),
		TodoList:      NewTodoListService(*repo),
	}
}
