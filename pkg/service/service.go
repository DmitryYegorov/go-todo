package service

import "github.com/DmitryYegorov/go-todo/pkg/repository"

type Authorization interface{}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
