package gosolid

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserReposiory interface {
	Create(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, user User) error
}

type DBUserRepository struct {
	db *gorm.DB
}

func (r *DBUserRepository) Create(ctx context.Context, user User) (*User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	return &user, err
}

func (r *DBUserRepository) Update(ctx context.Context, user User) error {
	return r.db.WithContext(ctx).Save(&user).Error
}

type UserCache interface {
	Create(user User) (*User, error)
	Update(user User) error
}

type MemoryUserRepository struct {
	users map[uuid.UUID]User
}

func (r *MemoryUserRepository) Create(user User) (*User, error) {
	if r.users == nil {
		r.users = map[uuid.UUID]User{}
	}
	if newUUID, err := uuid.NewV4(); err != nil {
		return nil, err
	} else {
		user.ID = newUUID
	}

	r.users[user.ID] = user

	return &user, nil
}

func (r *MemoryUserRepository) Update(user User) error {
	if r.users == nil {
		r.users = map[uuid.UUID]User{}
	}
	r.users[user.ID] = user

	return nil
}
