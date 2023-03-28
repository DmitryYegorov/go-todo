package repository

import (
	"fmt"
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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

func (r *TodoItemPostgres) UpdateListItem(itemId int, input todo.UpdateTodoItem) error {
	var setValues = make([]string, 0)
	var args = make([]interface{}, 0)
	var argId = 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done = $%d", argId))
		if *input.Done {
			args = append(args, "true")
		} else {
			args = append(args, "false")
		}
		argId++
	}

	args = append(args, itemId)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", todoItemsTable, strings.Join(setValues, ", "), argId)

	logrus.Debugf("Update Query %s", query)
	logrus.Debugf("Params: %v", args)

	_, err := r.db.Exec(query, args...)

	return err
}
