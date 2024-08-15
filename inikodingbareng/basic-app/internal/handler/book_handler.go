package handler

import (
	"basic-solid/internal/domain"
	"basic-solid/internal/usecase"
)

// ini @Manualwiring untuk menyambungkan handler(controller) dan usecase(service)
type BookHandler struct {
	BookUsecase usecase.BookUsecase
}

// ini untuk komunikasinya
func NewBookHandler(bookUsecase usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		BookUsecase: bookUsecase,
	}
}

func (b BookHandler) StoreNewBook(book domain.Book) {

}

func (b BookHandler) UpdateBook(book domain.Book) {

}
