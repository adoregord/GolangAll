package main

import (
	"encoding/json"
	"fmt"
	"httpPakeNet/internal/domain"
	"net/http"
)

// handler untuk get
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("content-type", "application/json")
	response := map[string]string{"message": "Hello World"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handler untuk post
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("content-type", "application/json")
	var data domain.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := data
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/post1", getHandler)
	http.HandleFunc("/post", postHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
