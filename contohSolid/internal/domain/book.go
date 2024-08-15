package domain

// ini domain(model) untuk Book ---->ilove IPUL
type Book struct {
	ID          int
	Title       string
	Author      string
	Type        string // Novel/Comic/Encyclopedia
	Genre       []string
	Description string
	Stock       int
}
