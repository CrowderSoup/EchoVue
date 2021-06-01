package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Static Files
	e.Static("/css", "./app/dist/css")
	e.Static("/js", "./app/dist/js")
	e.Static("/img", "./app/dist/img")

	// Front-end Entrypoint
	e.File("/", "./app/dist/index.html")

	e.Logger.Fatal(e.Start(":1313"))
}
