package domain

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title" validate:"noblank"`
	Author string `json:"author" validate:"noblank"`
	Stock  int    `json:"stock" validate:"required"`
}
