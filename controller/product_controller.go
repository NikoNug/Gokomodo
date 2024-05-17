package controller

import (
	"Gokomodo/db"
	"Gokomodo/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ViewProduct(c *fiber.Ctx) error {
	db := db.ConnectDB()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	result := models.Products{}

	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.SellerID); err != nil {
			return err
		}

		// Append to the array
		result.Products = append(result.Products, product)
	}

	return c.JSON(result)
}

func AddProduct(c *fiber.Ctx) error {
	db := db.ConnectDB()

	p := new(models.Product)
	u := new(models.Seller)

	if err := c.BodyParser(p); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// ProductID
	p.ProductID = int64(uuid.New().ID())

	// Insert to DB
	rows, err := db.Query("INSERT INTO products (ProductID, Product_Name, Description, Price, SellerID) VALUES (?,?,?,?,?)", p.ProductID, p.ProductName, p.Description, p.Price, u.SellerID)
	if err != nil {
		return err
	}
	rows.Close()
	db.Close()

	log.Println(rows)

	return c.JSON(p)
}
