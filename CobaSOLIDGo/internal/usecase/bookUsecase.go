// ini untuk servicenya (logic)
package usecase

import (
	"main.go/internal/domain"
	"main.go/internal/repository"
)

type BookUsecase struct{
	bookRepo repository.BookInterface
}

func NewBookUsecase(bookRepo repository.BookInterface) BookInt{
	return BookUsecase{
		bookRepo: bookRepo,
	}
}

type BookInt interface {
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

func (uc BookUsecase) Add(book *domain.Book) error{
	return nil
}
func (uc BookUsecase) Delete(bookId int) error{
	return nil
}
func (uc BookUsecase) ShowAll() error{
	return nil
}
