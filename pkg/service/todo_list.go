package service

import (
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/DmitryYegorov/go-todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateNewTodoList(todo todo.TodoList, userId int) (int, error) {
	return s.repo.CreateNew(todo, userId)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAllByUserId(userId)
}

func (s *TodoListService) GetTodoListById(listId int, userId int) (todo.TodoList, error) {
	return s.repo.GetListById(listId, userId)
}

func (s *TodoListService) DeleteListById(listId int, userId int) error {
	_, err := s.repo.GetListById(listId, userId)
	if err != nil {
		return err
	}
	return s.repo.DeleteById(listId)
}
