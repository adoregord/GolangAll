package usecase

import "basic-solid/internal/repository"

type BookUsecase struct {
	repository repository.BookRepository
}

type BookUsecaseInterface interface {
	AddBook
}

type AddBook interface {
	AddUser()
}

type BookUsecaseImpl struct {
	bookRepo repository.BookRepositoryInterface
}

func NewBookUsecase(bookRepo repository.BookRepositoryInterface) BookUsecaseInterface {
	return BookUsecaseImpl{
		bookRepo: bookRepo,
	}
}

func (usecase BookUsecaseImpl) AddUser()
