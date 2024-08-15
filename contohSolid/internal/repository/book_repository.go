package repository

import (
	"basic-solid/internal/domain"
	"errors"
	"fmt"
)

// ini interface untuk menampung semua interface yang ingin digunakan 0_0
type BookRepositoryInterface interface {
	BookSaver
	BookUpdater
	BookDeleter
}

// ini interface untuk diimplementasi di repo untuk menyimpan buku
type BookSaver interface {
	SaveBook(book *domain.Book) error
}

// ini interface untuk diimplementasi di repo untuk update buku
type BookUpdater interface {
	UpdateBook(book *domain.Book) error
}

// ini interface untuk diimplementasi di repo untuk delete buku
type BookDeleter interface {
	DeleteBook(bookID int) error
}

// ini ceritanya db untuk Book
type BookRepository struct {
	books map[int]domain.Book
}

// ini untuk koneksi ke db nya untuk ngubah/nampil data ^___^(❁´◡`❁)
func NewBookRepository() BookRepositoryInterface {
	return BookRepository{
		books: map[int]domain.Book{},
	}
}

// ini fungsi untuk menyimpan buku pada DB T_T
func (repo BookRepository) SaveBook(book *domain.Book) error {
	// mengecek jika id sudah ada pada db atau tidak
	// jika ada maka akan return error
	if _, exists := repo.books[book.ID]; exists {
		return errors.New("book already exists")
	}

	// jika tidak ada error maka akan menyimpan data pada db dan return nil(no error)
	repo.books[book.ID] = *book
	fmt.Println(repo.books)
	return nil
}

// ini fungsi untuk menghapus buku pada DB ☆*: .｡. o(≧▽≦)o .｡.:*☆
func (repo BookRepository) DeleteBook(bookID int) error {
	// mengecek jika id ada atau tidak pada DB
	// jika tidak ada maka akan return error
	if _, exists := repo.books[bookID]; !exists {
		return errors.New("book not found")
	}

	// jika ada maka akan delete data pada DB dan return nil(no error)
	delete(repo.books, bookID)
	return nil
}

// ini fungsi untuk update buku pada DB
func (repo BookRepository) UpdateBook(book *domain.Book) error {
	// mengecek jika id ada atau tidak pada DB
	// jika ada maka akan update buku sesuai masukan pada DB
	if _, exists := repo.books[book.ID]; exists {
		// book.ID = foundBook.ID
		repo.books[book.ID] = *book
		return nil
	}

	// kalau tidak ada maka akan return error
	return errors.New("book not found")
}

// func (repo *BookRepository) GetAll() []*domain.Book {
// 	list := make([]*domain.Book, 0)
// 	for _, book := range repo.books {
// 		list = append(list, &book)
// 	}
// 	return list
// }

// func (repo *BookRepository) GetByID(bookID int) (*domain.Book, error) {
// 	book, exists := repo.books[bookID]
// 	if !exists {
// 		return nil, errors.New("book not found")
// 	}

// 	return &book, nil
// }
