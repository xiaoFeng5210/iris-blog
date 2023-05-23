package booksapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Title string `json:"title"`
}

func Run() {
	e := echo.New()
	e.GET("/", helloWorld)
	e.GET("/books", list)
	e.Start(":1323")
}

func list(ctx echo.Context) error {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}
	return ctx.JSON(http.StatusOK, books)
}

func helloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
