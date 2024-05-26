package controller

import (
	"Gokomodo/config"
	"Gokomodo/db"
	"Gokomodo/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

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

func LoginBuyer(c *fiber.Ctx) error {
	db := db.ConnectDB()
	var userInput models.Buyer

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	var user models.Buyer
	// Check DB
	query := "SELECT Email, Password FROM buyers WHERE Email = ?"
	row := db.QueryRow(query, userInput.Email)
	err := row.Scan(&user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendString("Email/Pass is Incorrect")
		}
		fmt.Printf("Error scanning row : %v\n", err)
		return c.Status(500).SendString(err.Error())
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).SendString(err.Error())
	}

	// Create JWT Token
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gokomodo",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// declare algorithm for signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HTTPOnly: true,
	})

	return c.Status(200).SendString("Login Success")
}

func Logout(c *fiber.Ctx) error {
	// delete token di cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	})

	return c.Status(200).SendString("Logout Success")
}
