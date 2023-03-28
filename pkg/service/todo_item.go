package service

import (
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/DmitryYegorov/go-todo/pkg/repository"
)

type TodoItemService struct {
	todolistRepo repository.TodoList
	todoItemRepo repository.TodoItem
}

func NewTodoItemService(todoListRepo repository.TodoList, todoItemRepo repository.TodoItem) *TodoItemService {
	return &TodoItemService{todolistRepo: todoListRepo, todoItemRepo: todoItemRepo}
}

func (s *TodoItemService) CreateNewItem(listId int, userId int, input todo.TodoItem) (int, error) {
	_, err := s.todolistRepo.GetListById(listId, userId)

	if err != nil {
		return 0, err
	}

	return s.todoItemRepo.CreateNew(listId, input)
}

func (s *TodoItemService) GetAllItemByListId(listId int, userId int) ([]todo.TodoItem, error) {
	items, err := s.todoItemRepo.GetAllByListIdAndUserId(listId, userId)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *TodoItemService) UpdateItem(itemId int, input todo.UpdateTodoItem) error {
	return s.todoItemRepo.UpdateListItem(itemId, input)
}

func (s *TodoItemService) GetItemById(itemId int) (todo.TodoItem, error) {
	return s.todoItemRepo.GetItemById(itemId)
}
