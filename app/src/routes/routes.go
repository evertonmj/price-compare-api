package routes

import (
    "github.com/labstack/echo/v4"

	"price-compare-v3/api/handlers"
)

func DefineRoutes(e *echo.Echo) {
    e.POST("/products", handlers.AddProduct)
    e.PUT("/products/", handlers.UpdateProductPrices)
    e.GET("/products/:id", handlers.GetProduct)
    e.PUT("/products/:id", handlers.UpdateProductByID)
    e.DELETE("/products/:id", handlers.DeleteProductById)
    e.GET("/healthCheck", handlers.HealthCheck)
}