package main

import (
	"cashier/configs"
	"cashier/handlers"
	"cashier/repositories"
	"cashier/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	cfg, err := configs.LoadingConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	config := configs.Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := configs.InitDB(config.DBConn)
	if err != nil {
		fmt.Printf("Failed to initialize database")
	}

	productRepositories := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepositories)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("GET /api/category", handlers.GetCategories)
	http.HandleFunc("POST /api/category", handlers.PostCategory)
	http.HandleFunc("GET /api/category/{id}", handlers.GetCategory)
	http.HandleFunc("DELETE /api/category/{id}", handlers.DeleteCategory)
	http.HandleFunc("PUT /api/category/{id}", handlers.UpdateCategory)

	http.HandleFunc("GET /api/product", productHandler.GetAll)
	http.HandleFunc("POST /api/product", productHandler.PostProduct)
	http.HandleFunc("GET /api/product/{id}", productHandler.GetProduct)
	http.HandleFunc("DELETE /api/product/{id}", productHandler.DeleteProduct)
	http.HandleFunc("PUT /api/product/{id}", productHandler.UpdateProduct)

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	defer db.Close()

	serverAddr := "0.0.0.0:" + cfg.Port
	fmt.Printf("Server running in localhost%s\n", serverAddr)

	err = http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalf("Failed to run server ...")
	}
}
