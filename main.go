package main

import (
	"github.com/Milkado/stark-backend-test/routes"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	e.Use()

	routes.Routes(e)

	if err := e.Start(":1313"); err != nil {
		e.Logger.Error("server start fail")
	}
}
