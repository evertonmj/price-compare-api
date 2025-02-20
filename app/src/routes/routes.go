package routes

import (
    "github.com/labstack/echo/v4"

	"price-compare-v3/api/handlers"
)

func DefineRoutes(e *echo.Echo) {
    e.POST("/products", handlers.AddProduct)
    e.GET("/products/", handlers.GetAllProducts)
    e.GET("/products/:id", handlers.GetProduct)
    e.PUT("/products/:id", handlers.UpdateProductByID)
    e.DELETE("/products/:id", handlers.DeleteProductById)

    e.GET("/liveness", handlers.HealthCheck)
    e.GET("/readiness", handlers.Readiness)
}