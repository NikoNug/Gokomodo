package controller

import (
	"Gokomodo/db"
	"Gokomodo/models"

	"github.com/gofiber/fiber/v2"
)

func ViewOrders(c *fiber.Ctx) error {
	db := db.ConnectDB()

	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return err
	}

	defer rows.Close()
	defer db.Close()

	result := models.Orders{}

	for rows.Next() {
		order := models.Order{}
		if err := rows.Scan(&order.OrderID, &order.BuyerID, &order.SellerID, &order.DeliveryDestinationAddress, &order.DeliverySourceAddress, &order.Status); err != nil {
			return err
		}

		result.Orders = append(result.Orders, order)
	}

	return c.JSON(result)
}
