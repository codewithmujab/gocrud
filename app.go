package main

import (
	"log"

	database "github.com/codewithmujab/gocrud/config"
	"github.com/codewithmujab/gocrud/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//connect ke db
	database.ConnectDb()

	//fiber
	app := fiber.New()

	//routes
	routes.RouterApp(app)

	//cors
	app.Use(cors.New())

	//404
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) //404 not found
	})

	//port running
	log.Fatal(app.Listen(":3000"))
}
