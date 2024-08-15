package domain

import "sync"

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
	MUTEK   *sync.Mutex
}