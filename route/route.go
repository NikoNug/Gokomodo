package route

import (
	"Gokomodo/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {
	auth := c.Group("/auth")
	auth.Post("/registerSeller", controller.RegisterSeller)
	auth.Post("/registerBuyer", controller.RegisterBuyer)

	product := c.Group("/product")
	product.Get("/products", controller.ViewProduct)
	product.Post("/addProduct", controller.AddProduct)

	order := c.Group("/order")
	order.Get("/orders", controller.ViewOrders)
}
