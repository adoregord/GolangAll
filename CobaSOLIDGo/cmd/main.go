package main

import (
	"fmt"

	"main.go/internal/repository"
	"main.go/internal/usecase"
)

func main() {
	repo := repository.NewBookRepository()
	uc := usecase.NewBookUsecase(repo)

	fmt.Println(uc)
}		
