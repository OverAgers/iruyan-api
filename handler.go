package main

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Home Page!"))
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []string{"Alice", "Bob", "Charlie"}

	// JSON形式でレスポンス
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
