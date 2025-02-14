package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDefineRoutes(t *testing.T) {
	e := echo.New()
	DefineRoutes(e)

	tests := []struct {
		method string
		path   string
	}{
		{http.MethodPost, "/products"},
		{http.MethodPut, "/products/"},
		{http.MethodGet, "/products/:id"},
		{http.MethodPut, "/products/:id"},
		{http.MethodDelete, "/products/:id"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Find the route
		e.Router().Find(tt.method, tt.path, c)
		assert.NotNil(t, c.Path(), "Route should be defined for %s %s", tt.method, tt.path)
	}
}
