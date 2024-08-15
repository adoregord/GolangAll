package repository

import (
	"basic-solid/internal/domain"
	"errors"
	"fmt"
)

// ini interface untuk menampung semua interface yang ingin digunakan
type UserDatasource interface {
	Save(user *domain.User) error
	Find(userID int) (*domain.User, error)
	Delete(userID int) error
}

// ini ceritanya db
type UserRepository struct {
	users map[int]domain.User
}

// ini ceritanya dari oauth server
type UserOauthServer struct {
	users map[int]domain.User
}

// ini untuk koneksi ke db nya untuk ngubah/nampil data
func NewUserRepository() UserDatasource {
	return &UserRepository{
		users: make(map[int]domain.User),
	}
}

// ini buatan bene aku gatau -DD
func NewUserOauthServer() UserDatasource {
	return &UserOauthServer{}
}

// ini fungsi dari repo untuk menyimpan data pada DB User ^____^❤️😊😂🤣😍
func (repo *UserRepository) Save(user *domain.User) error {
	if _, exists := repo.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	repo.users[user.ID] = *user
	return nil
}

// ini fungsi dari repo untuk mencari data user by ID pada DB User
func (repo *UserRepository) Find(userID int) (*domain.User, error) {
	if user, exists := repo.users[userID]; exists {
		return &user, nil
	}

	return nil, errors.New("user not found")
}

// ini fungsi dari repo untuk menghapus data user by ID pada DB User
func (repo *UserRepository) Delete(userID int) error {
	if _, exists := repo.users[userID]; !exists {
		return errors.New("user not found")
	}

	delete(repo.users, userID)
	return nil
}

// Oauth Server Implementation
func (repo *UserOauthServer) Save(user *domain.User) error {
	fmt.Println("Saving user to oauth server...")
	return nil
}

// ini punya bene aku gatau
func (repo *UserOauthServer) Find(userID int) (*domain.User, error) {
	if user, exists := repo.users[userID]; exists {
		return &user, nil
	}

	return nil, errors.New("user not found")
}

// ini punya bene aku gatau
func (repo *UserOauthServer) Delete(userID int) error {
	return nil
}
