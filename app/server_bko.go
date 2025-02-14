package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type Product struct {
	ProductID      uuid.UUID `json:product_id`
	Name           string    `json:name`
	Description    string    `json:description`
	CurrentPrice   float64   `json:current_price`
	HistoricPrices []Price   `json:historic_prices`
	CreatedAt      time.Time `json:created_at`
	UpdatedAt      time.Time `json:updated_at`
}

func (i *Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Product) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}

type Price struct {
	ProductID uuid.UUID `json:product_id`
	StoreID   uuid.UUID `json:store_id`
	Price     float64   `json:price`
	CreatedAt time.Time `json:created_at`
}

func (i *Price) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

type Store struct {
	StoreID  uuid.UUID `json:store_id`
	Name     string    `json:name`
	Location string    `json:location`
	Products []Product `json:products`
}

func (i *Store) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func main() {
	startWebServer()
	fmt.Println("Server started!")
	getDBConnection()
	fmt.Println("Database connected!")
}

func startWebServer() {
	e := echo.New()

	defineRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func getDBConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Username: "default", // use your Redis user. More info https://redis.io/docs/latest/operate/oss_and_stack/management/security/acl/
		Password: "secret",  // use your Redis password
	},
	)
	return client
}

func defineRoutes(e *echo.Echo) {
	e.GET("/products", CompareProductPrices)
	e.POST("/products", AddProduct)
	e.PUT("/products/", UpdateProductPrices)
	e.GET("/products/:id", GetProduct)
	e.PUT("/products/:id", UpdateProductByID)
	e.DELETE("/products/:id", DeleteProductById)
}

func CompareProductPrices(c echo.Context) error {
	// This function will compare the prices of products
	dbConnection := getDBConnection()
	products, err := dbConnection.Get(c.Request().Context(), "products").Result()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, products)
}

func AddProduct(c echo.Context) error {
	fmt.Printf("Adding product")
	dbConnection := getDBConnection()

	product := new(Product)

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
	dbConnection := getDBConnection()
	product := new(Product)

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
	dbConnection := getDBConnection()
	productID, err := uuid.Parse(c.Param("product_id"))
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

func UpdateProductByID(c echo.Context) error {
	// This function will update a product by ID
	dbConnection := getDBConnection()
	product := new(Product)

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

func DeleteProductById(c echo.Context) error {
	dbConnection := getDBConnection()
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

func GetLastPricesFromStores(ProductID uuid.UUID, c echo.Context) error {
	// This function will get the last prices from stores
	dbConnection := getDBConnection()
	prices, err := dbConnection.Get(c.Request().Context(), ProductID.String()).Result()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, prices)
}
