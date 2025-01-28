package main

import (
	"ecommerce-microservices/config"
	"ecommerce-microservices/internal/database"
	"ecommerce-microservices/internal/handler"
	"ecommerce-microservices/internal/repository"
	"ecommerce-microservices/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadEnv()

	database.Init()

	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	r := chi.NewRouter()
	r.Post("/products", productHandler.Create)
	r.Get("/products", productHandler.GetAll)
	r.Get("/products/{id}", productHandler.GetById)

	port := config.GetEnv("PORT")
	log.Printf("Server starter at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
