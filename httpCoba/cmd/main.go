package main

import (
	"fmt"
	"http/internal/domain"
	http "http/package"
)

func main() {
	router := http.NewRouter()

	router.HandleFunc("GET", "/Didi", func(w http.ResponseWriter, r *http.Request) {
		w.WriteJSON(domain.User{
			Name:  "Didi",
			Email: "Didi@gmail.com",
			Age:   21,
		})
	})

	router.HandleFunc("POST", "/Didi/post", func(w http.ResponseWriter, r *http.Request) {
		var dodo domain.User

		if err := r.ParseJSON(&dodo); err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Error parsing JSON"))
			return
		}

		w.WriteJSON(dodo)
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// server starting
	fmt.Println("Server is running on port: 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
