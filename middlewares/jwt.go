package middlewares

import (
	"Gokomodo/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// func JWTMiddleware(c *fiber.Ctx) error {
// 	cookie := c.Cookies("token")
// 	if cookie == "" {
// 		return c.Status(http.StatusNotFound).SendString("Cookies not found")
// 	}

// 	tokenString := cookie

// 	claims := config.JWTClaim{}

// 	// parse the token
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
// 		return config.JWT_KEY, nil
// 	})

// 	if err != nil {
// 		v, _ := err.(*jwt.ValidationError)
// 		switch v.Errors {
// 		case jwt.ValidationErrorSignatureInvalid:
// 			return c.Status(http.StatusUnauthorized).SendString(err.Error())
// 		case jwt.ValidationErrorExpired:
// 			return c.Status(http.StatusUnauthorized).SendString(err.Error() + "Token expired")
// 		default:
// 			return c.Status(http.StatusUnauthorized).SendString(err.Error())
// 		}
// 	}

// 	if !token.Valid {
// 		return c.Status(http.StatusUnauthorized).SendString("You are not authorized")
// 	}

// 	return c.Next()
// }

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// extract token from cookie
		tokenString := c.Cookies("token")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing JWT Token",
			})
		}

		// parse and validate the token
		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid JWT Token",
			})
		}

		// Store email in the context locals
		c.Locals("user", claims.Email)

		return c.Next()
	}
}
