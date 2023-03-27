package repository

import (
	"fmt"
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{
		db: db,
	}
}

func (repo *ListPostgres) CreateNew(list todo.TodoList, userId int) (int, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ListPostgres) GetAllByUserId(userId int) ([]todo.TodoList, error) {
	var list []todo.TodoList
	getUsersListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s as ul LEFT JOIN %s as tl ON ul.list_id = tl.id WHERE ul.user_id = $1", usersListsTable, todoListsTable)

	err := r.db.Select(&list, getUsersListsQuery, userId)

	return list, err
}

func (r *ListPostgres) GetListById(listId int, userId int) (todo.TodoList, error) {
	var list todo.TodoList
	getListByIdQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s AS ul LEFT JOIN %s tl ON ul.list_id = tl.id WHERE ul.user_id = $1 AND ul.list_id = $2", usersListsTable, todoListsTable)

	err := r.db.Get(&list, getListByIdQuery, userId, listId)

	return list, err
}

func (r *ListPostgres) UpdateList(listId int, userId int, input todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, listId, userId)

	logrus.Debugf("Update Query: %s", query)
	logrus.Debugf("Params: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ListPostgres) DeleteById(listId int) error {
	deleteListByIdQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", todoListsTable)
	_, err := r.db.Exec(deleteListByIdQuery, listId)

	return err
}
