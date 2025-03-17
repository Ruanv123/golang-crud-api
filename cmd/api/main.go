package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ruanv123/api-go-crud/internal/database"
	"github.com/ruanv123/api-go-crud/internal/handler"
	"github.com/ruanv123/api-go-crud/internal/middleware"
	"github.com/ruanv123/api-go-crud/internal/repository"
	"github.com/ruanv123/api-go-crud/internal/service"
)

func main() {
	db := database.ConnectDB()

	app := fiber.New()
	// middleware de logs no  console
	app.Use(logger.New())
	// setando o middleware de cors
	app.Use(cors.New())

	app.Use(middleware.ErrorHandler())

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	api := app.Group("/api")
	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)
	api.Get("/profile", middleware.JWTProtected(), userHandler.Profile)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	api.Post("/products", middleware.JWTProtected(), productHandler.CreateProduct)
	api.Get("/products", middleware.JWTProtected(), productHandler.GetAllProducts)
	api.Get("/products/:id", middleware.JWTProtected(), productHandler.GetProductByID)
	api.Put("/products/:id", middleware.JWTProtected(), productHandler.UpdateProduct)
	api.Delete("/products/:id", middleware.JWTProtected(), productHandler.DeleteProduct)

	log.Fatal(app.Listen(":8000"))
}
