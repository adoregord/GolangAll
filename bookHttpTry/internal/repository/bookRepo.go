package repository

import (
	"errors"
	"httpBook/internal/domain"
)

type BookRepoInterface interface {
	BookAdd
	BookDelete
	BookView
}
type BookAdd interface {
	AddBook(book *domain.Book) error
}
type BookDelete interface {
	DeleteBook(id int) error
}
type BookView interface {
	ViewBook() ([]domain.Book, error)
}

type BookRepo struct{
	Book map[int]domain.Book
}

func NewBookRepo() BookRepoInterface{
	return BookRepo{
		Book: map[int]domain.Book{},
	}
}

// implement the func
func (repo BookRepo) AddBook(book *domain.Book) error {
	
	repo.Book[book.ID] = *book
	return nil
}
func (repo BookRepo) DeleteBook(id int) error {
	_, exist := repo.Book[id]
	if !exist {
		return errors.New("BOOK ID UNAVAILABLE 🤬🤬🤬")
	}
	delete(repo.Book, id)
	return nil
}
func (repo BookRepo) ViewBook() ([]domain.Book, error) {
	books := make([]domain.Book, 0, len(repo.Book))
	for _, book := range repo.Book	 {
		books = append(books, book)
	}
	return books, nil
}