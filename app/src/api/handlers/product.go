package handlers

import (
	"fmt"
	"net/http"
	configs_db "price-compare-v3/configs/db"
	"price-compare-v3/models"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// HealthCheck returns a simple message indicating the service is alive
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "I'm alive!")
}

// Readiness checks if the database connection is ready
func Readiness(c echo.Context) error {
	fmt.Printf("Testing database readiness")
	dbConnection := configs_db.NewConnection()
	err := dbConnection.Ping(c.Request().Context()).Err()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Database running and I'm ready!")
}

// AddProduct adds a new product to the database
func AddProduct(c echo.Context) error {
	fmt.Printf("Adding product")
	dbConnection := configs_db.NewConnection()

	product := new(models.Product)

	if err := c.Bind(&product); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	product.ProductID = uuid.New()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	err := dbConnection.Set(c.Request().Context(), "products:"+product.ProductID.String(), product, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, product)
}

// UpdateProductPrices updates the prices of an existing product
func UpdateProductPrices(c echo.Context) error {
	dbConnection := configs_db.NewConnection()

	product := new(models.Product)

	if err := c.Bind(&product); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	product.UpdatedAt = time.Now()
	err := dbConnection.Set(c.Request().Context(), "products:"+product.ProductID.String(), product, 0).Err()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, product)
}

// GetProduct retrieves a product by its ID
func GetProduct(c echo.Context) error {
	dbConnection := configs_db.NewConnection()

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	product, err := dbConnection.Get(c.Request().Context(), "products:"+productID.String()).Result()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, product)
}

// GetAllProducts retrieves all products from the database
func GetAllProducts(c echo.Context) error {
	dbConnection := configs_db.NewConnection()

	products, err := dbConnection.Get(c.Request().Context(), "products:").Result()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, products)
}

// DeleteProductById deletes a product by its ID
func DeleteProductById(c echo.Context) error {
	dbConnection := configs_db.NewConnection()

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	err = dbConnection.Del(c.Request().Context(), productID.String()).Err()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Product deleted successfully")
}

// UpdateProductByID updates a product by its ID
func UpdateProductByID(c echo.Context) error {
	dbConnection := configs_db.NewConnection()

	product := new(models.Product)

	if err := c.Bind(&product); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	product.UpdatedAt = time.Now()
	err := dbConnection.Set(c.Request().Context(), product.ProductID.String(), product, 0).Err()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, product)
}

// GetLastPricesFromStores retrieves the last prices from stores for a given product ID
func GetLastPricesFromStores(ProductID uuid.UUID, c echo.Context) error {
	dbConnection := configs_db.NewConnection()
	prices, err := dbConnection.Get(c.Request().Context(), ProductID.String()).Result()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, prices)
}
