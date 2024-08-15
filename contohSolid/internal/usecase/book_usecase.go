package usecase

import (
	"basic-solid/internal/domain"
	"basic-solid/internal/repository"
)

// type BookUsecase struct {
// 	repository repository.BookRepository
// }

type BookUsecaseInterface interface {
	AddBook
	DeleteBook
	UpdateBook
}

type AddBook interface {
	AddBook(book domain.Book) error
}
type DeleteBook interface {
	DeleteBook(book domain.Book) error
}
type UpdateBook interface {
	UpdateBook(book domain.Book) error
}

type BookUsecaseImpl struct {
	bookRepo repository.BookRepositoryInterface
}

func NewBookUsecase(bookRepo repository.BookRepositoryInterface) BookUsecaseInterface {
	return BookUsecaseImpl{
		bookRepo: bookRepo,
	}
}

func (uc BookUsecaseImpl) AddBook(book domain.Book) error {
	err := uc.bookRepo.SaveBook(&book)
	if err != nil {
		return err
	}
	return nil
}
func (uc BookUsecaseImpl) UpdateBook(book domain.Book) error {
	err := uc.bookRepo.UpdateBook(&book)
	if err != nil {
		return err
	}
	return nil
}
func (uc BookUsecaseImpl) DeleteBook(book domain.Book) error {
	err := uc.bookRepo.DeleteBook(book.ID)
	if err != nil {
		return err
	}
	return nil
}
