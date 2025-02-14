package handlers

import (
	"time"
	"fmt"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"price-compare-v3/configs/db"
	"price-compare-v3/models"
)

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


func GetLastPricesFromStores(ProductID uuid.UUID, c echo.Context) error {
	dbConnection := configs_db.NewConnection()
	prices, err := dbConnection.Get(c.Request().Context(), ProductID.String()).Result()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, prices)
}