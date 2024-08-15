package main

import (
	"fmt"
	"httpBook/internal/domain"
	"httpBook/internal/handler"
	"httpBook/internal/repository"
	"httpBook/internal/usecase"
	"net/http"
)

func main() {

	bookRepo := repository.NewBookRepo()
	BookUc := usecase.NewBookUsecase(bookRepo)
	BookH := handler.NewBookHandler(BookUc)

	books := []domain.Book{
		{ID: 1, Title: "Buku 1", Author: "", Stock: 10},
		{ID: 2, Title: "Buku 2", Author: "Penulis 2", Stock: 10},
		{ID: 3, Title: "Buku 3", Author: "Penulis 3", Stock: 10},
		{ID: 4, Title: "Buku 4", Author: "Penulis 4", Stock: 10},
	}
	for _, value := range books {
		BookUc.AddBook(value)
	}

	routes := http.NewServeMux()
	routes.HandleFunc("/books", BookH.AddBook)
	routes.HandleFunc("/booksGet", BookH.ViewBook)
	routes.HandleFunc("/booksDelete", BookH.DeleteBook)

	server := http.Server{}
	server.Handler = routes
	server.Addr = ":8080"

	fmt.Println("Server berjalan di http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
