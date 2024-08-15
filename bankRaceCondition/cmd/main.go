package main

import (
	"bankRaceCondition/internal/domain"
	"bankRaceCondition/internal/handler"
	"bankRaceCondition/internal/repository"
	"bankRaceCondition/internal/usecase"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(5)

	var wg sync.WaitGroup

	userRepo := repository.NewUserRepo()
	userUc := usecase.NewUserUsecase(userRepo)
	userH := handler.NewUserHandler(userUc)

	user := domain.User{
		ID:      1,
		Name:    "Didi",
		Balance: 2000,
		MUTEK:   &sync.Mutex{},
	}
	a := userH.CreateUser(user)

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go userH.Deposit(a.ID, 100, &wg)
		go userH.Withdraw(a.ID, 100, &wg)
	}
	wg.Wait()
	fmt.Println("Program Berhasil Dijalankan")
}

// type User struct {
// 	ID      int     `json:"id"`
// 	Name    string  `json:"name"`
// 	Balance float64 `json:"balance"`
// 	MUTEK   *sync.Mutex
// }

// func CreateUser() *User {
// 	return &User{
// 		ID:      1,
// 		Name:    "Didi",
// 		Balance: 1000,
// 		MUTEK:   &sync.Mutex{},
// 	}
// }

// func (u *User) Withdraw(amount float64, wg *sync.WaitGroup) error {
// 	u.MUTEK.Lock()
// 	defer wg.Done()
// 	defer u.MUTEK.Unlock()
// 	if u.Balance < amount {
// 		return fmt.Errorf("insufficient balance")
// 	}
// 	u.Balance -= amount
// 	fmt.Println(u.Balance)
// 	return nil
// }
// func (u *User) Deposit(amount float64, wg *sync.WaitGroup) error {
// 	u.MUTEK.Lock()
// 	defer wg.Done()
// 	defer u.MUTEK.Unlock()
// 	u.Balance += amount
// 	fmt.Println(u.Balance)
// 	return nil
// }

// func main() {
// 	runtime.GOMAXPROCS(2)

// 	var wg sync.WaitGroup

// 	user := CreateUser()
// 	for i := 0; i < 100; i++ {
// 		wg.Add(2)
// 		go user.Deposit(100, &wg)
// 		go user.Withdraw(100, &wg)
// 		// go func() {
// 		// 	defer wg.Done()
// 		// 	user.Deposit(100)
// 		// }()
// 		// go func() {
// 		// 	defer wg.Done()
// 		// 	user.Withdraw(100)
// 		// }()

// 	}
// 	// wg.Done()

// 	wg.Wait()
// 	print("SUCCESS PRINT")
// }
