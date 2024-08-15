package model

import "time"

// Pokemon struct to hold the information about a Pokemon
type StarRail struct {
	ID             int
	Name           string
	Type           string
	CatchRate      float32
	IsRare         bool
	RegisteredDate time.Time
}
