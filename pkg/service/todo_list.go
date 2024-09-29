package service

import (
	"restapi"
	"restapi/pkg/repository"
)

type ListService struct {
	repo repository.TodoList
}

func NewListService(repo repository.TodoList) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) Create(list restapi.TodoList) (int, error) {
	return 0, nil
}
