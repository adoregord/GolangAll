package main

import (
	"basic-solid/internal/domain"
	"basic-solid/internal/handler"
	"basic-solid/internal/repository"
	"basic-solid/internal/usecase"
	"fmt"
)

func main() {
	repo := repository.NewBookRepository()
	uc := usecase.NewBookUsecase(repo)
	h := handler.NewBookHandler(uc)

	bookReq := domain.Book{
		ID:          0,
		Title:       "Buku Ngoding Golang",
		Author:      "Bene",
		Type:        "Programming",
		Genre:       []string{"Horror"},
		Description: "Buku ini sangat menyenangkan",
		Stock:       1,
	}
	err := h.StoreNewBook(bookReq)
	if err != nil {
		fmt.Println(err.Error())
	}
}
