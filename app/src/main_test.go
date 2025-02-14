package main

import (
	"os"
	"price-compare-v3/routes"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Setup code if needed

	code := m.Run()

	// Teardown code if needed

	os.Exit(code)
}

func TestPortEnvVariable(t *testing.T) {
	os.Setenv("PORT", "8080")
	defer os.Unsetenv("PORT")

	port := os.Getenv("PORT")
	assert.Equal(t, "8080", port)
}

func TestDefaultPort(t *testing.T) {
	os.Unsetenv("PORT")

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	assert.Equal(t, "1323", port)
}

func TestRoutesDefined(t *testing.T) {
	e := echo.New()
	routes.DefineRoutes(e)

	// Assuming you have some routes defined, you can check if they exist
	assert.NotEmpty(t, e.Routes())
}
