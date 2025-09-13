package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func main() {
	repo := NewUserRepo()
	service := NewUserService(repo)

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// Example: always query for ID=3 (non-existent)
		user, err := service.GetUserProfile(3)
		if err != nil {
			// Log with context for developers
			log.Printf("handler error: %v", err)

			// Send safe message to client
			if errors.Is(err, ErrUserNotFound) {
				http.Error(w, "user not found", http.StatusNotFound)
			} else {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			return
		}

		// On success return JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
