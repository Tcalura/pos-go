package main

import (
	"net/http"

	"mod-apis/configs"
	"mod-apis/internal/entity"
	"mod-apis/internal/infra/database"
	"mod-apis/internal/infra/webserver/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//...
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userHandler := handlers.NewUserHandler(database.NewUser(db))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.CreateUser)
	// r.GET("/users", userHandler.GetUsers)
	// r.GET("/users/{id}", userHandler.GetUser)
	// r.PUT("/users/{id}", userHandler.UpdateUser)
	// r.DELETE("/users/{id}", userHandler.DeleteUser)

	http.ListenAndServe(":8000", r)
}
