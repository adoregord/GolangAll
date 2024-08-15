package usecase

import (
	"httpBook/internal/domain"
	"httpBook/internal/repository"
)

type BookUsecaseInterface interface {
	BookAdd
	BookDelete
	BookView
}
type BookAdd interface {
	AddBook(book domain.Book) error
}
type BookDelete interface {
	DeleteBook(id int) error
}
type BookView interface {
	ViewBook() ([]domain.Book, error)
}

type BookUsecase struct {
	BookRepo repository.BookRepoInterface
}

func NewBookUsecase(bookRepo repository.BookRepoInterface) BookUsecaseInterface {
	return BookUsecase{
		BookRepo: bookRepo,
	}
}

func (uc BookUsecase) AddBook(book domain.Book) error {
	return uc.BookRepo.AddBook(&book)
}
func (uc BookUsecase) DeleteBook(id int) error {
	return uc.BookRepo.DeleteBook(id)
}
func (uc BookUsecase) ViewBook() ([]domain.Book, error) {
	return uc.BookRepo.ViewBook()
}
