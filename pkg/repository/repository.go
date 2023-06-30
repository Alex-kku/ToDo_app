package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Todolist interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	Todolist
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
