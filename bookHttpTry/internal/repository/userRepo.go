package repository

import (
	"errors"
	"httpBook/internal/domain"
)

type UserRepoInterface interface {
	UserAdd
	UserDelete
	UserView
	UserViewById
}
type UserAdd interface {
	UserAdd(user *domain.User) error
}
type UserDelete interface {
	UserDelete(userId int) error
}
type UserView interface {
	UserView() ([]domain.User, error)
}
type UserViewById interface {
	UserViewById(userId int) (domain.User, error)
}

type UserRepo struct {
	User map[int]domain.User
}

func NewUserRepo() UserRepoInterface {
	return UserRepo{
		User: map[int]domain.User{},
	}
}

func (repo UserRepo) UserAdd(user *domain.User) error {
	if len(repo.User) == 0 || repo.User == nil {
		user.ID = 1
	} else {
		user.ID = repo.User[len(repo.User)-1].ID + 1
	}
	repo.User[user.ID] = *user
	return nil
}
func (repo UserRepo) UserDelete(userId int) error {
	if _, exist := repo.User[userId]; !exist {
		return errors.New("THERE'S NO USER WITH SUCH ID 🤬🤬🤬")
	}

	delete(repo.User, userId)
	return nil
}
func (repo UserRepo) UserView() ([]domain.User, error) {
	var users []domain.User
	for _, user := range repo.User {
		users = append(users, user)
	}
	return users, nil
}
func (repo UserRepo) UserViewById(userId int) (domain.User, error) {
	var user domain.User
	return user, nil
}
