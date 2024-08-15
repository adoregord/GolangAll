package repository

import (
	"bankRaceCondition/internal/domain"
	"fmt"
	"sync"
)

type UserRepoInterface interface {
	CreateUser(user *domain.User) domain.User
	Deposit(userId int, amount float64, wg *sync.WaitGroup) error
	Withdraw(userId int, amount float64, wg *sync.WaitGroup) error
}

type UserRepo struct {
	User  map[int]domain.User
	MUTEK *sync.Mutex
}

func NewUserRepo() UserRepoInterface {
	return &UserRepo{
		User:  make(map[int]domain.User),
		MUTEK: &sync.Mutex{},
	}
}
func (repo UserRepo) CreateUser(user *domain.User) domain.User {
	repo.User[user.ID] = *user
	return *user
}

func (repo UserRepo) Deposit(userId int, amount float64, wg *sync.WaitGroup) error {
	defer func() {
		repo.MUTEK.Unlock()
		wg.Done()
	}()
	repo.MUTEK.Lock()
	user1 := repo.User[userId]

	user1.Balance += amount
	repo.User[userId] = user1
	fmt.Println("Deposit: ",
		repo.User[userId].Balance, user1)
	return nil

}

func (repo UserRepo) Withdraw(userId int, amount float64, wg *sync.WaitGroup) error {
	defer func() {
		repo.MUTEK.Unlock()
		wg.Done()
	}()
	repo.MUTEK.Lock()
	user1 := repo.User[userId]

	user1.Balance -= amount
	repo.User[userId] = user1
	fmt.Println("Withdraw: ", repo.User[userId].Balance, user1)
	return nil

}
