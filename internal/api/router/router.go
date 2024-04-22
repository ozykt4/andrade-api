package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Router func(route fiber.Router)

func SetupRouter(app *fiber.App, routes ...Router) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metricas do serviço"}))
	api := app.Group("api")

	v1 := api.Group("v1")
	for _, route := range routes {
		route(v1)
	}
}
