package routes

import (
	"cashier/handlers"
	"net/http"

	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(catHandler *handlers.CategoriesHandler, prodHandler *handlers.ProductsHandler) http.Handler {

	// Category Routes
	http.HandleFunc("GET /api/category", catHandler.GetAll)
	http.HandleFunc("POST /api/category", catHandler.PostCategory)
	http.HandleFunc("GET /api/category/{id}", catHandler.GetCategory)
	http.HandleFunc("DELETE /api/category/{id}", catHandler.DeleteCategory)
	http.HandleFunc("PUT /api/category/{id}", catHandler.UpdateCategory)

	// Product Routes
	http.HandleFunc("GET /api/product", prodHandler.GetAll)
	http.HandleFunc("POST /api/product", prodHandler.PostProduct)
	http.HandleFunc("GET /api/product/{id}", prodHandler.GetProduct)
	http.HandleFunc("DELETE /api/product/{id}", prodHandler.DeleteProduct)
	http.HandleFunc("PUT /api/product/{id}", prodHandler.UpdateProduct)

	http.Handle("/", httpSwagger.WrapHandler)

	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(http.DefaultServeMux)
}
