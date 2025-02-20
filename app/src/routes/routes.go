package routes

import (
	"price-compare-v3/api/handlers"

	"github.com/labstack/echo/v4"
)

// DefineRoutes defines all the routes for the application
func DefineRoutes(e *echo.Echo) {
	// Product routes
	e.POST("/products", handlers.AddProduct)              // Add a new product
	e.GET("/products/", handlers.GetAllProducts)          // Get all products
	e.GET("/products/:id", handlers.GetProduct)           // Get a product by ID
	e.PUT("/products/:id", handlers.UpdateProductByID)    // Update a product by ID
	e.DELETE("/products/:id", handlers.DeleteProductById) // Delete a product by ID

	// Health check routes
	e.GET("/liveness", handlers.HealthCheck) // Check if the service is alive
	e.GET("/readiness", handlers.Readiness)  // Check if the service is ready
}
