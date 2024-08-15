package handler

import (
	"bankRaceCondition/internal/domain"
	"bankRaceCondition/internal/usecase"
	"sync"
)

type UserHandlerInterface interface {
	CreateUser(user domain.User) domain.User
	Deposit(userId int, amount float64, wg *sync.WaitGroup) error
	Withdraw(userId int, amount float64, wg *sync.WaitGroup) error
}

type UserHandler struct {
	UserUsercase usecase.UserUsecaseInterface
}

func NewUserHandler(userUsercase usecase.UserUsecaseInterface) UserHandlerInterface {
	return UserHandler{
		UserUsercase: userUsercase,
	}
}

func (h UserHandler) CreateUser(user domain.User) domain.User {
	return h.UserUsercase.CreateUser(user)
}
func (h UserHandler) Deposit(userId int, amount float64, wg *sync.WaitGroup) error {
	return h.UserUsercase.Deposit(userId, amount, wg)
}
func (h UserHandler) Withdraw(userId int,amount float64, wg *sync.WaitGroup) error {
	return h.UserUsercase.Withdraw(userId, amount, wg)
}
