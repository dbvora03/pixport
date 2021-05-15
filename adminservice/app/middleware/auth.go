package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"


func AuthRequest (c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")


	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){

		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map {
			"message": "unauthorized",
		})
	
	}
	fmt.Println(token)

	return err
}