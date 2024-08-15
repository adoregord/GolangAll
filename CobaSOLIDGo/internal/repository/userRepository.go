package repository

import "main.go/internal/domain"

// bikin db
type UserRepository struct {
	User map[int]domain.User
}

//bikin repo untuk koneksi ke db
func NewUserRepo() UserRepoInterface {
	return UserRepository{
		User: map[int]domain.User{},
	}
}

type UserRepoInterface interface {
	UserAdd
	UserDelete
	UserFindAll
}

type UserAdd interface {
	Add() error
}

type UserDelete interface {
	Delete() error
}

type UserFindAll interface {
	FindAll() error
}

func (repo UserRepository) Add() error {
	return nil
}

func (repo UserRepository) Delete() error {
	return nil
}

func (repo UserRepository) FindAll() error {
	return nil
}
