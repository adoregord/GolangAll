package usecase

import (
	"bankRaceCondition/internal/domain"
	"bankRaceCondition/internal/repository"
	"sync"
)

type UserUsecaseInterface interface {
	CreateUser(user domain.User) domain.User
	Deposit(userId int, amount float64, wg *sync.WaitGroup) error
	Withdraw(userId int, amount float64, wg *sync.WaitGroup) error
}

type UserUsecase struct {
	UserRepository repository.UserRepoInterface
}

func NewUserUsecase(userRepository repository.UserRepoInterface) UserUsecaseInterface {
	return UserUsecase{
		UserRepository: userRepository,
	}
}

func (uc UserUsecase) CreateUser(user domain.User) domain.User {
	return uc.UserRepository.CreateUser(&user)
}
func (uc UserUsecase) Deposit(userId int, amount float64, wg *sync.WaitGroup) error {
	return uc.UserRepository.Deposit(userId, amount, wg)
}
func (uc UserUsecase) Withdraw(userId int, amount float64, wg *sync.WaitGroup) error {
	return uc.UserRepository.Withdraw(userId, amount, wg)
}
