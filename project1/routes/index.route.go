package routes

import (
	"restapi/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouterApp(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", controllers.GetAllUsers)
	v1.Post("/users", controllers.CreateUsers)
	v1.Get("/users/:id", controllers.GetUserById)
	v1.Patch("/users/:id", controllers.UpdateUser)
	v1.Delete("/users/:id", controllers.DeleteUser)
}

//app := fiber.New()
// api := app.Group("/api", middleware) // /api

//   v1 := api.Group("/v1", middleware)   // /api/v1
//   v1.Get("/list", handler)             // /api/v1/list
//   v1.Get("/user", handler)             // /api/v1/user

//   v2 := api.Group("/v2", middleware)   // /api/v2
//   v2.Get("/list", handler)             // /api/v2/list
//   v2.Get("/user", handler)
