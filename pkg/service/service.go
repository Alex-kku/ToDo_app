package service

import "github.com/Alex-kku/todo-app/pkg/repository"

type Authorization interface {
}

type Todolist interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	Todolist
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
