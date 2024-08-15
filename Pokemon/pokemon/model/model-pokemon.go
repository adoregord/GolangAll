package model

import "time"

// Pokemon struct to hold the information about a Pokemon
type Pokemon struct {
	ID             int
	Name           string
	Type           string
	CatchRate      float32
	IsRare         bool
	RegisteredDate time.Time
}
