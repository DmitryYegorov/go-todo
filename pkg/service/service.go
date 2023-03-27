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

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(*repo),
	}
}
