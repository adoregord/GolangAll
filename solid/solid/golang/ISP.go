package gosolid

import (
	"errors"
	"time"
)

type Guest struct {
	cart ShoppingCart
}

type ShoppingCart struct {
	items []Product
}

type Product struct {
	Name  string
	Price float64
}

type Wallet struct {
	UserName string
	Total    float64
}

func (w *Wallet) Deduct(deductAmount float64) error {
	if w.Total-deductAmount < 0 {
		return errors.New("amount balance limit")
	} else {
		w.Total = w.Total - deductAmount
		return nil
	}
}

func (s *ShoppingCart) Add(product Product) {
	s.items = append(s.items, product)
}

func (g *Guest) AddToShoppingCart(product Product) {
	g.cart.Add(product)
}

func (g *Guest) IsLoggedIn() bool {
	return false
}

func (g *Guest) Pay(Money float64) error {
	return errors.New("user is not logged in")
}

func (g *Guest) HasPremium() bool {
	return false
}

func (g *Guest) HasDiscountFor(Product) bool {
	return false
}

type NormalCustomer struct {
	cart   ShoppingCart
	wallet Wallet
}

func (c *NormalCustomer) AddToShoppingCart(product Product) {
	c.cart.Add(product)
}

func (c *NormalCustomer) IsLoggedIn() bool {
	return true
}

func (c *NormalCustomer) Pay(Money float64) error {
	return c.wallet.Deduct(Money)
}

func (c *NormalCustomer) HasPremium() bool {
	return false
}

func (c *NormalCustomer) HasDiscountFor(Product) bool {
	return true
}

type Discount struct {
	Name         string
	EndValidDate time.Time
	Used         bool
}

type DiscountPolicy struct {
	AppliedProduct Product
	Discount
}

func (dp *DiscountPolicy) IsApplicableFor(c *PremiumCustomer, prod Product) bool {
	if c.HasDiscountFor(prod) {
		return true
	}
	if dp.AppliedProduct.Name == prod.Name {
		if time.Now().After(dp.EndValidDate) {
		} else {
			return false
		}
	} else {
		return false
	}
	return false
}

type PremiumCustomer struct {
	cart     ShoppingCart
	wallet   Wallet
	policies []DiscountPolicy
}

func (c *PremiumCustomer) AddToShoppingCart(product Product) {
	c.cart.Add(product)
}

func (c *PremiumCustomer) IsLoggedIn() bool {
	return true
}

func (c *PremiumCustomer) Pay(money float64) error {
	return c.wallet.Deduct(money)
}

func (c *PremiumCustomer) HasPremium() bool {
	return false
}

func (c *PremiumCustomer) HasDiscountFor(product Product) bool {
	for _, p := range c.policies {
		if p.IsApplicableFor(c, product) {
			return true
		}
	}

	return false
}

type AllUser interface {
	AddToShoppingCart(product Product)
	IsLoggedIn() bool
	Pay(money float64) error
	HasPremium() bool
	HasDiscountFor(product Product) bool
}

// interface yang cocok untuk guest
// namun tetap dipakai oleh normalCustomer dan premiumCustomer
type IUser interface {
	AddToShoppingCart(product Product)
}

// interface yang cocok untuk normalCustomer
// namun tetap dipakai oleh premiumCustomer
type ILoggedInUser interface {
	IUser
	Pay(money float64) error
}

// interface yang cocok untuk normalCustomer dan premiumCustomer
type IPremiumUser interface {
	ILoggedInUser
	HasDiscountFor(product Product) bool
}

// func NewGuest() IUser {
// 	return &Guest{}
// }

// func NewNormalUser() ILoggedInUser {
// 	return &NormalCustomer{}
// }

// func NewPremiumUser() IPremiumUser {
// 	return &PremiumCustomer{}
// }

func NewGuest() AllUser {
	return &Guest{}
}

func NewNormalUser() AllUser {
	return &NormalCustomer{}
}

func NewPremiumUser() AllUser {
	return &PremiumCustomer{}
}
