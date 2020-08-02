package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "/site/build/index.html")

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "site/build",
		HTML5: true,
	}))

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, os.Getenv("testvalue"))
	})

	e.Logger.Fatal(e.Start(":5000"))
}
