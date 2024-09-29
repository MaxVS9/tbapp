package service

import (
	"restapi"
	"restapi/pkg/repository"
)

type Authorization interface {
}

type TodoList interface {
	Create(list restapi.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoList: NewListService(repos.TodoList),
	}
}
