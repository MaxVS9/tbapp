package service

import (
	"restapi"
	"restapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user restapi.User) (int, error)
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
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewListService(repos.TodoList),
	}
}
