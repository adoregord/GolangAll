package handler

import (
	"basic-solid/internal/domain"
	"basic-solid/internal/usecase"
)

type BookHandlerInterface interface {
	StoreNewBook(book domain.Book) error
	UpdateBook(book domain.Book) error
	DeleteBook(book domain.Book) error
}

// ini @Manualwiring untuk menyambungkan handler(controller) dan usecase(service)
type BookHandler struct {
	BookUsecase usecase.BookUsecaseInterface
}

// ini untuk komunikasinya
func NewBookHandler(bookUsecase usecase.BookUsecaseInterface) BookHandlerInterface {
	return BookHandler{
		BookUsecase: bookUsecase,
	}
}

func (h BookHandler) StoreNewBook(book domain.Book) error {
	
	err := h.BookUsecase.AddBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (h BookHandler) UpdateBook(book domain.Book) error {
	err := h.BookUsecase.UpdateBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (h BookHandler) DeleteBook(book domain.Book) error {
	err := h.BookUsecase.DeleteBook(book)
	if err != nil {
		return err
	}
	return nil
}
