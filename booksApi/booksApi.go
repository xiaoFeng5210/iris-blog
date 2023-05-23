package booksapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	e.GET("/", helloWorld)
}

func helloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
