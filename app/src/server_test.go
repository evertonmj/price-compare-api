package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"log"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"price-compare-v3/api/handlers"
)
func TestAddProduct(t *testing.T) {
	e := echo.New()
	productJSON := `{"name":"Test Product","description":"Test Description","current_price":10.0}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.AddProduct(c), "AddProduct handler returned an error") {
		assert.Equal(t, http.StatusCreated, rec.Code, "Expected status code 201, got %v", rec.Code)
		log.Println("Response:", rec.Body.String())
	}
}

func TestUpdateProductPrices(t *testing.T) {
	e := echo.New()
	productJSON := `{"product_id":"` + uuid.New().String() + `","name":"Test Product","description":"Test Description","current_price":15.0}`
	req := httptest.NewRequest(http.MethodPut, "/products", strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.UpdateProductPrices(c), "UpdateProductPrices handler returned an error") {
		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200, got %v", rec.Code)
		log.Println("Response:", rec.Body.String())
	}
}

func TestUpdateProductByID(t *testing.T) {
	e := echo.New()
	productJSON := `{"product_id":"` + uuid.New().String() + `","name":"Test Product","description":"Test Description","current_price":20.0}`
	req := httptest.NewRequest(http.MethodPut, "/products/"+uuid.New().String(), strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(uuid.New().String())

	if assert.NoError(t, handlers.UpdateProductByID(c), "UpdateProductByID handler returned an error") {
		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200, got %v", rec.Code)
		log.Println("Response:", rec.Body.String())
	}
}

func TestDeleteProductById(t *testing.T) {
	e := echo.New()
	productID := uuid.New().String()
	req := httptest.NewRequest(http.MethodDelete, "/products/"+productID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(productID)

	if assert.NoError(t, handlers.DeleteProductById(c), "DeleteProductById handler returned an error") {
		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200, got %v", rec.Code)
		log.Println("Response:", rec.Body.String())
	}
}
