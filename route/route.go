package route

import (
	"Gokomodo/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {
	auth := c.Group("/auth")
	auth.Post("/registerSeller", controller.RegisterSeller)
	auth.Post("/registerBuyer", controller.RegisterBuyer)
}
