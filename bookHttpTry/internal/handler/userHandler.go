package handler

import (
	"httpBook/internal/domain"
	"httpBook/internal/usecase"
)

type UserHandlerInterface interface {
	UserAdd
	UserDelete
	UserView
	UserViewById
}

type UserAdd interface {
	UserAdd(user domain.User) error
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

type UserHandler struct{
	UserUsecase usecase.UserUsecaseInterface
}

func NewUserHandler(userUsecase usecase.UserUsecaseInterface) UserHandlerInterface {
	return UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h UserHandler) UserAdd(user domain.User) error {
	return h.UserUsecase.UserAdd(user)
}
func (h UserHandler) UserDelete(userId int) error {
	return h.UserUsecase.UserDelete(userId)
}
func (h UserHandler) UserView() ([]domain.User, error) {
	return h.UserUsecase.UserView()
}
func (h UserHandler) UserViewById(userId int) (domain.User, error) {
	return h.UserUsecase.UserViewById(userId)
}





