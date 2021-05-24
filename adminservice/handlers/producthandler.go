package producthandler

import (
	//"time"
	"net/http"
	"dhruv.com/svc/database"
	"dhruv.com/svc/models"
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	//jwt "github.com/form3tech-oss/jwt-go"
)

const SecretKey = "secret"


func GetAll(c *fiber.Ctx) error {
	var products []productmodel.Product
	database.DB.Find(&products)
	return c.JSON(products)
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var product_instance productmodel.Product
	database.DB.Find(&product_instance, id)
	return c.Status(500).JSON(&fiber.Map{
		"success": true,
		"product": product_instance,
		"message": "Product returned successfully",
	})
}

func AddProd(c *fiber.Ctx) error {

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

	claims := token.Claims.(*jwt.StandardClaims)


	product_instance := new(productmodel.Product)
	if err := c.BodyParser(product_instance) ; err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return c.Status(400).JSON(err)
	}
	database.DB.Create(&productmodel.Product {
		Title: product_instance.Title,
		Description: product_instance.Description,
		Image: product_instance.Image,
		Price: product_instance.Price,
	})

	c.Status(400).JSON(&fiber.Map{
		"success": true,
		"product": product_instance,
		"message": "Product Added Successfully",
		"claim": claims,
	})
	
	return nil
}

func DeleteProduct (c *fiber.Ctx) error {

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

	claims := token.Claims.(*jwt.StandardClaims)



	id := c.Params("id")
	result := database.DB.Find(&productmodel.Product{}, id)

	fmt.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error": "Object doesnt exist",
		  })
	}

	database.DB.Delete(&productmodel.Product{}, id)


	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "product deleted successfully",
		"claims": claims,
	  }); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": err,
		  })
		
	}
	return nil
}


func Login (c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	user := data["user"]
	pass := data["pass"]

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: "1",
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix() ,
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map {
			"message": "incorrect credentials",
		})
	}
	cookie := fiber.Cookie {
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24 * 7),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map {
		"message": "success",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie {
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "logged out",
	})
}