package main

import (
	"cashier/configs"
	_ "cashier/docs"
	"cashier/handlers"
	"cashier/repositories"
	"cashier/routes"
	"cashier/services"
	"fmt"
	"log"
	"net/http"
)

// @title           Cashier API
// @version         1.0
// @host            kasir-api-production-7789.up.railway.app
// @BasePath        /
func main() {
	cfg, err := configs.LoadingConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	db, err := configs.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	defer db.Close()

	catRepo := repositories.NewCategoriesRepository(db)
	catService := services.NewCategoriesService(catRepo)
	catHandler := handlers.NewCategoriesHandler(catService)

	prodRepo := repositories.NewProductsRepository(db)
	prodService := services.NewProductsService(prodRepo)
	prodHandler := handlers.NewProductsHandler(prodService)

	handler := routes.NewRouter(catHandler, prodHandler)

	serverAddr := "0.0.0.0:" + cfg.Port
	fmt.Printf("Server running at %s\n", serverAddr)

	if err := http.ListenAndServe(serverAddr, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
