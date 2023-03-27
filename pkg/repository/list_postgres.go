package repository

import (
	"fmt"
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/jmoiron/sqlx"
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
