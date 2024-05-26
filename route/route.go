package route

import (
	"Gokomodo/controller"
	"Gokomodo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {
	auth := c.Group("/auth")
	auth.Post("/registerSeller", controller.RegisterSeller)
	auth.Post("/registerBuyer", controller.RegisterBuyer)
	auth.Post("/loginBuyer", controller.LoginBuyer)
	auth.Get("/logout", controller.Logout)

	product := c.Group("/product")
	product.Use(middlewares.JWTMiddleware())
	product.Get("/products", controller.ViewProduct)
	product.Post("/addProduct", controller.AddProduct)

	order := c.Group("/order")
	order.Get("/orders", controller.ViewOrders)
}
