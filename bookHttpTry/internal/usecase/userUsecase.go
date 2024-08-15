package usecase

import (
	"httpBook/internal/domain"
	"httpBook/internal/repository"
)

type UserUsecaseInterface interface {
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

type UserUsecase struct {
	UserRepo repository.UserRepoInterface
}

func NewUserRepo(userRepo repository.UserRepoInterface) UserUsecaseInterface {
	return UserUsecase{
		UserRepo: userRepo,
	}
}

func (uc UserUsecase) UserAdd(user domain.User) error {
	return uc.UserRepo.UserAdd(&user)
}
func (uc UserUsecase) UserDelete(userId int) error {
	return uc.UserRepo.UserDelete(userId)
}
func (uc UserUsecase) UserView() ([]domain.User, error) {
	return uc.UserRepo.UserView()
}
func (uc UserUsecase) UserViewById(userId int) (domain.User, error) {
	return uc.UserRepo.UserViewById(userId)
}