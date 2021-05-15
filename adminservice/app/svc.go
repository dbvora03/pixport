package main

import (
	"log"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"dhruv.com/svc/models"
	"dhruv.com/svc/database"
	//"dhruv.com/svc/handlers"
	"dhruv.com/svc/router"

)

func main() {

    if err := database.Open(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	fmt.Println("Database connected as:", database.DB)
	database.DB.AutoMigrate(&productmodel.Product{})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	

	approuting.RoutingSetup(app)
	// "https://cdn.discordapp.com/attachments/791357377944485898/842179245459636284/image0.png"
	app.Listen(":3000")

}
