// ini untuk repository (koneksi ke DB)
package repository

import (
	"errors"

	"main.go/internal/domain"
)

// ini kumpulan interface//
type BookInterface interface {
	BookAdd
	BookDelete
	BookShowAll
}
type BookAdd interface {
	Add(book *domain.Book) error
}
type BookDelete interface {
	Delete(bookId int) error
}
type BookShowAll interface {
	ShowAll() error
}

// buat database untuk menyimpan data book
type BookRepository struct {
	Books map[int]domain.Book
}

// memaksa book untuk mengimplementasi interface
func NewBookRepository() BookInterface {
	return BookRepository{
		Books: map[int]domain.Book{},
	}
}

func (repo BookRepository) Add(book *domain.Book) error {
	_, exist := repo.Books[book.ID]
	if exist {
		return errors.New("BOOK ALREADY EXISTS")
	}
	return nil
}
func (repo BookRepository) Delete(bookId int) error {
	_, exist := repo.Books[bookId]
	if !exist {
		return errors.New("THERE'S NO BOOK WITH SUCH ID")
	}
	return nil
}
func (repo BookRepository) ShowAll() error {
	if len(repo.Books) == 0 {
		return errors.New("THERE'S NO BOOK TO DISPLAY")
	}
	return nil
}
