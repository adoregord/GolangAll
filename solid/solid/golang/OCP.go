package gosolid

import "fmt"

// open-closed principles
// type "Payment" is open for extension but closed for changes
// means you can extend or add new payment, without changing the current payment behaviour
// its possible becase Payment stuct has Payment Methods as interfaces
// its dynamic behaviour is best for development

type Payment struct{}
type PaymentMethod interface {
	Pay()
}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
}

type GOPAY struct {
	amount float64
}

func (cc GOPAY) Pay() {
	fmt.Printf("Paid %.2f using Gopay", cc.amount)
}

type OVO struct {
	amount float64
}

func (cc OVO) Pay() {
	fmt.Printf("Paid %.2f using OVO", cc.amount)
}

func main() {
	p := Payment{}
	gopay := GOPAY{15000.00}
	ovo := OVO{25000.00}
	p.Process(ovo)
	p.Process(gopay)

	var userDataSource UserDataSource

	// Contoh menggunakan LocalDatabase
	userDataSource = &LocalDatabase{}
	user, err := userDataSource.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", user)
	}

	// Contoh menggunakan API
	userDataSource = &API{}
	user, err = userDataSource.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", user)
	}

}
