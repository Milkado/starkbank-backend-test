package routes

import (
	"net/http"

	"github.com/Milkado/stark-backend-test/app"
	"github.com/labstack/echo/v5"
)

func Routes(server *echo.Echo) {

	server.GET("/test", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.POST("/webhook/payment", app.Listener)

	server.POST("/start-cron", app.StartCron)

	server.GET("/", app.DashboardHandler)
	server.GET("/data", app.DashboardDataHandler)
}
