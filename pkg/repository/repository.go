package repository

import (
	"github.com/jmoiron/sqlx"
	"restapi"
)

type Authorization interface {
}

type TodoList interface {
	Create(list restapi.TodoList) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoList: newListPostgres(db),
	}
}
