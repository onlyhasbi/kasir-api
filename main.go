package main

import (
	"cashier/configs"
	_ "cashier/docs"
	"cashier/handlers"
	"cashier/repositories"
	"cashier/services"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Cashier API
// @version         1.0
// @host            kasir-api-production-7789.up.railway.app
// @BasePath        /

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

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/category", categoryHandler.GetAll)
	mux.HandleFunc("POST /api/category", categoryHandler.PostCategory)
	mux.HandleFunc("GET /api/category/{id}", categoryHandler.GetCategory)
	mux.HandleFunc("DELETE /api/category/{id}", categoryHandler.DeleteCategory)
	mux.HandleFunc("PUT /api/category/{id}", categoryHandler.UpdateCategory)

	productRepositories := repositories.NewProductsRepository(db)
	productService := services.NewProductsService(productRepositories)
	productHandler := handlers.NewProductsHandler(productService)

	mux.HandleFunc("GET /api/product", productHandler.GetAll)
	mux.HandleFunc("POST /api/product", productHandler.PostProduct)
	mux.HandleFunc("GET /api/product/{id}", productHandler.GetProduct)
	mux.HandleFunc("DELETE /api/product/{id}", productHandler.DeleteProduct)
	mux.HandleFunc("PUT /api/product/{id}", productHandler.UpdateProduct)

	mux.Handle("/", httpSwagger.WrapHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	serverAddr := "0.0.0.0:" + cfg.Port
	fmt.Printf("Server running in localhost%s\n", serverAddr)

	err = http.ListenAndServe(serverAddr, handler)
	if err != nil {
		log.Fatalf("Failed to run server ...")
	}
}
