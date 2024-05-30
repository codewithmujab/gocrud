package routes

import (
	"github.com/codewithmujab/gocrud/controllers"
	"github.com/gofiber/fiber/v2"
)

// next routing update ke grouping dan version
func RouterApp(app *fiber.App) {
	//default
	app.Get("/", controllers.Hello)

	//books
	app.Get("/allbooks", controllers.AllBooks)
	app.Get("/book/:id", controllers.GetBook)
	app.Post("/book", controllers.AddBook)
	app.Put("/book/:id", controllers.Update)
	app.Delete("/book/:id", controllers.Destroy)

	//employee
}
