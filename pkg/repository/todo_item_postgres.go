package repository

import (
	"fmt"
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) CreateNew(listId int, input todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()

	var itemId int
	createTodoItemQuery := fmt.Sprintf("INSERT INTO %s (title, description, done) VALUES ($1, $2, $3) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createTodoItemQuery, input.Title, input.Description, input.Done)

	err = row.Scan(&itemId)
	if err != nil {
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAllByListIdAndUserId(listId int, userId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s AS ti LEFT JOIN %s AS il ON il.item_id = ti.id LEFT JOIN %s AS ul ON ul.list_id = il.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Select(&items, query, userId, listId)
	if err != nil {
		return nil, err
	}

	return items, nil
}
