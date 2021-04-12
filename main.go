package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	// Routes
	echo.GET("/", hello)

	// Start server
	echo.Logger.Fatal(echo.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World!")
	jsonMap := map[string]string{
		"hoge": "this is your domain",
	}
	// hoge := map[string]int{
	// 	"hoge": 123,
	// }

	return c.JSON(http.StatusOK, jsonMap)
}
