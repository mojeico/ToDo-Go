package service

import "myGoMod/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
