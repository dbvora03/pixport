package approuting

import (
	"dhruv.com/svc/handlers"
	"github.com/gofiber/fiber/v2"
)


func RoutingSetup (app *fiber.App) error {

	api := app.Group("/api")

	api.Get("/products/all", producthandler.GetAll)
	api.Get("/products/:id", producthandler.GetOne)
	api.Post("/products/add", producthandler.AddProd)
	api.Post("/login", producthandler.Login)
	api.Post("/logout", producthandler.Logout)
	api.Delete("/delete/:id", producthandler.DeleteProduct)


	return nil

}
