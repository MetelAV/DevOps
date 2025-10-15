package main

import (
	"encoding/json"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response map[string]string

	switch r.Method {
	case http.MethodGet:
		response = map[string]string{"message": "Hello, World!"}
	case http.MethodPost:
		response = map[string]string{"message": "You sent a POST request!"}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response = map[string]string{"error": "Method not allowed"}
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
