package main

import (
    "log"
    "os"

    "github.com/labstack/echo/v4"
    "price-compare-v3/routes"
)

func main() {
    e := echo.New()

    routes.DefineRoutes(e)

    port := os.Getenv("PORT")
    if port == "" {
        port = "1323"
    }

    log.Fatal(e.Start(":" + port))
}