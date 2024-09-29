package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"restapi"
)

type ListPostgres struct {
	db *sqlx.DB
}

func newListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (r *ListPostgres) Create(list restapi.TodoList) (int, error) {
	tx, err := r.db.Begin()

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (list_id) VALUES ($1)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
