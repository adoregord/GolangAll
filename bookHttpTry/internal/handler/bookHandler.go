package handler

import (
	"encoding/json"
	"httpBook/internal/domain"
	"httpBook/internal/usecase"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

type BookHandlerInterface interface {
	BookAdd
	BookDelete
	BookView
}
type BookAdd interface {
	AddBook(w http.ResponseWriter, r *http.Request)
}
type BookDelete interface {
	DeleteBook(w http.ResponseWriter, r *http.Request)
}
type BookView interface {
	ViewBook(w http.ResponseWriter, r *http.Request)
}

type BookHandler struct {
	BookUsecase usecase.BookUsecaseInterface
}

type response struct {
	Msg string
}


func NewBookHandler(bookUsecase usecase.BookUsecaseInterface) BookHandlerInterface {
	return BookHandler{
		BookUsecase: bookUsecase,
	}
}

func (h BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Info().
			Int("httpStatusCode", http.StatusMethodNotAllowed).
			Str("httpMethod", r.Method).
			Msg("Method not allowed")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var book domain.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Info().
			Int("httpStatusCode", http.StatusBadRequest).
			Str("httpMethod", r.Method).
			Msg("Method not allowed")
		return
	}
	// validate the input
	if err := validate.Struct(book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Info().
			Int("httpStatusCode", http.StatusBadRequest).
			Str("httpMethod", r.Method).
			Msg("Method not allowed")
		return
	}
	if err := h.BookUsecase.AddBook(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Info().Msg("Internal server error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Info().
			Int("httpStatusCode", http.StatusMethodNotAllowed).
			Str("httpMethod", "DELETE").
			Msg("Method not allowed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	bookIdStr := r.URL.Query().Get("id")
	if bookIdStr == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		log.Info().
			Int("httpStatusCode", http.StatusBadRequest).
			Str("httpMethod", "DELETE").
			Msg("Missing book ID")
		return
	}
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Info().
			Int("httpStatusCode", http.StatusBadRequest).
			Str("httpMethod", "DELETE").
			Msg("Invalid book ID")
		return
	} else if err := h.BookUsecase.DeleteBook(bookId); err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{Msg: err.Error()})
		log.Info().
			Int("httpStatusCode", http.StatusInternalServerError).
			Str("httpMethod", "DELETE").
			Msg("Internal server error")
		return
	}
	w.WriteHeader(http.StatusOK)
}
// vars := mux.Vars(r)
// idStr := vars["id"]
// id, err := strconv.Atoi(idStr)

// if err != nil {
// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	log.Info().
// 		Int("httpStatusCode", http.StatusBadRequest).
// 		Str("httpMethod", "DELETE").
// 		Msg("Invalid ID")
// 	return
// } else if err := h.BookUsecase.DeleteBook(id); err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	log.Info().Msg("Internal server error")
// 	return
// }

func (h BookHandler) ViewBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Info().
			Int("httpStatusCode", http.StatusMethodNotAllowed).
			Str("httpMethod", r.Method).
			Msg("Method not allowed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	books, err := h.BookUsecase.ViewBook()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Info().Msg("Internal server error")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
