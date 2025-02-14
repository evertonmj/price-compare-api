package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(`{"name":"Test Product"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, AddProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUpdateProductPrices(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/products", strings.NewReader(`{"productID":"123e4567-e89b-12d3-a456-426614174000","price":100}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, UpdateProductPrices(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteProductById(t *testing.T) {
	e := echo.New()
	productID := uuid.New()
	req := httptest.NewRequest(http.MethodDelete, "/products/"+productID.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(productID.String())

	if assert.NoError(t, DeleteProductById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateProductByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/products", strings.NewReader(`{"productID":"123e4567-e89b-12d3-a456-426614174000","name":"Updated Product"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, UpdateProductByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}