package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", Hello)
	e.Logger.Fatal(e.Start(":8082"))
}

func Hello(ctxt echo.Context) error {
	return ctxt.String(http.StatusOK, "hi")
}
