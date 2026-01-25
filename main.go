package main

import (
	"category/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /api/category", handlers.GetCategories)
	http.HandleFunc("POST /api/category", handlers.PostCategory)
	http.HandleFunc("GET /api/category/", handlers.GetCategory)
	http.HandleFunc("DELETE /api/category/", handlers.DeleteCategory)
	http.HandleFunc("PUT /api/category/", handlers.UpdateCategory)
	http.HandleFunc("GET /health", handlers.GetHealthCheck)

	fmt.Println("Server running di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to run server ...")
	}
}
