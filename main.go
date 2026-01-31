package main

import (
	"cashier/configs"
	"cashier/handlers"
	"cashier/repositories"
	"cashier/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg, err := configs.LoadingConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := configs.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer db.Close()

	categoryRepositories := repositories.NewCategoriesRepository(db)
	categoryService := services.NewCategoriesService(categoryRepositories)
	categoryHandler := handlers.NewCategoriesHandler(categoryService)

	http.HandleFunc("GET /api/category", categoryHandler.GetAll)
	http.HandleFunc("POST /api/category", categoryHandler.PostCategory)
	http.HandleFunc("GET /api/category/{id}", categoryHandler.GetCategory)
	http.HandleFunc("DELETE /api/category/{id}", categoryHandler.DeleteCategory)
	http.HandleFunc("PUT /api/category/{id}", categoryHandler.UpdateCategory)

	productRepositories := repositories.NewProductsRepository(db)
	productService := services.NewProductsService(productRepositories)
	productHandler := handlers.NewProductsHandler(productService)

	http.HandleFunc("GET /api/product", productHandler.GetAll)
	http.HandleFunc("POST /api/product", productHandler.PostProduct)
	http.HandleFunc("GET /api/product/{id}", productHandler.GetProduct)
	http.HandleFunc("DELETE /api/product/{id}", productHandler.DeleteProduct)
	http.HandleFunc("PUT /api/product/{id}", productHandler.UpdateProduct)

	serverAddr := "0.0.0.0:" + cfg.Port
	fmt.Printf("Server running in localhost%s\n", serverAddr)

	err = http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalf("Failed to run server ...")
	}
}
