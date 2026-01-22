package main

import (
	"encoding/json"
	"fmt"
	"go-kasir-api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"hello": "world",
		})
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetAllCategories(w, r)
		case "POST":
			handlers.StoreCategory(w, r)
		}
	})

	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetCategoryById(w, r)
		case "PUT":
			handlers.UpdateCategory(w, r)
		case "DELETE":
			handlers.DeleteCategory(w, r)
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API is Running",
		})
	})

	fmt.Println("Server running on localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to running server")
	}
}
