package controller

import (
	"Gokomodo/db"
	"Gokomodo/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterSeller(c *fiber.Ctx) error {
	db := db.ConnectDB()
	user := new(models.Seller)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// New user to database
	user.SellerID = int(uuid.New().ID())
	rows, err := db.Query("INSERT INTO sellers (SellerID, Email, Name, Password, Pickup) VALUES (?,?,?,?,?)", user.SellerID, user.Email, user.Name, hashedPassword, user.Pickup)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Close()

	log.Println(rows)

	fmt.Println("User Registered!")
	return c.JSON(user)
}

func RegisterBuyer(c *fiber.Ctx) error {
	db := db.ConnectDB()
	user := new(models.Buyer)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// New user to database
	user.BuyerID = int(uuid.New().ID())
	rows, err := db.Query("INSERT INTO buyers (BuyerID, Email, Name, Password, Alamat_Pengiriman) VALUES (?,?,?,?,?)", user.BuyerID, user.Email, user.Name, hashedPassword, user.Alamat_Pengiriman)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Close()

	log.Println(rows)

	fmt.Println("User Registered!")
	return c.JSON(user)
}
