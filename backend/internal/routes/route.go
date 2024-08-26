package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", handlers.Login)

	app.Get("/products", middleware.AuthMiddleware(), handlers.GetAllProducts)
	app.Post("/products", middleware.AuthMiddleware(), handlers.CreateProduct)
	app.Get("/products/:id", middleware.AuthMiddleware(), handlers.GetProduct)
	app.Put("/products/:id", middleware.AuthMiddleware(), handlers.UpdateProduct)
	app.Delete("/products/:id", middleware.AuthMiddleware(), handlers.DeleteProduct)

	// Rotas para Categorias
	app.Get("/categories", handlers.GetAllCategories)
	app.Post("/categories", handlers.CreateCategory)
	app.Get("/categories/:id", handlers.GetCategory)
	app.Put("/categories/:id", handlers.UpdateCategory)
	app.Delete("/categories/:id", handlers.DeleteCategory)
}
